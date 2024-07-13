package pingo

import (
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/code-brew-lab/pingo/pkg/netcore"
)

type Request struct {
	ip       net.IP
	fd       int
	id       netcore.ID
	doneChan chan bool
}

func NewRequest(ip net.IP) (*Request, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return nil, fmt.Errorf("pingo.NewRequest: %v", err)
	}

	doneChan := make(chan bool)

	return &Request{
		ip:       ip,
		fd:       fd,
		id:       netcore.NewID(),
		doneChan: doneChan,
	}, nil
}

func (r *Request) Make() {
	go func() {
		var seq uint16 = 0
		for {
			icmp := netcore.NewICMP(netcore.ControlKindEchoRequest, r.id, seq)
			syscall.Sendto(r.fd, icmp.Marshal(), 0, &syscall.SockaddrInet4{Addr: [4]byte(r.ip.To4())})
			seq++
			time.Sleep(time.Second * 1)
		}
	}()
}

func (r *Request) FD() int {
	return r.fd
}

func (r *Request) DoneChannel() <-chan bool {
	return r.doneChan
}

func (r *Request) Close() error {
	return syscall.Close(r.fd)
}
