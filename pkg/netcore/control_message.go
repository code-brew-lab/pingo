package netcore

type (
	ControlKind uint8
	ControlCode uint8
)

const (
	// ControlKind definitions
	ControlKindEchoReply                    ControlKind = 0x0  // 0 - Echo Reply
	ControlKindDestinationUnreachable       ControlKind = 0x3  // 3 - Destination Unreachable
	ControlKindSourceQuench                 ControlKind = 0x4  // 4 - Source Quench
	ControlKindRedirectMessage              ControlKind = 0x5  // 5 - Redirect Message
	ControlKindEchoRequest                  ControlKind = 0x8  // 8 - Echo Request
	ControlKindRouterAdvertisement          ControlKind = 0x9  // 9 - Router Advertisement
	ControlKindRouterSolicitation           ControlKind = 0xA  // 10 - Router Solicitation
	ControlKindTimeExceeded                 ControlKind = 0xB  // 11 - Time Exceeded
	ControlKindParameterProblem             ControlKind = 0xC  // 12 - Parameter Problem: Bad IP header
	ControlKindTimestamp                    ControlKind = 0xD  // 13 - Timestamp
	ControlKindTimestampReply               ControlKind = 0xE  // 14 - Timestamp Reply
	ControlKindInformationRequest           ControlKind = 0xF  // 15 - Information Request
	ControlKindInformationReply             ControlKind = 0x10 // 16 - Information Reply
	ControlKindAddressMaskRequest           ControlKind = 0x11 // 17 - Address Mask Request
	ControlKindAddressMaskReply             ControlKind = 0x12 // 18 - Address Mask Reply
	ControlKindTraceroute                   ControlKind = 0x1E // 30 - Traceroute
	ControlKindDeprecatedDatagramConversion ControlKind = 0x1F // 31 - Deprecated Datagram Conversion Error
	ControlKindExtendedEchoRequest          ControlKind = 0x2A // 42 - Extended Echo Request
	ControlKindExtendedEchoReply            ControlKind = 0x2B // 43 - Extended Echo Reply
)

