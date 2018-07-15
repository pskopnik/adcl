package parser

import (
	"errors"
	"strconv"

	"github.com/seoester/adcl/protocol/encoding"
)

// Error variables related to MessageReader.
var (
	ErrInvalidNamedParameter = errors.New("invalid named parameter, the parameter cannot be interpreted as a named parameter")
)

type MessageReader struct {
	lexer Lexer
}

// NewMessageReader creates a new MessageReader reading the passed in string.
//
// Equivalent to:
//     var messageReader MessageReader
//     messageReader.Reset(s)
func NewMessageReader(s string) *MessageReader {
	m := &MessageReader{}
	m.lexer.Reset(s)
	return m
}

// Reset sets s as the string to read from and resets the internal state.
func (m *MessageReader) Reset(s string) {
	m.lexer.Reset(s)
}

func (m *MessageReader) ReadNamed() (param Named, err error) {
	tok, err := m.lexer.Next()
	if err != nil {
		return
	}
	if len(tok) < 2 || !(encoding.IsUpperAlpha(tok[0]) &&
		encoding.IsUpperAlphaNum(tok[1])) {
		err = ErrInvalidNamedParameter
		return
	}

	param = Named{
		Raw: tok,
	}

	return
}

func (m *MessageReader) ReturnNamed(param *Named) error {
	return m.lexer.Put(param.Raw)
}

func (m *MessageReader) ReadPositional() (param Positional, err error) {
	tok, err := m.lexer.Next()
	if err != nil {
		return
	}

	param = Positional{
		Raw: tok,
	}

	return
}

func (m *MessageReader) ReturnPositional(param *Positional) error {
	return m.lexer.Put(param.Raw)
}

// Positional represents a positional parameter. It is a helper type used by
// MessageReader. All defined methods assume that Raw contains a valid
// positional parameter.
type Positional struct {
	Raw string
}

func (p *Positional) RawValue() string {
	return p.Raw
}

func (p *Positional) ValueBytes() ([]byte, error) {
	return []byte(p.RawValue()), nil
}

func (p *Positional) ValueString() (string, error) {
	return encoding.DecodeADCString(p.RawValue())
}

func (p *Positional) ValueBase32Value() (*encoding.Base32Value, error) {
	return encoding.ParseBase32Value(p.RawValue())
}

func (p *Positional) ValueBase32Bytes() ([]byte, error) {
	return encoding.DecodeBase32String(p.RawValue())
}

func (p *Positional) ValueUint64() (uint64, error) {
	return strconv.ParseUint(p.RawValue(), 10, 64)
}

func (p *Positional) ValueInt64() (int64, error) {
	return strconv.ParseInt(p.RawValue(), 10, 64)
}

func (p *Positional) ValueFloat64() (float64, error) {
	return strconv.ParseFloat(p.RawValue(), 64)
}

func (p *Positional) ValueFloat32() (float32, error) {
	f, err := strconv.ParseFloat(p.RawValue(), 32)
	return float32(f), err
}

// Named represents a named parameter. It is a helper type used by
// MessageReader. All defined methods assume that Raw contains a valid named
// parameter.
type Named struct {
	Raw string
}

func (n *Named) Name() string {
	return n.Raw[0:2]
}

func (n *Named) RawValue() string {
	return n.Raw[2:]
}

func (n *Named) ValueBytes() ([]byte, error) {
	return []byte(n.RawValue()), nil
}

func (n *Named) ValueString() (string, error) {
	return encoding.DecodeADCString(n.RawValue())
}

func (n *Named) ValueBase32Value() (*encoding.Base32Value, error) {
	return encoding.ParseBase32Value(n.RawValue())
}

func (n *Named) ValueBase32Bytes() ([]byte, error) {
	return encoding.DecodeBase32String(n.RawValue())
}

func (n *Named) ValueUint64() (uint64, error) {
	return strconv.ParseUint(n.RawValue(), 10, 64)
}

func (n *Named) ValueInt64() (int64, error) {
	return strconv.ParseInt(n.RawValue(), 10, 64)
}

func (n *Named) ValueFloat64() (float64, error) {
	return strconv.ParseFloat(n.RawValue(), 64)
}

func (n *Named) ValueFloat32() (float32, error) {
	f, err := strconv.ParseFloat(n.RawValue(), 32)
	return float32(f), err
}
