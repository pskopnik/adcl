package message

var _ ParamAccessor = &RCMContent{}

type RCMContent struct {
	Protocol    string
	protocolStr string
	Token       string
	tokenStr    string

	Flags map[string]string

	// Known additional flags
	// KY?; EXT § 3.17. SUDP - Encrypting UDP traffic (EXT v1.0.8)
}

func (r *RCMContent) Positional() []string {
	return []string{r.protocolStr, r.tokenStr}
}

func (r *RCMContent) PosLen() int {
	return 2
}

func (r *RCMContent) PosAt(i int) string {
	switch i {
	case 0:
		return r.protocolStr
	case 1:
		return r.tokenStr
	default:
		panic("index out of range")
	}
}

func (r *RCMContent) Named() map[string]string {
	return r.Flags
}

func (r *RCMContent) NamedGet(key string) (string, bool) {
	val, ok := r.Flags[key]
	return val, ok
}
