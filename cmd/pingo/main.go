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

	for {
		buf := make([]byte, 1024)
		numRead, err := syscall.Read(fd, buf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		datagram, err := netcore.ParseDatagram(buf[:numRead], netcore.ProtocolICMP)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Raw Datagram: %s\n", hex.EncodeToString(datagram.Raw()))
		ip := datagram.IP()
		t := datagram.Transporter()
		fmt.Println(hex.EncodeToString(ip.Marshal()))
		fmt.Print(hex.EncodeToString(t.Marshal()))
	}
}
