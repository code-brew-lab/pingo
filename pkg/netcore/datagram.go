package netcore

import (
	"errors"
	"fmt"
	"time"
)

type (
	Datagram struct {
		ts  time.Time
		raw []byte
		ip  *IP
		t   Transporter
	}

	Transporter interface {
		Marshal() []byte
	}
)

func ParseDatagram(b []byte, p Protocol) (*Datagram, error) {
	raw := make([]byte, len(b))
	copy(raw, b)

	ip, i, err := ParseIP(b, p)
	if err != nil {
		return nil, fmt.Errorf("netcore.ParseDatagram: %v", err)
	}

	b = b[i:]
	t, _, err := parseTransporter(b, p)
	if err != nil {
		return nil, fmt.Errorf("netcore.ParseDatagram: %v", err)
	}

	return &Datagram{
		ts:  time.Now(),
		raw: raw,
		ip:  ip,
		t:   t,
	}, nil
}

func NewDatagram(ip *IP, t Transporter) (*Datagram, error) {
	if ip == nil {
		return nil, errors.New("netcore.NewDatagram: ip is empty")
	}

	if t == nil {
		return nil, errors.New("netcore.NewDatagram: transporter is empty")
	}

	return &Datagram{
		ts:  time.Now(),
		raw: append(ip.Marshal(), t.Marshal()...),
		ip:  ip,
		t:   t,
	}, nil
}

func (d *Datagram) Timestamp() time.Time {
	return d.ts
}

func (d *Datagram) Raw() []byte {
	return d.raw
}

func (d *Datagram) IP() *IP {
	return d.ip
}

func (d *Datagram) Transporter() Transporter {
	return d.t
}

func parseTransporter(b []byte, p Protocol) (Transporter, int, error) {
	switch p {
	case ProtocolICMP:
		return ParseICMP(b)
	default:
		return nil, 0, fmt.Errorf("unsupported protocol: %s", p.String())
	}
}
