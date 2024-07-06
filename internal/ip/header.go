package ip

import (
	"fmt"

	"github.com/code-brew-lab/pingo/internal/checksum"
)

type (
	Header struct {
		version        uint8 // 4 bits
		headerLength   uint8 // 4 bits
		serviceType    uint8
		datagramLength uint16
		id             uint16
		flags          uint16
		ttl            uint8
		protocol       uint8
		checksum       uint16
		srcIP          []byte
		dstIP          []byte
	}

	HeaderBuilder struct {
		*Header
	}
)

func NewHeaderBuilder() *HeaderBuilder {
	header := &Header{
		version:      4,
		headerLength: 5,
		protocol:     1,
		srcIP:        []byte{127, 0, 0, 1},
	}
	return &HeaderBuilder{header}
}

func (hb *HeaderBuilder) Version(v uint8) *HeaderBuilder {
	hb.version = v
	return hb
}

func (hb *HeaderBuilder) ServiceType(st uint8) *HeaderBuilder {
	hb.serviceType = st
	return hb
}

func (hb *HeaderBuilder) ID(id uint16) *HeaderBuilder {
	hb.id = id
	return hb
}

func (hb *HeaderBuilder) Flags(f uint16) *HeaderBuilder {
	hb.flags = f
	return hb
}

func (hb *HeaderBuilder) Protocol(p uint8) *HeaderBuilder {
	hb.protocol = p
	return hb
}

func (hb *HeaderBuilder) SourceIP(ip [4]byte) *HeaderBuilder {
	hb.srcIP = ip[:]
	return hb
}

func (hb *HeaderBuilder) DestinationIP(ip [4]byte) *HeaderBuilder {
	hb.dstIP = ip[:]
	return hb
}

func (hb *HeaderBuilder) Build() (*Header, error) {
	return hb.Header, nil
}

func (h *Header) Marshal() ([]byte, error) {
	buff := []byte{}

	var vh uint8 = (h.version << 4) + (h.headerLength & 0x0F)
	buff = append(buff, vh)

	buff = append(buff, h.serviceType)

	buff = append(buff, byte(h.datagramLength>>8))
	buff = append(buff, byte(h.datagramLength<<8>>8))

	buff = append(buff, byte(h.id>>8))
	buff = append(buff, byte(h.id<<8>>8))

	buff = append(buff, byte(h.flags>>8))
	buff = append(buff, byte(h.flags<<8>>8))

	buff = append(buff, h.ttl)

	buff = append(buff, h.protocol)

	buff = append(buff, byte(h.checksum>>8))
	buff = append(buff, byte(h.checksum<<8>>8))

	buff = append(buff, h.srcIP...)
	buff = append(buff, h.dstIP...)

	ch, err := checksum.Calculate(buff)
	if err != nil {
		return nil, fmt.Errorf("ip.Header.Marshal: %v", err)
	}

	buff[10] = byte(ch >> 8)
	buff[11] = byte(ch << 8 >> 8)

	return buff, nil
}

// 450000000000000000013bfd7f000001
