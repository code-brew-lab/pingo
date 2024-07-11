package icmp

type (
	Header struct {
		kind       uint8
		code       uint8
		checksum   uint16
		identifier uint16
		sequence   uint16
		timestamp  [8]byte
	}

	controlMessage struct {
		kind uint8
		code uint8
	}
)
