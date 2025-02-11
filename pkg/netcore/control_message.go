package netcore

type (
	ControlKind uint8
	ControlCode uint8
)

const (
	// ControlKind definitions
	ControlKindEchoReply              ControlKind = 0x0  // 0   - Echo Reply
	ControlKindDestinationUnreachable ControlKind = 0x3  // 3   - Destination Unreachable
	ControlKindEchoRequest            ControlKind = 0x8  // 8   - Echo Request
	ControlKindExtendedEchoRequest    ControlKind = 0x2A // 42  - Extended Echo Request
	ControlKindExtendedEchoReply      ControlKind = 0x2B // 43  - Extended Echo Reply
)

// ControlCode definitions for Echo Reply (ControlKind = 0)
const (
	ControlCodeEchoReply ControlCode = 0x0
)

// ControlCode definitions for Destination Unreachable (ControlKind = 3)
const (
	ControlCodeNetworkUnreachable       ControlCode = 0x0 // 0 - Destination network unreachable
	ControlCodeHostUnreachable          ControlCode = 0x1 // 1 - Destination host unreachable
	ControlCodeProtocolUnreachable      ControlCode = 0x2 // 2 - Destination protocol unreachable
	ControlCodePortUnreachable          ControlCode = 0x3 // 3 - Destination port unreachable
	ControlCodeFragmentationNeeded      ControlCode = 0x4 // 4 - Fragmentation required, and DF flag set
	ControlCodeSourceRouteFailed        ControlCode = 0x5 // 5 - Source route failed
	ControlCodeNetworkUnknown           ControlCode = 0x6 // 6 - Destination network unknown
	ControlCodeHostUnknown              ControlCode = 0x7 // 7 - Destination host unknown
	ControlCodeSourceHostIsolated       ControlCode = 0x8 // 8 - Source host isolated
	ControlCodeNetworkProhibited        ControlCode = 0x9 // 9 - Network administratively prohibited
	ControlCodeHostProhibited           ControlCode = 0xA // 10 - Host administratively prohibited
	ControlCodeNetworkUnreachableForToS ControlCode = 0xB // 11 - Network unreachable for ToS
	ControlCodeHostUnreachableForToS    ControlCode = 0xC // 12 - Host unreachable for ToS
	ControlCodeCommunicationProhibited  ControlCode = 0xD // 13 - Communication administratively prohibited
	ControlCodeHostPrecedenceViolation  ControlCode = 0xE // 14 - Host Precedence Violation
	ControlCodePrecedenceCutoffInEffect ControlCode = 0xF // 15 - Precedence cutoff in effect
)

// ControlCode definitions for Echo Request (ControlKind = 8)
const (
	ControlCodeEchoRequest ControlCode = 0x0
)

// ControlCode definitions for Extended Echo Reply (ControlKind = 43)
const (
	ControlCodeNoError                        ControlCode = 0x0
	ControlCodeMalformedQuery                 ControlCode = 0x1
	ControlCodeNoSuchInterface                ControlCode = 0x2
	ControlCodeNoSuchTableEntry               ControlCode = 0x3
	ControlCodeMultipleInterfacesSatisfyQuery ControlCode = 0x4
)

// ControlCode definitions for Extended Echo Request (ControlKind = 42)
const (
	ControlCodeExtendedEchoRequest ControlCode = 0x0
)

func ParseControlKind(kind uint8) ControlKind {
	switch kind {
	case uint8(ControlKindEchoReply):
		return ControlKindEchoReply
	case uint8(ControlKindDestinationUnreachable):
		return ControlKindDestinationUnreachable
	case uint8(ControlKindEchoRequest):
		return ControlKindEchoRequest
	case uint8(ControlKindExtendedEchoReply):
		return ControlKindExtendedEchoReply
	case uint8(ControlKindExtendedEchoRequest):
		return ControlKindExtendedEchoRequest
	default:
		return ControlKind(255) // Reserved
	}
}

func (c ControlKind) Uint8() uint8 {
	return uint8(c)
}

func (c ControlKind) String() string {
	switch c {
	case ControlKindEchoReply:
		return "Echo Reply"
	case ControlKindDestinationUnreachable:
		return "Destination Unreachable"
	case ControlKindEchoRequest:
		return "Echo Request"
	case ControlKindExtendedEchoReply:
		return "Extended Echo Reply"
	case ControlKindExtendedEchoRequest:
		return "Extended Echo Request"
	default:
		return "Reserved"
	}
}

