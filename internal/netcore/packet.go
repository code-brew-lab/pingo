package netcore

import (
	"time"
)

type Packet struct {
	timestamp time.Time
	raw       []byte
}

func NewPacket(raw []byte) *Packet {
	return &Packet{
		timestamp: time.Now(),
		raw:       raw,
	}
}

func (p *Packet) Timestamp() time.Time {
	return p.timestamp
}

func (p *Packet) Raw() []byte {
	return p.raw
}
