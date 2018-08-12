package message

import (
	"github.com/seoester/adcl/protocol/maybe"
)

type RESFlag string

const (
	RESFlagFN RESFlag = "FN"
	RESFlagSI         = "SI"
	RESFlagSL         = "SL"
	RESFlagTO         = "TO"
	RESFlagTR         = "TR"
	RESFlagTD         = "TD"
)

var _ ParamAccessor = &RESContent{}

type RESContent struct {
	FN    string
	fnStr string
	SI    int
	siStr string
	SL    maybe.Int
	slStr string
	TO    string
	toStr string

	// TR is
	// Specified in EXT § 3.1 TIGR - Tiger tree hash support (EXT v1.0.8).
	// Tiger tree Hash root, encoded with base32.
	TR    maybe.Base32Value
	trStr string
	// TD is
	// Specified in EXT § 3.1 TIGR - Tiger tree hash support (EXT v1.0.8).
	// Tree depth, index of the highest level of tree data available, root-only = 0, first level (2 leaves) = 1, second level = 2, etc…
	TD    maybe.Int
	tdStr string

	Flags map[string]string

	// Known additional flags
	// FI, FO, DA; EXT § 3.27 ASCH - Extended searching capability (EXT v1.0.8)

}

func (r *RESContent) Positional() []string {
	return []string{}
}

func (r *RESContent) PosLen() int {
	return 0
}

func (r *RESContent) PosAt(i int) string {
	panic("index out of range")
}

func (r *RESContent) Named() map[string]string {
	ma := make(map[string]string)

	for k, v := range r.Flags {
		ma[k] = v
	}

	ma[r.fnStr[:2]] = r.fnStr[2:]

	ma[r.siStr[:2]] = r.siStr[2:]

	if r.SL.IsSet {
		ma[r.slStr[:2]] = r.slStr[2:]
	}

	ma[r.toStr[:2]] = r.toStr[2:]

	if r.TR.IsSet {
		ma[r.trStr[:2]] = r.trStr[2:]
	}

	if r.TD.IsSet {
		ma[r.tdStr[:2]] = r.tdStr[2:]
	}

	return ma
}

func (r *RESContent) NamedGet(key string) (string, bool) {
	if len(key) == 2 {
		switch RESFlag(key) {
		case RESFlagFN:
			return r.fnStr[2:], true
		case RESFlagSI:
			return r.siStr[2:], true
		case RESFlagSL:
			return r.slStr[2:], r.SL.IsSet
		case RESFlagTO:
			return r.toStr[2:], true
		case RESFlagTR:
			return r.trStr[2:], r.TR.IsSet
		case RESFlagTD:
			return r.tdStr[2:], r.TD.IsSet
		}
	}

	val, ok := r.Flags[key]
	return val, ok
}
