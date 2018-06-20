package message

import (
	"errors"

	"github.com/seoester/adcl/protocol/encoding"
)

// Error variables related to header fields.
var (
	ErrInvalidCommandName = errors.New("invalid command name")
	ErrInvalidType        = errors.New("invalid message type")
)

type Type byte

const (
	TypeBroadcast        Type = 'B'
	TypeClientmessage         = 'C'
	TypeDirectmessage         = 'D'
	TypeEchomessage           = 'E'
	TypeFeaturebroadcast      = 'F'
	TypeHubmessage            = 'H'
	TypeInfomessage           = 'I'
	TypeUDPmessage            = 'U'
)

// ParseType returns a Type typed version of a byte. If the passed in byte is
// not a known Type, ErrInvalidType is returned.
func ParseType(b byte) (Type, error) {
	switch Type(b) {
	case TypeBroadcast:
		return TypeBroadcast, nil
	case TypeClientmessage:
		return TypeClientmessage, nil
	case TypeDirectmessage:
		return TypeDirectmessage, nil
	case TypeEchomessage:
		return TypeEchomessage, nil
	case TypeFeaturebroadcast:
		return TypeFeaturebroadcast, nil
	case TypeHubmessage:
		return TypeHubmessage, nil
	case TypeInfomessage:
		return TypeInfomessage, nil
	case TypeUDPmessage:
		return TypeUDPmessage, nil
	default:
		return Type(0), ErrInvalidType
	}
}

type Command string

const (
	CommandSTA Command = "STA"
	CommandSUP         = "SUP"
	CommandSID         = "SID"
	CommandINF         = "INF"
	CommandMSG         = "MSG"
	CommandSCH         = "SCH"
	CommandRES         = "RES"
	CommandCTM         = "CTM"
	CommandRCM         = "RCM"
	CommandGPA         = "GPA"
	CommandPAS         = "PAS"
	CommandQUI         = "QUI"
	CommandGET         = "GET"
	CommandGFI         = "GFI"
	CommandSND         = "SND"
)

// ParseCommand returns a Command typed version of a string. The second return
// value indicates whether the command is known, i.e. has a Command constant
// in this package.
//
// If the string is a known command, the associated Command constant is
// returned. If the string is not known, but syntactically a valid command,
// the string is returned as a Command type. If the string is not a valid
// command, ErrInvalidCommandName is returned.
//
// Using the constants allows (slightly) faster equality checking.
func ParseCommand(s string) (Command, bool, error) {
	switch Command(s) {
	case CommandSTA:
		return CommandSTA, true, nil
	case CommandSUP:
		return CommandSUP, true, nil
	case CommandSID:
		return CommandSID, true, nil
	case CommandINF:
		return CommandINF, true, nil
	case CommandMSG:
		return CommandMSG, true, nil
	case CommandSCH:
		return CommandSCH, true, nil
	case CommandRES:
		return CommandRES, true, nil
	case CommandCTM:
		return CommandCTM, true, nil
	case CommandRCM:
		return CommandRCM, true, nil
	case CommandGPA:
		return CommandGPA, true, nil
	case CommandPAS:
		return CommandPAS, true, nil
	case CommandQUI:
		return CommandQUI, true, nil
	case CommandGET:
		return CommandGET, true, nil
	case CommandGFI:
		return CommandGFI, true, nil
	case CommandSND:
		return CommandSND, true, nil
	default:
		if !(len(s) != 3 &&
			encoding.IsUpperAlpha(s[1]) &&
			encoding.IsUpperAlphaNum(s[2]) &&
			encoding.IsUpperAlphaNum(s[3])) {
			return Command(""), false, ErrInvalidCommandName
		}

		return Command(s), false, nil
	}
}

type HeaderFields interface{}

// BroadcastHeaderFields represents the additional header fields for messages
// with type Broadcast.
type BroadcastHeaderFields struct {
	MySID *encoding.Base32Value
}

// DEHeaderFields is the shared implementation of the additional header fields
// for messages of type Clientmessage, Infomessage and Hubmessage.
type CIHHeaderFields struct {
}

// ClientHeaderFields represents the additional header fields for messages of
// type Clientmessage.
type ClientHeaderFields CIHHeaderFields

// InfoHeaderFields represents the additional header fields for messages of
// type Infomessage.
type InfoHeaderFields CIHHeaderFields

// HubHeaderFields represents the additional header fields for messages of
// type Hubmessage.
type HubHeaderFields CIHHeaderFields

// DEHeaderFields is the shared implementation of the additional header fields
// for messages of type Directmessage and Echomessage.
type DEHeaderFields struct {
	MySID     *encoding.Base32Value
	TargetSID *encoding.Base32Value
}

// DirectHeaderFields represents the additional header fields for messages of
// type Directmessage.
type DirectHeaderFields DEHeaderFields

// EchoHeaderFields represents the additional header fields for messages of
// type Echomessage.
type EchoHeaderFields DEHeaderFields

// FeatureHeaderFields represents the additional header fields for messages of
// type Featurebroadcast.
type FeatureHeaderFields struct {
	MySID    *encoding.Base32Value
	Features []FeatureOp
}

// UDPHeaderFields represents the additional header fields for messages of
// type UDPmessage.
type UDPHeaderFields struct {
	MyCID *encoding.Base32Value
}
