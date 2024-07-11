package netcore

type Protocol uint8

const (
	ProtocolICMP       Protocol = 0x1  // 1
	ProtocolIP         Protocol = 0x4  // 4
	ProtocolTCP        Protocol = 0x6  // 6
	ProtocolUDP        Protocol = 0x11 // 17
	ProtocolMUX        Protocol = 0x12 // 18
	ProtocolRDP        Protocol = 0x1B // 27
	ProtocolSMP        Protocol = 0x79 //121
	ProtocolUnassigned Protocol = 0x92 //146
)

func ParseProtocol(proto uint8) Protocol {
	switch proto {
	case uint8(ProtocolICMP):
		return ProtocolICMP
	case uint8(ProtocolIP):
		return ProtocolIP
	case uint8(ProtocolTCP):
		return ProtocolTCP
	case uint8(ProtocolUDP):
		return ProtocolUDP
	case uint8(ProtocolMUX):
		return ProtocolMUX
	case uint8(ProtocolRDP):
		return ProtocolRDP
	case uint8(ProtocolSMP):
		return ProtocolSMP
	default:
		return ProtocolUnassigned
	}
}

func (p Protocol) Uint8() uint8 {
	return uint8(p)
}

func (p Protocol) String() string {
	switch p {
	case ProtocolICMP:
		return "ICMP"
	case ProtocolIP:
		return "IP-in-IP"
	case ProtocolTCP:
		return "TCP"
	case ProtocolUDP:
		return "UDP"
	case ProtocolMUX:
		return "MUX"
	case ProtocolRDP:
		return "RDP"
	case ProtocolSMP:
		return "SMP"
	default:
		return "Unassigned"
	}
}
