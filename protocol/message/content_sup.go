package message

var _ ParamAccessor = &SUPContent{}

type SUPContent struct {
	FeatureOps  []FeatureOp
	featureStrs []string

	Flags map[string]string

	// No known additional flags
}

func (s *SUPContent) Positional() []string {
	return s.featureStrs
}

func (s *SUPContent) PosLen() int {
	return len(s.featureStrs)
}

func (s *SUPContent) PosAt(i int) string {
	return s.featureStrs[i]
}

func (s *SUPContent) Named() map[string]string {
	return s.Flags
}

func (s *SUPContent) NamedGet(key string) (string, bool) {
	val, ok := s.Flags[key]
	return val, ok
}
