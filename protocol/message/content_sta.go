package message

var _ ParamAccessor = &STAContent{}

type STAContent struct {
	Code           StatusCode
	codeStr        string
	Description    string
	descriptionStr string

	Flags map[string]string

	// Known additional flags
	// FC, TL, TO, PR, FM, FB, I4, I6; BASE $ 5.3.1. STA (BASE v1.0.3)
	// RF; EXT § 3.10 RF - Referrer notification (EXT v1.0.8)
	// QP; EXT § 3.11 QP - Upload queue notification (EXT v1.0.8)
	// FC, TO, RC; EXT § 3.27 ASCH - Extended searching capability (EXT v1.0.8)
}

func (s *STAContent) Positional() []string {
	return []string{s.codeStr, s.descriptionStr}
}

func (s *STAContent) PosLen() int {
	return 2
}

func (s *STAContent) PosAt(i int) string {
	switch i {
	case 0:
		return s.codeStr
	case 1:
		return s.descriptionStr
	default:
		panic("index out of range")
	}
}

func (s *STAContent) Named() map[string]string {
	return s.Flags
}

func (s *STAContent) NamedGet(key string) (string, bool) {
	val, ok := s.Flags[key]
	return val, ok
}
