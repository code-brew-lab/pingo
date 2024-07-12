package checksum

import (
	"encoding/binary"
)

func Calculate(b []byte) uint16 {
	raw := make([]byte, len(b))
	copy(raw, b)

	if len(raw)%2 != 0 {
		raw = append(raw, 0)
	}

	var sum uint32
	for i := 0; i < len(raw); i += 2 {
		sum += uint32(binary.BigEndian.Uint16(raw[i : i+2]))
	}

	for sum>>16 != 0 {
		sum = (sum & 0xFFFF) + (sum >> 16)
	}

	return ^uint16(sum)
}

func Verify(b []byte) bool {
	return Calculate(b) == 0
}
