package netcore

import (
	"errors"
	"fmt"
	"time"
)

type (
	Datagram struct {
		ts   time.Time
		ip   *IP
		icmp *ICMP
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
	icmp, _, err := ParseICMP(b)
	if err != nil {
		return nil, fmt.Errorf("netcore.ParseDatagram: %v", err)
	}

	return &Datagram{
		ts:   time.Now(),
		ip:   ip,
		icmp: icmp,
	}, nil
}

func NewDatagram(ip *IP, icmp *ICMP) (*Datagram, error) {
	if ip == nil {
		return nil, errors.New("netcore.NewDatagram: IP is empty")
	}

	if icmp == nil {
		return nil, errors.New("netcore.NewDatagram: ICMP is empty")
	}

	ip.datagramLen += uint16(len(icmp.Marshal()))

	return &Datagram{
		ts:   time.Now(),
		ip:   ip,
		icmp: icmp,
	}, nil
}

func (d *Datagram) Timestamp() time.Time {
	return d.ts
}

func (d *Datagram) Marshal() []byte {
	return append(d.ip.Marshal(), d.icmp.Marshal()...)
}

func (d *Datagram) IP() *IP {
	return d.ip
}

func (d *Datagram) ICMP() *ICMP {
	return d.icmp
}
