package ip

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/code-brew-lab/pingo/internal/checksum"
)

type (
	Header struct {
		version     uint8 // 4 bits
		headerLen   uint8 // 4 bits
		serviceType uint8
		datagramLen uint16
		id          uint16
		flags       uint16
		ttl         uint8
		proto       uint8
		checksum    uint16
		srcIP       []byte
		dstIP       []byte
		options     []byte
	}

	HeaderBuilder struct {
		*Header
	}
)

const (
	headerMultiplier uint8 = 4
	minHeaderLen     uint8 = 20
)

func ParseHeader(h []byte) (*Header, int, error) {
	if len(h) < int(minHeaderLen) {
		return nil, 0, fmt.Errorf("ip.ParseHeader: Header length must be at least %d bytes", minHeaderLen)
	}

	version := h[0] >> 4
	if version != 4 {
		return nil, 0, fmt.Errorf("ip.ParseHeader: Unsupported IP version %d. Only IPv4 is supported", version)
	}

	headerLen := h[0] & 0x0F
	fmt.Println(hex.EncodeToString(h[:headerLen*headerMultiplier]))
	if !checksum.Verify(h[:headerLen*headerMultiplier]) {
		return nil, 0, errors.New("ip.ParseHeader: Checksum verification failed")
	}

	be := binary.BigEndian

	serviceType := h[1]

	datagramLen := be.Uint16(h[2:4])
	id := be.Uint16(h[4:6])
	flags := be.Uint16(h[6:8])

	ttl := h[8]
	proto := h[9]

	cs := be.Uint16(h[10:12])

	srcIP := make([]byte, 4)
	copy(srcIP, h[12:16])
	dstIP := make([]byte, 4)
	copy(dstIP, h[16:20])

	var options []byte
	if headerLen > 5 {
		options = make([]byte, (headerLen-5)*headerMultiplier)
		copy(options, h[20:headerLen*headerMultiplier])
	}

	return &Header{
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
		int(headerLen * headerMultiplier),
		nil
}

func NewHeaderBuilder() *HeaderBuilder {
	header := &Header{
		version:   4,
		headerLen: 5,
		proto:     1,
		ttl:       255,
		srcIP:     []byte{127, 0, 0, 1},
		dstIP:     []byte{0, 0, 0, 0},
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
	hb.proto = p
	return hb
}

func (hb *HeaderBuilder) SourceIP(ip []byte) *HeaderBuilder {
	hb.srcIP = ip
	return hb
}

func (hb *HeaderBuilder) DestinationIP(ip []byte) *HeaderBuilder {
	hb.dstIP = ip
	return hb
}

func (hb *HeaderBuilder) Build() (*Header, error) {
	if hb.version != 4 {
		return nil, fmt.Errorf("ip.HeaderBuilder.Build: Unsupported IP version %d. Only IPv4 is supported", hb.version)
	}

	headerLen := hb.headerLen * headerMultiplier
	if headerLen < minHeaderLen {
		return nil, fmt.Errorf("ip.HeaderBuilder.Build: Invalid header length %d. Header length must be at least %d bytes", headerLen, minHeaderLen)
	}

	return hb.Header, nil
}

func (h *Header) Marshal() []byte {
	buff := make([]byte, h.headerLen*headerMultiplier)
	be := binary.BigEndian

	var vh uint8 = (h.version << 4) + (h.headerLen & 0x0F)
	buff[0] = vh
	buff[1] = h.serviceType

	be.PutUint16(buff[2:4], h.datagramLen)
	be.PutUint16(buff[4:6], h.id)
	be.PutUint16(buff[6:8], h.flags)

	buff[8] = h.ttl
	buff[9] = h.proto

	i := 12
	i += copy(buff[i:i+4], h.srcIP[:])
	i += copy(buff[i:i+4], h.dstIP[:])

	copy(buff[i:h.headerLen*headerMultiplier], h.options)

	ch := checksum.Calculate(buff)
	be.PutUint16(buff[10:12], ch)

	return buff
}

// Wireshark => Req: 45000054df2100004001da0ac0a8016dd8ef2678,  Resp: 45000054000000003801c12cd8ef2678c0a8016d
// Pingo     => 45004000000000003801c12cd8ef2678c0a8016d
