package message

type ParamAccessor interface {
	Positional() []string
	PosLen() int
	PosAt(i int) string

	Named() map[string]string
	NamedGet(key string) (string, bool)
}

var _ ParamAccessor = &Message{}

type Message struct {
	Type         Type
	Command      Command
	HeaderFields HeaderFields

	Content ParamAccessor
}

func (m *Message) Positional() []string {
	return m.Content.Positional()
}

func (m *Message) PosLen() int {
	return m.Content.PosLen()
}

func (m *Message) PosAt(i int) string {
	return m.Content.PosAt(i)
}

func (m *Message) Named() map[string]string {
	return m.Content.Named()
}

func (m *Message) NamedGet(key string) (string, bool) {
	return m.Content.NamedGet(key)
}

var _ ParamAccessor = &GenericContent{}

type GenericContent struct {
	PositionalParams []string
	NamedParams      map[string]string
}

func (g *GenericContent) Positional() []string {
	return g.PositionalParams
}

func (g *GenericContent) PosLen() int {
	return len(g.PositionalParams)
}

func (g *GenericContent) PosAt(i int) string {
	return g.PositionalParams[i]
}

func (g *GenericContent) Named() map[string]string {
	return g.NamedParams
}

func (g *GenericContent) NamedGet(key string) (string, bool) {
	val, ok := g.NamedParams[key]
	return val, ok
}
