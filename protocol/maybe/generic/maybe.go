package maybe

import (
	"github.com/cheekybits/genny/generic"
)

type Type generic.Type

type MaybeType struct {
	Value Type
	IsSet bool
}

func (m *MaybeType) Get() (Type, bool) {
	return m.Value, m.IsSet
}

func (m *MaybeType) GetDefault(def Type) Type {
	if m.IsSet {
		return m.Value
	} else {
		return def
	}
}

func (m *MaybeType) Set(val Type) {
	m.Value = val
	m.IsSet = true
}

func (m *MaybeType) Unset() {
	m.IsSet = false
}
