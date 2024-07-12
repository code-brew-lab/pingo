package main

import (
	"encoding/hex"
	"fmt"
	"syscall"

	"github.com/code-brew-lab/pingo/pkg/netcore"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer syscall.Close(fd)

	icmp, err := netcore.NewICMP(netcore.ControlKindEchoRequest, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	addr := &syscall.SockaddrInet4{
		Port: 0,
		Addr: [4]byte{216, 239, 38, 120},
	}

	err = syscall.Sendto(fd, icmp.Marshal(), 0, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		buff := make([]byte, 1024)
		numRead, err := syscall.Read(fd, buff)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Raw Datagram:%s\n", hex.EncodeToString(buff[:numRead]))
		d, err := netcore.ParseDatagram(buff[:numRead], netcore.ProtocolICMP)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ip := d.IP()
		icmp := d.ICMP()
		fmt.Printf("%s -> %s\n", ip.SourceIP(), ip.DestinationIP())
		fmt.Printf("Kind: %s, Status: %s\n", icmp.Kind(), icmp.Code().String(icmp.Kind()))

	}
}
