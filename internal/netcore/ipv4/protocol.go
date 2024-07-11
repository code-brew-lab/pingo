package ipv4

type Proto uint8

const (
	ProtoICMP       Proto = 0x1  // 1
	ProtoIP         Proto = 0x4  // 4
	ProtoTCP        Proto = 0x6  // 6
	ProtoUDP        Proto = 0x11 // 17
	ProtoMUX        Proto = 0x12 // 18
	ProtoRDP        Proto = 0x1B // 27
	ProtoSMP        Proto = 0x79 //121
	ProtoUnassigned Proto = 0x92 //146
)

func ParseProto(proto uint8) Proto {
	switch proto {
	case uint8(ProtoICMP):
		return ProtoICMP
	case uint8(ProtoIP):
		return ProtoIP
	case uint8(ProtoTCP):
		return ProtoTCP
	case uint8(ProtoUDP):
		return ProtoUDP
	case uint8(ProtoMUX):
		return ProtoMUX
	case uint8(ProtoRDP):
		return ProtoRDP
	case uint8(ProtoSMP):
		return ProtoSMP
	default:
		return ProtoUnassigned
	}
}

func (p Proto) Uint8() uint8 {
	return uint8(p)
}

func (p Proto) String() string {
	switch p {
	case ProtoICMP:
		return "ICMP"
	case ProtoIP:
		return "IP-in-IP"
	case ProtoTCP:
		return "TCP"
	case ProtoUDP:
		return "UDP"
	case ProtoMUX:
		return "MUX"
	case ProtoRDP:
		return "RDP"
	case ProtoSMP:
		return "SMP"
	default:
		return "Unassigned"
	}
}
