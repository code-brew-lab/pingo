package checksum

import (
	"encoding/binary"
)

func Calculate(b []byte) uint16 {
	if len(b)%2 != 0 {
		b = append(b, 0)
	}

	var sum uint32
	for i := 0; i < len(b); i += 2 {
		sum += uint32(binary.BigEndian.Uint16(b[i : i+2]))
	}

	for sum>>16 != 0 {
		sum = (sum & 0xFFFF) + (sum >> 16)
	}

	return ^uint16(sum)
}

func Verify(b []byte) bool {
	return Calculate(b) == 0
}