func ParseControlCode(kind ControlKind, code uint8) ControlCode {
	switch kind {
	case ControlKindEchoReply:
		return ControlCodeEchoReply
	case ControlKindDestinationUnreachable:
		switch code {
		case uint8(ControlCodeNetworkUnreachable):
			return ControlCodeNetworkUnreachable
		case uint8(ControlCodeHostUnreachable):
			return ControlCodeHostUnreachable
		case uint8(ControlCodeProtocolUnreachable):
			return ControlCodeProtocolUnreachable
		case uint8(ControlCodePortUnreachable):
			return ControlCodePortUnreachable
		case uint8(ControlCodeFragmentationNeeded):
			return ControlCodeFragmentationNeeded
		case uint8(ControlCodeSourceRouteFailed):
			return ControlCodeSourceRouteFailed
		case uint8(ControlCodeNetworkUnknown):
			return ControlCodeNetworkUnknown
		case uint8(ControlCodeHostUnknown):
			return ControlCodeHostUnknown
		case uint8(ControlCodeSourceHostIsolated):
			return ControlCodeSourceHostIsolated
		case uint8(ControlCodeNetworkProhibited):
			return ControlCodeNetworkProhibited
		case uint8(ControlCodeHostProhibited):
			return ControlCodeHostProhibited
		case uint8(ControlCodeNetworkUnreachableForToS):
			return ControlCodeNetworkUnreachableForToS
		case uint8(ControlCodeHostUnreachableForToS):
			return ControlCodeHostUnreachableForToS
		case uint8(ControlCodeCommunicationProhibited):
			return ControlCodeCommunicationProhibited
		case uint8(ControlCodeHostPrecedenceViolation):
			return ControlCodeHostPrecedenceViolation
		case uint8(ControlCodePrecedenceCutoffInEffect):
			return ControlCodePrecedenceCutoffInEffect
		}
	case ControlKindEchoRequest:
		return ControlCodeEchoRequest
	case ControlKindExtendedEchoReply:
		switch code {
		case uint8(ControlCodeNoError):
			return ControlCodeNoError
		case uint8(ControlCodeMalformedQuery):
			return ControlCodeMalformedQuery
		case uint8(ControlCodeNoSuchInterface):
			return ControlCodeNoSuchInterface
		case uint8(ControlCodeNoSuchTableEntry):
			return ControlCodeNoSuchTableEntry
		case uint8(ControlCodeMultipleInterfacesSatisfyQuery):
			return ControlCodeMultipleInterfacesSatisfyQuery
		}
	case ControlKindExtendedEchoRequest:
		return ControlCodeExtendedEchoRequest
	}
	return ControlCode(255) // Reserved
}

func (c ControlCode) Uint8() uint8 {
	return uint8(c)
}

func (c ControlCode) String(kind ControlKind) string {
	switch kind {
	case ControlKindEchoReply:
		return "Echo Reply"
	case ControlKindDestinationUnreachable:
		switch c {
		case ControlCodeNetworkUnreachable:
			return "Destination network unreachable"
		case ControlCodeHostUnreachable:
			return "Destination host unreachable"
		case ControlCodeProtocolUnreachable:
			return "Destination protocol unreachable"
		case ControlCodePortUnreachable:
			return "Destination port unreachable"
		case ControlCodeFragmentationNeeded:
			return "Fragmentation required, and DF flag set"
		case ControlCodeSourceRouteFailed:
			return "Source route failed"
		case ControlCodeNetworkUnknown:
			return "Destination network unknown"
		case ControlCodeHostUnknown:
			return "Destination host unknown"
		case ControlCodeSourceHostIsolated:
			return "Source host isolated"
		case ControlCodeNetworkProhibited:
			return "Network administratively prohibited"
		case ControlCodeHostProhibited:
			return "Host administratively prohibited"
		case ControlCodeNetworkUnreachableForToS:
			return "Network unreachable for ToS"
		case ControlCodeHostUnreachableForToS:
			return "Host unreachable for ToS"
		case ControlCodeCommunicationProhibited:
			return "Communication administratively prohibited"
		case ControlCodeHostPrecedenceViolation:
			return "Host Precedence Violation"
		case ControlCodePrecedenceCutoffInEffect:
			return "Precedence cutoff in effect"
		}
	case ControlKindEchoRequest:
		return "Echo Request"
	case ControlKindExtendedEchoReply:
		switch c {
		case ControlCodeNoError:
			return "No Error"
		case ControlCodeMalformedQuery:
			return "Malformed Query"
		case ControlCodeNoSuchInterface:
			return "No Such Interface"
		case ControlCodeNoSuchTableEntry:
			return "No Such Table Entry"
		case ControlCodeMultipleInterfacesSatisfyQuery:
			return "Multiple Interfaces Satisfy Query"
		}
	case ControlKindExtendedEchoRequest:
		return "Extended Echo Request"
	}

	return "Reserved"
}
