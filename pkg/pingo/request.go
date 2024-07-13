package pingo

import (
	"context"
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/code-brew-lab/pingo/pkg/netcore"
)

type Request struct {
	ip net.IP
	fd int
	id netcore.ID
}

func NewRequest(ip net.IP) (*Request, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return nil, fmt.Errorf("pingo.NewRequest: %v", err)
	}

	return &Request{
		ip: ip,
		fd: fd,
		id: netcore.NewID(),
	}, nil
}

func (r *Request) Make(ctx context.Context, interval time.Duration) {
	var seq uint16 = 0
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			icmp := netcore.NewICMP(netcore.ControlKindEchoRequest, r.id, seq)
			if err := syscall.Sendto(r.fd, icmp.Marshal(), 0, &syscall.SockaddrInet4{Addr: [4]byte(r.ip.To4())}); err != nil {
				fmt.Println("Failed to send ICMP packet:", err)
				return
			}
			seq++
		}
	}
}

func (r *Request) FD() int {
	return r.fd
}

func (r *Request) ID() netcore.ID {
	return r.id
}

func (r *Request) Close() error {
	return syscall.Close(r.fd)
}
