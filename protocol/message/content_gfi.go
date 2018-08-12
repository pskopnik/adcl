package message

var _ ParamAccessor = &GFIContent{}

type GFIContent struct {
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

	Flags map[string]string

	// No known additional flags
}

func (g *GFIContent) Positional() []string {
	return []string{g.namespaceStr, g.identifierStr}
}

func (g *GFIContent) PosLen() int {
	return 2
}

func (g *GFIContent) PosAt(i int) string {
	switch i {
	case 0:
		return g.namespaceStr
	case 1:
		return g.identifierStr
	default:
		panic("index out of range")
	}
}

func (g *GFIContent) Named() map[string]string {
	return g.Flags
}

func (g *GFIContent) NamedGet(key string) (string, bool) {
	val, ok := g.Flags[key]
	return val, ok
}
