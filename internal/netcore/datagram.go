package netcore

import (
	"fmt"
	"time"

	"github.com/code-brew-lab/pingo/internal/netcore/ipv4"
)

type (
	Transporter interface {
		Raw() []byte
	}

	ParserFunc func(b []byte) (any, int, error)

	Datagram struct {
		ts  time.Time
		raw []byte
		ip  *ipv4.IP
	}
)

func ParseDatagram(b []byte) (*Datagram, error) {
	raw := make([]byte, len(b))
	copy(raw, b)

	ip, i, err := ipv4.Parse(b)
	if err != nil {
		return nil, fmt.Errorf("netcore.ParseDatagram: %v", err)
	}

	b = b[i:]

	return &Datagram{
		ts:  time.Now(),
		raw: raw,
		ip:  ip,
	}, nil
}

func (d *Datagram) Timestamp() time.Time {
	return d.ts
}

func (d *Datagram) Raw() []byte {
	return d.raw
}
