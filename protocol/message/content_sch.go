package message

import (
	"github.com/seoester/adcl/protocol/maybe"
)

type SCHFlag string

const (
	SCHFlagTR SCHFlag = "TR"
	SCHFlagTD         = "TD"
)

var _ ParamAccessor = &SCHContent{}

type SCHContent struct {
	SearchTerms    []SearchTerm
	searchTermStrs []string

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
	// KY; EXT § 3.17. SUDP - Encrypting UDP traffic (EXT v1.0.8)
	// GR, RX; EXT § 3.20 SEGA - Grouping of file extensions in SCH (EXT v1.0.8)
	// MT, PP, OT, NT, MR, PA, RE; EXT § 3.27 ASCH - Extended searching capability (EXT v1.0.8)
}

func (s *SCHContent) Positional() []string {
	return s.searchTermStrs
}

func (s *SCHContent) PosLen() int {
	return len(s.searchTermStrs)
}

func (s *SCHContent) PosAt(i int) string {
	return s.searchTermStrs[i]
}

func (s *SCHContent) Named() map[string]string {
	ma := make(map[string]string)

	for k, v := range s.Flags {
		ma[k] = v
	}

	if s.TR.IsSet {
		ma[s.trStr[:2]] = s.trStr[2:len(s.trStr)]
	}
	if s.TD.IsSet {
		ma[s.tdStr[:2]] = s.tdStr[2:len(s.tdStr)]
	}

	return ma
}

func (s *SCHContent) NamedGet(key string) (string, bool) {
	if len(key) == 2 {
		switch SCHFlag(key) {
		case SCHFlagTR:
			return s.trStr, s.TR.IsSet
		case SCHFlagTD:
			return s.tdStr, s.TD.IsSet
		}
	}

	val, ok := s.Flags[key]
	return val, ok
}
