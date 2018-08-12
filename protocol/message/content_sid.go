package message

import (
	"github.com/seoester/adcl/protocol/encoding"
)

var _ ParamAccessor = &SIDContent{}

type SIDContent struct {
	SID    *encoding.Base32Value
	sidStr string

	Flags map[string]string

	// No known additional flags
}

func (s *SIDContent) Positional() []string {
	return []string{s.sidStr}
}

func (s *SIDContent) PosLen() int {
	return 1
}

func (s *SIDContent) PosAt(i int) string {
	switch i {
	case 0:
		return s.sidStr
	default:
		panic("index out of range")
	}
}

func (s *SIDContent) Named() map[string]string {
	return s.Flags
}

func (s *SIDContent) NamedGet(key string) (string, bool) {
	val, ok := s.Flags[key]
	return val, ok
}
