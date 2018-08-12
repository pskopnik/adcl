package message

var _ ParamAccessor = &CTMContent{}

type CTMContent struct {
	Protocol    string
	protocolStr string
	Port        string
	portStr     string
	Token       string
	tokenStr    string

	Flags map[string]string

	// No known additional flags
}

func (c *CTMContent) Positional() []string {
	return []string{c.protocolStr, c.portStr, c.tokenStr}
}

func (c *CTMContent) PosLen() int {
	return 3
}

func (c *CTMContent) PosAt(i int) string {
	switch i {
	case 0:
		return c.protocolStr
	case 1:
		return c.portStr
	case 2:
		return c.tokenStr
	default:
		panic("index out of range")
	}
}

func (c *CTMContent) Named() map[string]string {
	return c.Flags
}

func (c *CTMContent) NamedGet(key string) (string, bool) {
	val, ok := c.Flags[key]
	return val, ok
}
