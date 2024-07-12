package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/code-brew-lab/pingo/pkg/pingo"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer syscall.Close(fd)

	doneChan := make(chan bool)
	go func() {
		dataChan, errChan := pingo.Read(doneChan, fd)

		for {
			select {
			case d := <-dataChan:
				ip := d.IP()
				icmp := d.ICMP()
				fmt.Printf("%s -> %s\n", ip.SourceIP(), ip.DestinationIP())
				fmt.Printf("Kind: %s, Status: %s\n", icmp.Kind(), icmp.Code().String(icmp.Kind()))
			case e := <-errChan:
				fmt.Println(e)
			}
		}
	}()
	time.Sleep(time.Second * 5)
	doneChan <- true
}
