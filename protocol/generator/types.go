package generator

import (
	"errors"
)

var intTypeSpec = &TypeSpec{
	Name:          "int",
	Mappers:       []*Mapper{BasicMapper},
	DefaultMapper: BasicMapper,
}

var floatTypeSpec = &TypeSpec{
	Name:          "float",
	Mappers:       []*Mapper{BasicMapper},
	DefaultMapper: BasicMapper,
}

var stringTypeSpec = &TypeSpec{
	Name:          "string",
	Mappers:       []*Mapper{BasicMapper},
	DefaultMapper: BasicMapper,
}

var base32TypeSpec = &TypeSpec{
	Name:          "base32",
	Mappers:       []*Mapper{BasicMapper},
	DefaultMapper: BasicMapper,
}

var ipTypeSpec = &TypeSpec{
	Name:          "ip",
	Mappers:       []*Mapper{BasicMapper},
	DefaultMapper: BasicMapper,
}

// Error variables related to type specifications and type and mapper
// resolution.
var (
	ErrUnknownTypeName   = errors.New("unknown type, cannot find type by name")
	ErrNoDefaultMapper   = errors.New("no default mapper configured by type spec, one needs to be selected manually")
	ErrUnknownMapperName = errors.New("unknown mapper name (for type), cannot find mapper by name")
)

func TypeSpecFromName(name string) (*TypeSpec, error) {
	switch name {
	case "int":
		return intTypeSpec, nil
	case "float":
		return intTypeSpec, nil
	case "string":
		return stringTypeSpec, nil
	case "base32":
		return base32TypeSpec, nil
	case "ip":
		return ipTypeSpec, nil
	default:
		return nil, ErrUnknownTypeName
	}
}

func ResolveMapperFromParam(param *Param) (*Mapper, error) {
	typeSpec, err := TypeSpecFromName(param.Type)
	if err != nil {
		return nil, err
	}

	if len(param.Mapper) == 0 {
		if typeSpec.DefaultMapper == nil {
			return nil, ErrNoDefaultMapper
		}

		return typeSpec.DefaultMapper, nil
	}

	for _, mapper := range typeSpec.Mappers {
		if mapper.Name == param.Mapper {
			return mapper, nil
		}
	}

	return nil, ErrUnknownMapperName
}
