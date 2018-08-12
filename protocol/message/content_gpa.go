package message

import (
	"github.com/seoester/adcl/protocol/encoding"
)

var _ ParamAccessor = &GPAContent{}

type GPAContent struct {
	Data    *encoding.Base32Value
	dataStr string

	Flags map[string]string

	// No known additional flags
}

func (g *GPAContent) Positional() []string {
	return []string{g.dataStr}
}

func (g *GPAContent) PosLen() int {
	return 1
}

func (g *GPAContent) PosAt(i int) string {
	switch i {
	case 0:
		return g.dataStr
	default:
		panic("index out of range")
	}
}

func (g *GPAContent) Named() map[string]string {
	return g.Flags
}

func (g *GPAContent) NamedGet(key string) (string, bool) {
	val, ok := g.Flags[key]
	return val, ok
}
