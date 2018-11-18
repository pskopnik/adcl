package message

type ParamAccessor interface {
	Positional() []string
	PosLen() int
	PosAt(i int) string

	Named() map[string]string
	NamedGet(key string) (string, bool)
}
