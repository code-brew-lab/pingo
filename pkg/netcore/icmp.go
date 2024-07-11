package netcore

type (
	ICMP struct {
	}
)

func ParseICMP(b []byte) (*ICMP, int, error) {
	return &ICMP{}, 0, nil
}

func (icmp *ICMP) Marshal() []byte {
	return []byte{}
}
