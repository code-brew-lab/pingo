package main

import (
	"encoding/hex"
	"fmt"
	"syscall"

	"github.com/code-brew-lab/pingo/internal/netcore"
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

		datagram, err := netcore.ParseDatagram(buf[:numRead])
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Raw Datagram: %s\n", hex.EncodeToString(datagram.Raw()))
	}
}
