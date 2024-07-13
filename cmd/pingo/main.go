package main

import (
	"fmt"
	"log"
	"net"

	"github.com/code-brew-lab/pingo/pkg/pingo"
)

func main() {
	req, err := pingo.NewRequest(net.IPv4(149, 0, 16, 25))
	if err != nil {
		log.Fatalln(err)
	}

	req.Make()

	dataChan, errChan := pingo.Read(req.DoneChannel(), req.FD())

	for {
		select {
		case d := <-dataChan:
			ip := d.IP()
			icmp := d.ICMP()
			fmt.Printf("[%s -> %s]", ip.SourceIP(), ip.DestinationIP())
			fmt.Printf("  ")
			fmt.Printf("Seq: %d, Kind: %s, StatusCode: %s\n", icmp.Sequence(), icmp.Kind(), icmp.Code().String(icmp.Kind()))
		case err := <-errChan:
			fmt.Println(err)
		}
	}
}
