package message

import (
	"github.com/seoester/adcl/protocol/maybe"
)

type INFFlag string

const (
	INFFlagID INFFlag = "ID"
	INFFlagPD         = "PD"
	INFFlagI4         = "I4"
	INFFlagI6         = "I6"
	INFFlagU4         = "U4"
	INFFlagU6         = "U6"
	INFFlagSS         = "SS"
	INFFlagSF         = "SF"
	INFFlagVE         = "VE"
	INFFlagUS         = "US"
	INFFlagDS         = "DS"
	INFFlagSL         = "SL"
	INFFlagAS         = "AS"
	INFFlagAM         = "AM"
	INFFlagEM         = "EM"
	INFFlagNI         = "NI"
	INFFlagDE         = "DE"
	INFFlagHN         = "HN"
	INFFlagHR         = "HR"
	INFFlagHO         = "HO"
	INFFlagTO         = "TO"
	INFFlagCT         = "CT"
	INFFlagAW         = "AW"
	INFFlagSU         = "SU"
	INFFlagRF         = "RF"
)

var _ ParamAccessor = &INFContent{}

type INFContent struct {
	// ID is
	// The CID of the client. Mandatory for C-C connections.
	// Specified in BASE.
	ID    maybe.Base32Value
	idStr string
	// PD is
	// The PID of the client. Hubs must check that the hash(PID) == CID and then discard the field before broadcasting it to other clients. Must not be sent in C-C connections.
	// Specified in BASE.
	PD    maybe.Base32Value
	pdStr string

	// I4 is
	// IPv4 address without port. A zero address (0.0.0.0) means that the server should replace it with the real IP of the client. Hubs must check that a specified address corresponds to what the client is connecting from to avoid DoS attacks and only allow trusted clients to specify a different address. Clients should use the zero address when connecting, but may opt not to do so at the user’s discretion.
	// Specified in BASE.
	I4    maybe.IP
	i4Str string
	// I6 is
	// IPv6 address without port. A zero address (::) means that the server should replace it with the IP of the client.
	// Specified in BASE.
	I6    maybe.IP
	i6Str string

	// U4 is
	// The Client UDP port.
	// Specified in BASE.
	U4    maybe.Int
	u4Str string
	// U6 is
	// Same as U4, but for IPv6.
	// Specified in BASE.
	U6    maybe.Int
	u6Str string

	// SS is
	// Share size in bytes
	// Specified in BASE.
	SS    maybe.Int
	ssStr string
	// SF is
	// Number of shared files
	// Specified in BASE.
	SF    maybe.Int
	sfStr string

	// VE is
	// Client identification, version (client-specific, a short identifier then a dotted version number is recommended)
	// Specified in BASE.
	VE    maybe.String
	veStr string

	// US is
	// Maximum upload speed, bytes/second
	// Specified in BASE.
	US    maybe.Int
	usStr string
	// DS is
	// Maximum downloads speed, bytes/second
	// Specified in BASE.
	DS    maybe.Int
	dsStr string

	// SL is
	// Maximum simultaneous upload connections (slots)
	// Specified in BASE.
	SL    maybe.Int
	slStr string

	// AS is
	// Automatic slot allocator speed limit, bytes/sec. The client keeps opening slots as long as its total upload speed doesn’t exceed this value.
	// Specified in BASE.
	AS    maybe.Int
	asStr string
	// AM is
	// Minimum simultaneous upload connections in automatic slot manager mode
	// Specified in BASE.
	AM    maybe.Int
	amStr string

	// EM is
	// E-mail address
	// Specified in BASE.
	EM    maybe.String
	emStr string
	// NI is
	// Nickname (or hub name). The hub must ensure that this is unique in the hub up to case-sensitivity. Valid are all characters in the Unicode character set with code point above 32, although hubs may limit this further as they like with an appropriate error message.
	// Specified in BASE.
	NI    maybe.String
	niStr string
	// DE is
	// Description. Valid are all characters in the Unicode character set with code point equal to or greater than 32.
	// Specified in BASE.
	DE    maybe.String
	deStr string

	// HN is
	// Hubs where user is a normal user and in NORMAL state
	// Specified in BASE.
	HN    maybe.Int
	hnStr string
	// HR is
	// Hubs where user is registered (had to supply password) and in NORMAL state
	// Specified in BASE.
	HR    maybe.Int
	hrStr string
	// HO is
	// Hubs where user is op and in NORMAL state
	// Specified in BASE.
	HO    maybe.Int
	hoStr string

	// TO is
	// Token, as received in RCM/CTM, when establishing a C-C connection.
	// Specified in BASE.
	TO    maybe.String
	toStr string

	// CT is
	// Client (user) type, 1=bot, 2=registered user, 4=operator, 8=super user, 16=hub owner, 32=hub (used when the hub sends an INF about itself). Multiple types are specified by adding the numbers together.
	// Specified in BASE.
	CT    maybe.Int
	ctStr string
	// AW is
	// 1=Away, 2=Extended away, not interested in hub chat (hubs may skip sending broadcast type MSG commands to clients with this flag)
	// Specified in BASE.
	AW    maybe.Int
	awStr string

	// SU is
	// Comma-separated list of feature FOURCC’s. This notifies other clients of extended capabilities of the connecting client.
	// Specified in BASE.
	SU    []string
	suStr string

	// RF is
	// URL of referrer (hub in case of redirect, web page)
	// Specified in BASE.
	RF    maybe.String
	rfStr string

	Flags map[string]string

	// Known additional flags
	// HH, WS, NE, OW, UC, SS, SF, MS, XS, ML, XL, MU, MR, MO, XU, XR, XO, MC, UP; EXT $ 3.4 PING - Pinger extension (EXT v1.0.8)
	// LC; EXT § 3.13 LC - Locale specification (EXT v1.0.8)
	// KP; EXT § 3.16 KEYP - Certificate substitution protection in conjunction with ADCS (EXT v1.0.8)
	// FO; EXT § 3.21 FO - Failover hub addresses (EXT v1.0.8)
	// FS; EXT § 3.22 FS - Free slots in client (EXT v1.0.8)
	// AP; EXT § 3.24 Application and version separation in INF (EXT v1.0.8)
	// RP; EXT § 3.32 RDEX - Redirects Extended (EXT v1.0.8)

}

