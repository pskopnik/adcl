package message

import (
	"github.com/seoester/adcl/protocol/maybe"
)

type GETFlag string

const (
	GETFlagRE GETFlag = "RE"
)

var _ ParamAccessor = &GETContent{}

type GETContent struct {
	// Namespace is
	// Specified in BASE.
	// Known values
	// file, list; BASE $ 5.3.13. GET (BASE v1.0.3)
	// tthl; EXT $ 3.1 TIGR - Tiger tree hash support (EXT v1.0.8)
	// blom; EXT $ 3.8 BLOM - Bloom filter (EXT v1.0.8)
	Namespace     string
	namespaceStr  string
	Identifer     string
	identifierStr string
	StartPos      int
	startPosStr   string
	Bytes         int
	bytesStr      string

	RE    maybe.Int
	reStr string

	Flags map[string]string

	// Known additional flags
	// ZL; EXT § 3.3. ZLIB - Compressed communication (EXT v1.0.8)
	// BK, BH; EXT $ 3.8 BLOM - Bloom filter (EXT v1.0.8)
	// DB; EXT § 3.31 Downloaded progress report for uploaders in GET (EXT v1.0.8)
}

func (g *GETContent) Positional() []string {
	return []string{g.namespaceStr, g.identifierStr, g.startPosStr, g.bytesStr}
}

func (g *GETContent) PosLen() int {
	return 4
}

func (g *GETContent) PosAt(i int) string {
	switch i {
	case 0:
		return g.namespaceStr
	case 1:
		return g.identifierStr
	case 2:
		return g.startPosStr
	case 3:
		return g.bytesStr
	default:
		panic("index out of range")
	}
}

func (g *GETContent) Named() map[string]string {
	m := make(map[string]string)

	for k, v := range g.Flags {
		m[k] = v
	}

	if g.RE.IsSet {
		m[g.reStr[:2]] = g.reStr[2:len(g.reStr)]
	}

	return m
}

func (g *GETContent) NamedGet(key string) (string, bool) {
	if len(key) == 2 {
		switch GETFlag(key) {
		case GETFlagRE:
			return g.reStr, g.RE.IsSet
		}
	}

	val, ok := g.Flags[key]
	return val, ok
}
