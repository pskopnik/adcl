package message

import (
	"github.com/seoester/adcl/protocol/encoding"
)

var _ ParamAccessor = &PASContent{}

type PASContent struct {
	Password    *encoding.Base32Value
	passwordStr string

	Flags map[string]string

	// No known additional flags
}

func (p *PASContent) Positional() []string {
	return []string{p.passwordStr}
}

func (p *PASContent) PosLen() int {
	return 1
}

func (p *PASContent) PosAt(i int) string {
	switch i {
	case 0:
		return p.passwordStr
	default:
		panic("index out of range")
	}
}

func (p *PASContent) Named() map[string]string {
	return p.Flags
}

func (p *PASContent) NamedGet(key string) (string, bool) {
	val, ok := p.Flags[key]
	return val, ok
}