func (i *INFContent) Positional() []string {
	return []string{}
}

func (i *INFContent) PosLen() int {
	return 0
}

func (i *INFContent) PosAt(_ int) string {
	panic("index out of range")
}

func (i *INFContent) Named() map[string]string {
	m := make(map[string]string)

	for k, v := range i.Flags {
		m[k] = v
	}

	if i.ID.IsSet {
		m[i.idStr[:2]] = i.idStr[2:len(i.idStr)]
	}
	if i.PD.IsSet {
		m[i.pdStr[:2]] = i.pdStr[2:len(i.pdStr)]
	}
	if i.I4.IsSet {
		m[i.i4Str[:2]] = i.i4Str[2:len(i.i4Str)]
	}
	if i.I6.IsSet {
		m[i.i6Str[:2]] = i.i6Str[2:len(i.i6Str)]
	}
	if i.U4.IsSet {
		m[i.u4Str[:2]] = i.u4Str[2:len(i.u4Str)]
	}
	if i.U6.IsSet {
		m[i.u6Str[:2]] = i.u6Str[2:len(i.u6Str)]
	}
	if i.SS.IsSet {
		m[i.ssStr[:2]] = i.ssStr[2:len(i.ssStr)]
	}
	if i.SF.IsSet {
		m[i.sfStr[:2]] = i.sfStr[2:len(i.sfStr)]
	}
	if i.VE.IsSet {
		m[i.veStr[:2]] = i.veStr[2:len(i.veStr)]
	}
	if i.US.IsSet {
		m[i.usStr[:2]] = i.usStr[2:len(i.usStr)]
	}
	if i.DS.IsSet {
		m[i.dsStr[:2]] = i.dsStr[2:len(i.dsStr)]
	}
	if i.SL.IsSet {
		m[i.slStr[:2]] = i.slStr[2:len(i.slStr)]
	}
	if i.AS.IsSet {
		m[i.asStr[:2]] = i.asStr[2:len(i.asStr)]
	}
	if i.AM.IsSet {
		m[i.amStr[:2]] = i.amStr[2:len(i.amStr)]
	}
	if i.EM.IsSet {
		m[i.emStr[:2]] = i.emStr[2:len(i.emStr)]
	}
	if i.NI.IsSet {
		m[i.niStr[:2]] = i.niStr[2:len(i.niStr)]
	}
	if i.DE.IsSet {
		m[i.deStr[:2]] = i.deStr[2:len(i.deStr)]
	}
	if i.HN.IsSet {
		m[i.hnStr[:2]] = i.hnStr[2:len(i.hnStr)]
	}
	if i.HR.IsSet {
		m[i.hrStr[:2]] = i.hrStr[2:len(i.hrStr)]
	}
	if i.HO.IsSet {
		m[i.hoStr[:2]] = i.hoStr[2:len(i.hoStr)]
	}
	if i.TO.IsSet {
		m[i.toStr[:2]] = i.toStr[2:len(i.toStr)]
	}
	if i.CT.IsSet {
		m[i.ctStr[:2]] = i.ctStr[2:len(i.ctStr)]
	}
	if i.AW.IsSet {
		m[i.awStr[:2]] = i.awStr[2:len(i.awStr)]
	}
	if len(i.SU) > 0 {
		m[i.suStr[:2]] = i.suStr[2:len(i.suStr)]
	}
	if i.RF.IsSet {
		m[i.rfStr[:2]] = i.rfStr[2:len(i.rfStr)]
	}

	return m
}

