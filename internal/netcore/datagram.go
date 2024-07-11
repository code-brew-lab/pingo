package netcore

import (
	"fmt"
	"time"

	"github.com/code-brew-lab/pingo/internal/netcore/icmp"
	"github.com/code-brew-lab/pingo/internal/netcore/ipv4"
)

type (
	Transporter interface {
		Raw() []byte
	}

	Datagram struct {
		ts  time.Time
		raw []byte
		ip  *ipv4.IP
		t   Transporter
	}
)

func ParseDatagram(b []byte, p ipv4.Proto) (*Datagram, error) {
	raw := make([]byte, len(b))
	copy(raw, b)

	ip, i, err := ipv4.Parse(b, p)
	if err != nil {
		return nil, fmt.Errorf("netcore.ParseDatagram: %v", err)
	}

	b = b[i:]
	payload, _, err := parsePayload(b, p)
	if err != nil {
		return nil, fmt.Errorf("netcore.ParseDatagram: %v", err)
	}

	return &Datagram{
		ts:  time.Now(),
		raw: raw,
		ip:  ip,
		t:   payload,
	}, nil
}

func (d *Datagram) Timestamp() time.Time {
	return d.ts
}

func (d *Datagram) Raw() []byte {
	return d.raw
}

func parsePayload(b []byte, p ipv4.Proto) (Transporter, int, error) {
	switch p {
	case ipv4.ProtoICMP:
		return icmp.Parse(b)
	default:
		return nil, 0, fmt.Errorf("unsupported protocol: %s", p.String())
	}
}
