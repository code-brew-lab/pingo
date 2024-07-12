package main

import (
	"encoding/hex"
	"fmt"
	"net"
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

	ip, err := netcore.NewIPBuilder(net.IPv4(216, 239, 38, 120)).
		SourceIP(net.IPv4(198, 19, 249, 159)).
		Protocol(netcore.ProtocolICMP).
		Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	icmp, err := netcore.NewICMP(netcore.ControlKindEchoRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	datagram, err := netcore.NewDatagram(ip, icmp)
	if err != nil {
		fmt.Println(err)
		return
	}

	addr := &syscall.SockaddrInet4{
		Port: 0,
		Addr: [4]byte{198, 19, 249, 1},
	}

	err = syscall.Sendto(fd, datagram.Marshal(), 0, addr)
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

		datagram, err := netcore.ParseDatagram(buff[:numRead], netcore.ProtocolICMP)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Raw Datagram: %s\n", hex.EncodeToString(datagram.Marshal()))
		ip := datagram.IP()
		icmp := datagram.ICMP()
		fmt.Printf("%s -> %s\n", ip.SourceIP(), ip.DestinationIP())
		fmt.Printf("Kind: %s, Status: %s\n", icmp.Kind(), icmp.Code().String(icmp.Kind()))

	}
}
