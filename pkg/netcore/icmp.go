package netcore

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/code-brew-lab/pingo/pkg/netcore/checksum"
)

type (
	ICMP struct {
		kind       ControlKind
		code       ControlCode
		checksum   uint16
		headerRest []byte
	}
)

const (
	minICMPLen uint16 = 8
	maxICMPLen uint16 = 576
)

func ParseICMP(b []byte) (*ICMP, int, error) {
	if len(b) < int(minICMPLen) {
		return nil, 0, fmt.Errorf("netcore.ParseICMP: ICMP length must be at least %d bytes", minICMPLen)
	}
	if len(b) > int(maxICMPLen) {
		return nil, 0, fmt.Errorf("netcore.ParseICMP: ICMP length must not exceed %d bytes", maxICMPLen)
	}
	if !checksum.Verify(b) {
		return nil, 0, errors.New("netcore.ParseIP: Checksum verification failed")
	}

	be := binary.BigEndian

	kind := ParseControlKind(b[0])
	code := ParseControlCode(kind, b[1])

	ch := be.Uint16(b[2:4])
	hr := b[4:8]

	return &ICMP{
		kind:       kind,
		code:       code,
		checksum:   ch,
		headerRest: hr,
	}, len(b), nil
}

func (icmp *ICMP) Marshal() []byte {
	buff := make([]byte, minICMPLen)
	be := binary.BigEndian

	buff[0] = icmp.Kind().Uint8()
	buff[1] = icmp.Code().Uint8()

	be.PutUint16(buff[2:4], icmp.checksum)

	copy(buff[4:8], icmp.headerRest)

	return buff
}

func (icmp *ICMP) Kind() ControlKind {
	return icmp.kind
}

func (icmp *ICMP) Code() ControlCode {
	return icmp.code
}

func (icmp *ICMP) Checksum() uint16 {
	return icmp.checksum
}

func (icmp *ICMP) HeaderRest() [4]byte {
	return [4]byte(icmp.headerRest)
}
