package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/code-brew-lab/gonq/pkg/dns"
	"github.com/code-brew-lab/pingo/pkg/pingo"
)

func main() {
	args := setupFlags()
	interval := time.Duration(args.interval) * time.Second
	timeout := time.Duration(args.timeout) * time.Millisecond

	var ip net.IP

	if parsedIP := net.ParseIP(args.host); parsedIP != nil {
		ip = parsedIP
	} else {
		dnsReq, err := dns.NewRequest(args.client, 53)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dnsReq.AddQuery(args.host, dns.TypeA, dns.ClassINET)

		dnsResp, err := dnsReq.Make()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ip = dnsResp.IPs()[0]
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := pingo.NewRequest(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer req.Close()

	go req.Make(cancel, interval)

	dataChan, errChan := pingo.Read(ctx, req.FD(), req.ID(), timeout)

	for {
		select {
		case d := <-dataChan:
			fmt.Println(d)
		case err := <-errChan:
			fmt.Println(err)
		case <-ctx.Done():
			return
		}
	}
}
