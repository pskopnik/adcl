package generator

import (
	"fmt"
)

type ParamMode int

const (
	ParamModePositional ParamMode = iota
	ParamModeNamed
)

func (p ParamMode) String() string {
	switch p {
	case ParamModePositional:
		return "positional"
	case ParamModeNamed:
		return "named"
	default:
		return fmt.Sprintf("ParamMode(%d)", p)
	}
}

type Message struct {
	Command          string
	PositionalParams []*Param
	NamedParams      []*Param
	Flags            []*Flag
}

type Param struct {
	Mode     ParamMode
	Name     string
	FlagName string
	Type     string
	Mapper   string
	Required bool
	Comment  string
}

type Flag struct {
	Comment string
}
