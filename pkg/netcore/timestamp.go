package netcore

import (
	"encoding/binary"
	"errors"
	"time"
)

type Timestamp int64

func TimestampNow() Timestamp {
	return Timestamp(time.Now().UnixNano())
}

func ParseTimestamp(b []byte) (Timestamp, error) {
	if len(b) > 8 {
		return 0, errors.New("timestamp should not be greater than 8 bytes")
	}

	padded := make([]byte, 8)
	copy(padded[8-len(b):], b)

	be := binary.BigEndian

	return Timestamp(be.Uint64(padded)), nil
}

func Duration(t1 Timestamp, t2 Timestamp) time.Duration {
	diff := t1 - t2

	if diff < 0 {
		diff *= -1
	}

	return time.Duration(diff)
}
