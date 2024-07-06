package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/code-brew-lab/pingo/internal/checksum"
	"github.com/code-brew-lab/pingo/internal/ip"
)

func main() {
	h, err := ip.NewHeaderBuilder().Build()
	if err != nil {
		log.Fatalln(err)
	}

	b, err := h.Marshal()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(hex.EncodeToString(b))
	fmt.Println(checksum.Verify(b))

	// fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// req, err := hex.DecodeString("45000054289a000040019092c0a8016dd8ef267845000054289a000040019092c0a8016dd8ef2678")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// syscall.Sendto(fd, req, 0, &syscall.SockaddrInet4{
	// 	Port: 0,
	// 	Addr: [4]byte{127, 0, 0, 1},
	// })

	// for {
	// 	buf := make([]byte, 1024)
	// 	numRead, err := syscall.Read(fd, buf)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Printf("% X\n", buf[:numRead])
	// }
}
