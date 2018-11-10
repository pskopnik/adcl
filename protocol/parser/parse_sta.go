package parser

import (
	"io"

	"github.com/seoester/adcl/protocol/message"
)

func ParseSTAContent(m *MessageReader) (mes message.STAContent, err error) {
	// cons := STAContentConstructor{&mes}

	var positionalParam Positional

	positionalParam, err = m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}

	statusCode, err := message.ParseStatusCode(positionalParam.RawValue())
	if err != nil {
		return
	}
	mes.Code = statusCode
	// cons.SetCode(statusCode, positionalParam.Raw)

	positionalParam, err = m.ReadPositional()
	if err == io.EOF {
		err = ErrIncompleteMessage
		return
	} else if err != nil {
		return
	}
	description, err := positionalParam.ValueString()
	if err != nil {
		return
	}
	mes.Description = description
	// cons.SetDescription(description, positionalParam.Raw)

	for {
		var namedParam Named
		namedParam, err = m.ReadNamed()
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}

		if mes.Flags == nil {
			mes.Flags = make(map[string]string)
		}
		mes.Flags[namedParam.Name()] = namedParam.RawValue()
	}

	return
}
