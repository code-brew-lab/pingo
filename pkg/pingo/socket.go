package pingo

import (
	"syscall"

	"github.com/code-brew-lab/pingo/pkg/netcore"
)

func Read(doneChan <-chan bool, fd int) (<-chan *netcore.Datagram, <-chan error) {
	dataChan := make(chan *netcore.Datagram, 10)
	errChan := make(chan error, 10)

	go func(dataChan chan<- *netcore.Datagram, errChan chan<- error, fd int) {
		defer close(dataChan)
		defer close(errChan)

		for {
			select {
			case <-doneChan:
				return
			default:
				read(dataChan, errChan, fd)
			}
		}
	}(dataChan, errChan, fd)

	return dataChan, errChan
}

func read(dataChan chan<- *netcore.Datagram, errChan chan<- error, fd int) {
	buff := make([]byte, 1024)
	numRead, err := syscall.Read(fd, buff)
	if err != nil {
		errChan <- err
		return
	}
	d, err := netcore.ParseDatagram(buff[:numRead], netcore.ProtocolICMP)
	if err != nil {
		errChan <- err
		return
	}

	dataChan <- d
}