func (i *INFContent) NamedGet(key string) (string, bool) {
	if len(key) == 2 {
		switch INFFlag(key) {
		case INFFlagID:
			return i.idStr, i.ID.IsSet
		case INFFlagPD:
			return i.pdStr, i.PD.IsSet
		case INFFlagI4:
			return i.i4Str, i.I4.IsSet
		case INFFlagI6:
			return i.i6Str, i.I6.IsSet
		case INFFlagU4:
			return i.u4Str, i.U4.IsSet
		case INFFlagU6:
			return i.u6Str, i.U6.IsSet
		case INFFlagSS:
			return i.ssStr, i.SS.IsSet
		case INFFlagSF:
			return i.sfStr, i.SF.IsSet
		case INFFlagVE:
			return i.veStr, i.VE.IsSet
		case INFFlagUS:
			return i.usStr, i.US.IsSet
		case INFFlagDS:
			return i.dsStr, i.DS.IsSet
		case INFFlagSL:
			return i.slStr, i.SL.IsSet
		case INFFlagAS:
			return i.asStr, i.AS.IsSet
		case INFFlagAM:
			return i.amStr, i.AM.IsSet
		case INFFlagEM:
			return i.emStr, i.EM.IsSet
		case INFFlagNI:
			return i.niStr, i.NI.IsSet
		case INFFlagDE:
			return i.deStr, i.DE.IsSet
		case INFFlagHN:
			return i.hnStr, i.HN.IsSet
		case INFFlagHR:
			return i.hrStr, i.HR.IsSet
		case INFFlagHO:
			return i.hoStr, i.HO.IsSet
		case INFFlagTO:
			return i.toStr, i.TO.IsSet
		case INFFlagCT:
			return i.ctStr, i.CT.IsSet
		case INFFlagAW:
			return i.awStr, i.AW.IsSet
		case INFFlagSU:
			return i.suStr, len(i.SU) > 0
		case INFFlagRF:
			return i.rfStr, i.RF.IsSet
		}
	}

	val, ok := i.Flags[key]
	return val, ok
}
