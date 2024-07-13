package netcore

import (
	"encoding/binary"
	"fmt"

	"github.com/code-brew-lab/pingo/pkg/netcore/checksum"
)

//0800f7ff00000000

type (
	ICMP struct {
		kind     ControlKind
		code     ControlCode
		checksum uint16
		id       ID
		seq      uint16
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

	be := binary.BigEndian

	kind := ParseControlKind(b[0])
	code := ParseControlCode(kind, b[1])

	id, err := ParseID(b[4:6])
	if err != nil {
		return nil, 0, fmt.Errorf("netcore.ParseICMP: %v", err)
	}

	seq := be.Uint16(b[6:])

	ch := be.Uint16(b[2:4])

	return &ICMP{
		kind:     kind,
		code:     code,
		checksum: ch,
		id:       id,
		seq:      seq,
	}, len(b), nil
}

func NewICMP(kind ControlKind, id ID, seq uint16) *ICMP {
	icmp := &ICMP{
		kind: kind,
		code: 0,
		id:   id,
		seq:  seq,
	}

	return icmp
}

func (icmp *ICMP) Marshal() []byte {
	buff := make([]byte, minICMPLength)
	be := binary.BigEndian

	buff[0] = icmp.Kind().Uint8()
	buff[1] = icmp.Code().Uint8()

	be.PutUint16(buff[4:], icmp.id.ToUint16())
	be.PutUint16(buff[6:], icmp.seq)

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

func (icmp *ICMP) ID() ID {
	return icmp.id
}

func (icmp *ICMP) Sequence() uint16 {
	return icmp.seq
}
