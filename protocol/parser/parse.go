package parser

import (
	"errors"
	"io"

	"github.com/seoester/adcl/protocol/encoding"
	"github.com/seoester/adcl/protocol/message"
)

var (
	ErrInvalidMessage         = errors.New("message invalid")
	ErrIncompleteMessage      = errors.New("message incomplete, required elements are missing")
	ErrInvalidFeatureEncoding = errors.New("feature invalid encoded in feature broadcast header")
)

// Constants which are used throughout the parser package.
const (
	space byte = ' '
	eol        = '\n'
)

func ParseMessage(m *MessageReader) (mes message.Message, err error) {
	mes, err = ParseHeader(m)
	if err != nil {
		return
	}

	mes.Content, err = ParseContent(m, mes.Command)
	if err != nil {
		return
	}

	return
}

func ParseHeader(m *MessageReader) (mes message.Message, err error) {
	fourcc, err := m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}
	if len(fourcc.RawValue()) != 4 {
		err = ErrInvalidMessage
		return
	}

	mes.Type, err = message.ParseType(fourcc.RawValue()[0])
	if err != nil {
		return
	}

	mes.Command, _, err = message.ParseCommand(fourcc.RawValue()[1:])
	if err != nil {
		return
	}

	mes.HeaderFields, err = ParseHeaderFields(m, mes.Type)
	if err != nil {
		return
	}

	return
}

func ParseHeaderFields(m *MessageReader, typ message.Type) (message.HeaderFields, error) {
	switch typ {
	case message.TypeBroadcast:
		return ParseBroadcastHeaderFields(m)
	case message.TypeClientmessage:
		return ParseCIHHeaderFields(m)
	case message.TypeDirectmessage:
		return ParseDEHeaderFields(m)
	case message.TypeEchomessage:
		return ParseDEHeaderFields(m)
	case message.TypeFeaturebroadcast:
		return ParseFeatureHeaderFields(m)
	case message.TypeHubmessage:
		return ParseCIHHeaderFields(m)
	case message.TypeInfomessage:
		return ParseCIHHeaderFields(m)
	case message.TypeUDPmessage:
		return ParseUDPHeaderFields(m)
	default:
		return nil, message.ErrInvalidType
	}
}

func ParseBroadcastHeaderFields(m *MessageReader) (fields message.BroadcastHeaderFields, err error) {
	param, err := m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}

	fields.MySID, err = param.ValueBase32Value()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}

	return
}

func ParseCIHHeaderFields(m *MessageReader) (fields message.CIHHeaderFields, err error) {
	return
}

func ParseDEHeaderFields(m *MessageReader) (fields message.DEHeaderFields, err error) {
	param, err := m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}

	fields.MySID, err = param.ValueBase32Value()
	if err != nil {
		return
	}

	param, err = m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}

	fields.TargetSID, err = param.ValueBase32Value()
	if err != nil {
		return
	}

	return
}

func ParseFeatureHeaderFields(m *MessageReader) (fields message.FeatureHeaderFields, err error) {
	param, err := m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}

	fields.MySID, err = param.ValueBase32Value()
	if err != nil {
		return
	}

	param, err = m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}

	if len(param.RawValue())%5 != 0 {
		return fields, ErrInvalidFeatureEncoding
	}

	for i := 0; i < len(param.RawValue()); i += 5 {
		var op message.FeatureOp

		switch param.RawValue()[i] {
		case '+':
			op.OpAction = message.FeatureOpAdd
		case '-':
			op.OpAction = message.FeatureOpRemove
		default:
			return fields, ErrInvalidFeatureEncoding
		}

		if !(encoding.IsUpperAlpha(param.RawValue()[i+1]) &&
			encoding.IsUpperAlphaNum(param.RawValue()[i+2]) &&
			encoding.IsUpperAlphaNum(param.RawValue()[i+3]) &&
			encoding.IsUpperAlphaNum(param.RawValue()[i+4])) {
			return fields, ErrInvalidFeatureEncoding
		}

		op.Feature = param.RawValue()[i+1 : i+5]

		fields.Features = append(fields.Features, op)
	}

	return
}

func ParseUDPHeaderFields(m *MessageReader) (fields message.UDPHeaderFields, err error) {
	param, err := m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}

	fields.MyCID, err = param.ValueBase32Value()
	if err != nil {
		return
	}

	return
}

func ParseContent(m *MessageReader, cmd message.Command) (cnt message.ParamAccessor, err error) {
	switch cmd {
	case message.CommandSTA:
		mes, err := ParseSTAContent(m)
		if err != nil {
			return nil, err
		}
		return &mes, err
	// case message.CommandSUP:
	// 	mes, err := ParseSUPContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandSID:
	// 	mes, err := ParseSIDContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandINF:
	// 	mes, err := ParseINFContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandMSG:
	// 	mes, err := ParseMSGContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandSCH:
	// 	mes, err := ParseSCHContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandRES:
	// 	mes, err := ParseRESContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandCTM:
	// 	mes, err := ParseCTMContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandRCM:
	// 	mes, err := ParseRCMContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandGPA:
	// 	mes, err := ParseGPAContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandPAS:
	// 	mes, err := ParsePASContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandQUI:
	// 	mes, err := ParseQUIContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandGET:
	// 	mes, err := ParseGETContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandGFI:
	// 	mes, err := ParseGFIContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	// case message.CommandSND:
	// 	mes, err := ParseSNDContent(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &mes, err
	default:
		mes, err := ParseGenericContent(m)
		if err != nil {
			return nil, err
		}
		return &mes, err
	}
}

func ParseGenericContent(m *MessageReader) (cnt message.GenericContent, err error) {
	positional, err := m.ReadPositional()
	for ; err == nil; positional, err = m.ReadPositional() {
		cnt.PositionalParams = append(cnt.PositionalParams, positional.RawValue())
	}

	if err == io.EOF {
		err = nil
	}

	return
}
