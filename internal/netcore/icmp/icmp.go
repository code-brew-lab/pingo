package icmp

type ICMP struct {
	raw []byte
}

func Parse(b []byte) (*ICMP, int, error) {
	raw := make([]byte, len(b))
	copy(raw, b)

	return &ICMP{raw: raw}, 0, nil
}

func (i ICMP) Raw() []byte {
	return i.raw
}
