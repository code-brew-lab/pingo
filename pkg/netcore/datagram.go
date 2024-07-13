package netcore

import (
	"errors"
	"fmt"
	"strings"
)

type (
	Datagram struct {
		ts   Timestamp
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
		ts:   TimestampNow(),
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
		ts:   TimestampNow(),
		ip:   ip,
		icmp: icmp,
	}, nil
}

func (d *Datagram) Timestamp() Timestamp {
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

func (d *Datagram) String() string {
	ip := d.IP()
	icmp := d.ICMP()
	rtt := Duration(d.Timestamp(), icmp.Timestamp())

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s -> %s]", ip.SourceIP(), ip.DestinationIP()))
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf("[RTT: %dms, TTL: %d]", rtt.Milliseconds(), ip.TTL()))
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf("[Seq: %d, Code: %s]", icmp.Sequence(), icmp.Code().String(icmp.Kind())))

	return sb.String()
}
