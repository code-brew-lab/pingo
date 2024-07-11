package ipv4

import (
	"errors"
	"fmt"
)

type (
	IP struct {
		raw    []byte
		header *Header
	}
)

func Parse(b []byte) (*IP, int, error) {
	header, i, err := parseHeader(b)
	if err != nil {
		return nil, 0, fmt.Errorf("ipv4.Parse: %s", err)
	}

	raw := make([]byte, len(b[:i]))
	copy(raw, b)

	return &IP{
		raw:    raw,
		header: header,
	}, i, nil
}

func NewWithHeader(h *Header) (*IP, error) {
	if h == nil {
		return nil, errors.New("ipv4.Parse: header is empty")
	}

	return &IP{
		raw:    h.Marshal(),
		header: h,
	}, nil
}