// ControlCode definitions for Echo Reply (ControlKind = 0)
const (
	ControlCodeEchoReply ControlCode = 0x0 // 0 - Echo reply
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

// ControlCode definitions for Redirect Message (ControlKind = 5)
const (
	ControlCodeRedirectDatagramForNetwork ControlCode = 0x0 // 0 - Redirect Datagram for the Network
	ControlCodeRedirectDatagramForHost    ControlCode = 0x1 // 1 - Redirect Datagram for the Host
	ControlCodeRedirectDatagramForToSNet  ControlCode = 0x2 // 2 - Redirect Datagram for the ToS & network
	ControlCodeRedirectDatagramForToSHost ControlCode = 0x3 // 3 - Redirect Datagram for the ToS & host
)

// ControlCode definitions for Time Exceeded (ControlKind = 11)
const (
	ControlCodeTTLExpiredInTransit            ControlCode = 0x0 // 0 - TTL expired in transit
	ControlCodeFragmentReassemblyTimeExceeded ControlCode = 0x1 // 1 - Fragment reassembly time exceeded
)

// ControlCode definitions for Parameter Problem (ControlKind = 12)
const (
	ControlCodePointerIndicatesError ControlCode = 0x0 // 0 - Pointer indicates the error
	ControlCodeMissingRequiredOption ControlCode = 0x1 // 1 - Missing a required option
	ControlCodeBadLength             ControlCode = 0x2 // 2 - Bad length
)

// ControlCode definitions for Extended Echo Reply (ControlKind = 43)
const (
	ControlCodeNoError                        ControlCode = 0x0 // 0 - No Error
	ControlCodeMalformedQuery                 ControlCode = 0x1 // 1 - Malformed Query
	ControlCodeNoSuchInterface                ControlCode = 0x2 // 2 - No Such Interface
	ControlCodeNoSuchTableEntry               ControlCode = 0x3 // 3 - No Such Table Entry
	ControlCodeMultipleInterfacesSatisfyQuery ControlCode = 0x4 // 4 - Multiple Interfaces Satisfy Query
)

func ParseControlKind(kind uint8) ControlKind {
	switch kind {
	case uint8(ControlKindEchoReply):
		return ControlKindEchoReply
	case uint8(ControlKindDestinationUnreachable):
		return ControlKindDestinationUnreachable
	case uint8(ControlKindSourceQuench):
		return ControlKindSourceQuench
	case uint8(ControlKindRedirectMessage):
		return ControlKindRedirectMessage
	case uint8(ControlKindEchoRequest):
		return ControlKindEchoRequest
	case uint8(ControlKindRouterAdvertisement):
		return ControlKindRouterAdvertisement
	case uint8(ControlKindRouterSolicitation):
		return ControlKindRouterSolicitation
	case uint8(ControlKindTimeExceeded):
		return ControlKindTimeExceeded
	case uint8(ControlKindParameterProblem):
		return ControlKindParameterProblem
	case uint8(ControlKindTimestamp):
		return ControlKindTimestamp
	case uint8(ControlKindTimestampReply):
		return ControlKindTimestampReply
	case uint8(ControlKindInformationRequest):
		return ControlKindInformationRequest
	case uint8(ControlKindInformationReply):
		return ControlKindInformationReply
	case uint8(ControlKindAddressMaskRequest):
		return ControlKindAddressMaskRequest
	case uint8(ControlKindAddressMaskReply):
		return ControlKindAddressMaskReply
	case uint8(ControlKindTraceroute):
		return ControlKindTraceroute
	case uint8(ControlKindDeprecatedDatagramConversion):
		return ControlKindDeprecatedDatagramConversion
	case uint8(ControlKindExtendedEchoRequest):
		return ControlKindExtendedEchoRequest
	case uint8(ControlKindExtendedEchoReply):
		return ControlKindExtendedEchoReply
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
	case ControlKindSourceQuench:
		return "Source Quench"
	case ControlKindRedirectMessage:
		return "Redirect Message"
	case ControlKindEchoRequest:
		return "Echo Request"
	case ControlKindRouterAdvertisement:
		return "Router Advertisement"
	case ControlKindRouterSolicitation:
		return "Router Solicitation"
	case ControlKindTimeExceeded:
		return "Time Exceeded"
	case ControlKindParameterProblem:
		return "Parameter Problem: Bad IP header"
	case ControlKindTimestamp:
		return "Timestamp"
	case ControlKindTimestampReply:
		return "Timestamp Reply"
	case ControlKindInformationRequest:
		return "Information Request"
	case ControlKindInformationReply:
		return "Information Reply"
	case ControlKindAddressMaskRequest:
		return "Address Mask Request"
	case ControlKindAddressMaskReply:
		return "Address Mask Reply"
	case ControlKindTraceroute:
		return "Traceroute"
	case ControlKindDeprecatedDatagramConversion:
		return "Deprecated Datagram Conversion Error"
	case ControlKindExtendedEchoRequest:
		return "Extended Echo Request"
	case ControlKindExtendedEchoReply:
		return "Extended Echo Reply"
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
	case ControlKindRedirectMessage:
		switch code {
		case uint8(ControlCodeRedirectDatagramForNetwork):
			return ControlCodeRedirectDatagramForNetwork
		case uint8(ControlCodeRedirectDatagramForHost):
			return ControlCodeRedirectDatagramForHost
		case uint8(ControlCodeRedirectDatagramForToSNet):
			return ControlCodeRedirectDatagramForToSNet
		case uint8(ControlCodeRedirectDatagramForToSHost):
			return ControlCodeRedirectDatagramForToSHost
		}
	case ControlKindTimeExceeded:
		switch code {
		case uint8(ControlCodeTTLExpiredInTransit):
			return ControlCodeTTLExpiredInTransit
		case uint8(ControlCodeFragmentReassemblyTimeExceeded):
			return ControlCodeFragmentReassemblyTimeExceeded
		}
	case ControlKindParameterProblem:
		switch code {
		case uint8(ControlCodePointerIndicatesError):
			return ControlCodePointerIndicatesError
		case uint8(ControlCodeMissingRequiredOption):
			return ControlCodeMissingRequiredOption
		case uint8(ControlCodeBadLength):
			return ControlCodeBadLength
		}
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
	}
	return ControlCode(255) // Reserved
}

func (c ControlCode) Uint8() uint8 {
	return uint8(c)
}

func (c ControlCode) String(kind ControlKind) string {
	switch kind {
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
	case ControlKindRedirectMessage:
		switch c {
		case ControlCodeRedirectDatagramForNetwork:
			return "Redirect Datagram for the Network"
		case ControlCodeRedirectDatagramForHost:
			return "Redirect Datagram for the Host"
		case ControlCodeRedirectDatagramForToSNet:
			return "Redirect Datagram for the ToS & network"
		case ControlCodeRedirectDatagramForToSHost:
			return "Redirect Datagram for the ToS & host"
		}
	case ControlKindTimeExceeded:
		switch c {
		case ControlCodeTTLExpiredInTransit:
			return "TTL expired in transit"
		case ControlCodeFragmentReassemblyTimeExceeded:
			return "Fragment reassembly time exceeded"
		}
	case ControlKindParameterProblem:
		switch c {
		case ControlCodePointerIndicatesError:
			return "Pointer indicates the error"
		case ControlCodeMissingRequiredOption:
			return "Missing a required option"
		case ControlCodeBadLength:
			return "Bad length"
		}
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
	}
	return "Reserved"
}
