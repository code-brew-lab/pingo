package netcore

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/code-brew-lab/pingo/pkg/netcore/checksum"
)

//0800f7ff00000000

type (
	ICMP struct {
		kind       ControlKind
		code       ControlCode
		checksum   uint16
		headerRest []byte
	}
)

const (
	minICMPLength uint16 = 8
	maxICMPLength uint16 = 576
)

func ParseICMP(b []byte) (*ICMP, int, error) {
	if len(b) < int(minICMPLength) {
		return nil, 0, fmt.Errorf("netcore.ParseICMP: ICMP length must be at least %d bytes", minICMPLength)
	}
	if len(b) > int(maxICMPLength) {
		return nil, 0, fmt.Errorf("netcore.ParseICMP: ICMP length must not exceed %d bytes", maxICMPLength)
	}
	if !checksum.Verify(b) {
		return nil, 0, errors.New("netcore.ParseICMP: Checksum verification failed")
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

func NewICMP(kind ControlKind, headerRest ...byte) (*ICMP, error) {
	if len(headerRest) != 0 && len(headerRest) != 4 {
		return nil, errors.New("netcore.NewICMP: Reaming header must be 4 bytes long")
	}

	icmp := &ICMP{
		kind:       kind,
		code:       0,
		headerRest: headerRest,
	}

	return icmp, nil
}

func (icmp *ICMP) Marshal() []byte {
	buff := make([]byte, minICMPLength)
	be := binary.BigEndian

	buff[0] = icmp.Kind().Uint8()
	buff[1] = icmp.Code().Uint8()

	copy(buff[4:8], icmp.headerRest)

	ch := checksum.Calculate(buff)
	be.PutUint16(buff[2:4], ch)

	return buff
}

func (icmp *ICMP) Kind() ControlKind {
	return icmp.kind
}

func (icmp *ICMP) Code() ControlCode {
	return icmp.code
}

func (icmp *ICMP) HeaderRest() []byte {
	return icmp.headerRest
}
