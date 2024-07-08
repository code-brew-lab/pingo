package main

import (
	"encoding/hex"
	"fmt"
	"syscall"

	"github.com/code-brew-lab/pingo/internal/ip"
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
		}

		packet := netcore.NewPacket(buf[:numRead])

		fmt.Printf("Raw Packet: %s\n", hex.EncodeToString(packet.Raw()))

		header, i, err := ip.ParseHeader(packet.Raw())
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(header)
		fmt.Println(i)
	}
}

/*
	Wireshark:
		Packet_1 [RESP]: 45000054000000003801c12cd8ef2678c0a8016d0000f2c767470000668bcac80001899808090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637

		Packet_2 [RESP]: 45000054000000003801c12cd8ef2678c0a8016d0000dffa67470001668bcac900019c6308090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637
*/

/*
	Pingo:
		Packet_1 [RESP]:	45004000000000003801c12cd8ef2678c0a8016d0000f2c767470000668bcac80001899808090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637

		Packet_2 [RESP]: 45004000000000003801c12cd8ef2678c0a8016d0000dffa67470001668bcac900019c6308090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637
*/
