package message

var _ ParamAccessor = &SNDContent{}

type SNDContent struct {
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

	Flags map[string]string

	// Known additional flags
	// ZL; EXT § 3.3. ZLIB - Compressed communication (EXT v1.0.8)
}

func (s *SNDContent) Positional() []string {
	return []string{s.namespaceStr, s.identifierStr, s.startPosStr, s.bytesStr}
}

func (s *SNDContent) PosLen() int {
	return 4
}

func (s *SNDContent) PosAt(i int) string {
	switch i {
	case 0:
		return s.namespaceStr
	case 1:
		return s.identifierStr
	case 2:
		return s.startPosStr
	case 3:
		return s.bytesStr
	default:
		panic("index out of range")
	}
}

func (s *SNDContent) Named() map[string]string {
	return s.Flags
}

func (s *SNDContent) NamedGet(key string) (string, bool) {
	val, ok := s.Flags[key]
	return val, ok
}
