package message

import (
	"github.com/seoester/adcl/protocol/maybe"
)

type MSGFlag string

const (
	MSGFlagPM MSGFlag = "PM"
	MSGFlagME         = "ME"
)

var _ ParamAccessor = &MSGContent{}

type MSGContent struct {
	Text    string
	textStr string

	PM    maybe.Base32Value
	pmStr string
	ME    maybe.Int
	meStr string

	Flags map[string]string

	// Known additional flags
	// TS, EXT § 3.5 TS - Timestamp in MSG (EXT v1.0.8)
}

func (m *MSGContent) Positional() []string {
	return []string{m.textStr}
}

func (m *MSGContent) PosLen() int {
	return 1
}

func (m *MSGContent) PosAt(i int) string {
	switch i {
	case 0:
		return m.textStr
	default:
		panic("index out of range")
	}
}

func (m *MSGContent) Named() map[string]string {
	ma := make(map[string]string)

	for k, v := range m.Flags {
		ma[k] = v
	}

	if m.PM.IsSet {
		ma[m.pmStr[:2]] = m.pmStr[2:len(m.pmStr)]
	}
	if m.ME.IsSet {
		ma[m.meStr[:2]] = m.meStr[2:len(m.meStr)]
	}

	return ma
}

func (m *MSGContent) NamedGet(key string) (string, bool) {
	if len(key) == 2 {
		switch MSGFlag(key) {
		case MSGFlagPM:
			return m.pmStr, m.PM.IsSet
		case MSGFlagME:
			return m.meStr, m.ME.IsSet
		}
	}

	val, ok := m.Flags[key]
	return val, ok
}
