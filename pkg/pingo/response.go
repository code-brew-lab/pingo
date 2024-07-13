package pingo

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/code-brew-lab/pingo/pkg/netcore"
)

func Read(ctx context.Context, fd int, id netcore.ID, timeout time.Duration) (<-chan *netcore.Datagram, <-chan error) {
	dataChan := make(chan *netcore.Datagram, 10)
	errChan := make(chan error, 10)
	ticker := time.NewTicker(time.Nanosecond)

	go func() {
		defer close(dataChan)
		defer close(errChan)
		defer ticker.Stop()

		for {
			timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			select {
			case <-ctx.Done():
				return
			case <-timeoutCtx.Done():
				errChan <- fmt.Errorf("pingo.Read: %v", timeoutCtx.Err())
			case <-ticker.C:
				datagram, err := read(fd)
				if datagram.ICMP().ID() != id {
					continue
				}
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
