package pingo

import (
	"context"
	"syscall"
	"time"

	"github.com/code-brew-lab/pingo/pkg/netcore"
)

func Read(doneChan <-chan bool, fd int) (<-chan *netcore.Datagram, <-chan error) {
	dataChan := make(chan *netcore.Datagram, 10)
	errChan := make(chan error, 10)

	go func() {
		defer close(dataChan)
		defer close(errChan)

		for {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			select {
			case <-doneChan:
				return
			case <-ctx.Done():
				errChan <- ctx.Err()
			default:
				datagram, err := read(fd)
				if err != nil {
					errChan <- err
				} else {
					dataChan <- datagram
				}
			}
		}
	}()

	return dataChan, errChan
}

func read(fd int) (*netcore.Datagram, error) {
	buff := make([]byte, 1024)
	numRead, err := syscall.Read(fd, buff)
	if err != nil {
		return nil, err
	}
	d, err := netcore.ParseDatagram(buff[:numRead], netcore.ProtocolICMP)
	if err != nil {
		return nil, err
	}

	return d, nil
}
