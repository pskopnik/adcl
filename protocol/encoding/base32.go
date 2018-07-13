package encoding

import (
	"encoding/base32"
)

var noPaddingEncoding *base32.Encoding = base32.StdEncoding.WithPadding(base32.NoPadding)

func EncodeToBase32String(src []byte) string {
	return noPaddingEncoding.EncodeToString(src)
}

func DecodeBase32String(s string) ([]byte, error) {
	return noPaddingEncoding.DecodeString(s)
}

type Base32Value struct {
	raw    []byte
	strSet bool
	str    string
}

func NewBase32Value(raw []byte) *Base32Value {
	return &Base32Value{
		raw: raw,
	}
}

func ParseBase32Value(str string) (*Base32Value, error) {
	raw, err := DecodeBase32String(str)
	if err != nil {
		return nil, err
	}

	return &Base32Value{
		raw:    raw,
		strSet: true,
		str:    str,
	}, nil
}

func (b *Base32Value) String() string {
	if !b.strSet {
		b.str = EncodeToBase32String(b.raw)
		b.strSet = true
	}

	return b.str
}

func (b *Base32Value) Raw() []byte {
	return b.raw
}
