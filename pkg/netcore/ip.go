package netcore

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"

	"github.com/code-brew-lab/pingo/pkg/netcore/checksum"
)

type (
	IP struct {
		version     uint8
		headerLen   uint8
		serviceType uint8
		datagramLen uint16
		id          uint16
		flags       uint16
		ttl         uint8
		proto       Protocol
		checksum    uint16
		srcIP       net.IP
		dstIP       net.IP
		options     []byte
	}

	IPBuilder struct {
		*IP
	}
)

const (
	ipMultiplier uint8 = 4
	minIPLen     uint8 = 20
)

func ParseIP(b []byte, p Protocol) (*IP, int, error) {
	if len(b) < int(minIPLen) {
		return nil, 0, fmt.Errorf("netcore.ParseIP: IP length must be at least %d bytes", minIPLen)
	}

	version := b[0] >> 4
	if version != 4 {
		return nil, 0, fmt.Errorf("netcore.ParseIP: Unsupported IP version %d. Only IPv4 is supported", version)
	}

	headerLen := b[0] & 0x0F
	totalLen := headerLen * ipMultiplier
	if !checksum.Verify(b[:totalLen]) {
		return nil, 0, errors.New("netcore.ParseIP: Checksum verification failed")
	}

	be := binary.BigEndian

	serviceType := b[1]

	datagramLen := be.Uint16(b[2:4])
	id := be.Uint16(b[4:6])
	flags := be.Uint16(b[6:8])

	ttl := b[8]
	proto := ParseProtocol(b[9])
	if proto != p {
		return nil, 0, fmt.Errorf("netcore.ParseIP: Unsupported protocol: %s", p.String())
	}

	cs := be.Uint16(b[10:12])

	srcIP := make([]byte, 4)
	copy(srcIP, b[12:16])
	dstIP := make([]byte, 4)
	copy(dstIP, b[16:20])

	var options []byte
	if totalLen > minIPLen {
		options = make([]byte, (headerLen-5)*ipMultiplier)
		copy(options, b[minIPLen:totalLen])
	}

	return &IP{
			version:     version,
			headerLen:   headerLen,
			serviceType: serviceType,
			datagramLen: datagramLen,
			id:          id,
			flags:       flags,
			ttl:         ttl,
			proto:       proto,
			checksum:    cs,
			srcIP:       srcIP,
			dstIP:       dstIP,
			options:     options,
		},
		int(headerLen * ipMultiplier),
		nil
}

func NewIPBuilder(dstIP net.IP) *IPBuilder {
	ip := &IP{
		version:   4,
		headerLen: 5,
		proto:     1,
		ttl:       255,
		srcIP:     net.IPv4(127, 0, 0, 1),
		dstIP:     dstIP,
	}

	return &IPBuilder{ip}
}

func (ib *IPBuilder) Version(v uint8) *IPBuilder {
	ib.version = v
	return ib
}

func (ib *IPBuilder) ServiceType(st uint8) *IPBuilder {
	ib.serviceType = st
	return ib
}

func (ib *IPBuilder) ID(id uint16) *IPBuilder {
	ib.id = id
	return ib
}

func (ib *IPBuilder) Flags(f uint16) *IPBuilder {
	ib.flags = f
	return ib
}

func (ib *IPBuilder) Protocol(p Protocol) *IPBuilder {
	ib.proto = p
	return ib
}

func (ib *IPBuilder) SourceIP(ip net.IP) *IPBuilder {
	ib.srcIP = ip
	return ib
}

func (ib *IPBuilder) Build() (*IP, error) {
	if ib.version != 4 {
		return nil, fmt.Errorf("netcore.IPBuilder.Build: Unsupported IP version %d. Only IPv4 is supported", ib.version)
	}

	headerLen := ib.headerLen * ipMultiplier
	if headerLen < minIPLen {
		return nil, fmt.Errorf("netcore.IPBuilder.Build: Invalid header length %d. Header length must be at least %d bytes", headerLen, minIPLen)
	}

	return ib.IP, nil
}

func (ip *IP) Marshal() []byte {
	buff := make([]byte, ip.headerLen*ipMultiplier)
	be := binary.BigEndian

	var vh uint8 = (ip.version << 4) + (ip.headerLen & 0x0F)
	buff[0] = vh
	buff[1] = ip.serviceType

	be.PutUint16(buff[2:4], ip.datagramLen)
	be.PutUint16(buff[4:6], ip.id)
	be.PutUint16(buff[6:8], ip.flags)

	buff[8] = ip.ttl
	buff[9] = ip.proto.Uint8()

	i := 12
	i += copy(buff[i:i+4], ip.srcIP[:])
	i += copy(buff[i:i+4], ip.dstIP[:])

	copy(buff[i:ip.headerLen*ipMultiplier], ip.options)

	ch := checksum.Calculate(buff)
	be.PutUint16(buff[10:12], ch)

	return buff
}

func (ip *IP) Version() uint8 {
	return ip.version
}

func (ip *IP) HeaderLength() uint8 {
	return ip.headerLen
}

func (ip *IP) DatagramLength() uint16 {
	return ip.datagramLen
}

func (ip *IP) ID() uint16 {
	return ip.id
}

func (ip *IP) Flags() uint16 {
	return ip.flags
}

func (ip *IP) TTL() uint8 {
	return ip.ttl
}

func (ip *IP) Protocol() Protocol {
	return ip.proto
}

func (ip *IP) Checksum() uint16 {
	return ip.checksum
}

func (ip *IP) SourceIP() net.IP {
	return ip.srcIP
}

func (ip *IP) DestinationIP() net.IP {
	return ip.dstIP
}