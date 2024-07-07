package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/code-brew-lab/pingo/internal/ip"
)

func main() {
	h, err := ip.NewHeaderBuilder().Build()
	if err != nil {
		log.Fatalln(err)
	}

	req := h.Marshal()

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_IP)
	if err != nil {
		fmt.Println(err)
		return
	}

	syscall.Sendto(fd, req, 0, &syscall.SockaddrInet4{
		Port: 0,
		Addr: [4]byte{127, 0, 0, 1},
	})

	for {
		buf := make([]byte, 1024)
		numRead, err := syscall.Read(fd, buf)
		if err != nil {
			fmt.Println(err)
		}

		header, i, err := ip.ParseHeader(buf[:numRead])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(header)
		fmt.Println(i)
	}
}
