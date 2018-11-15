package generator

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// Constants used by mapper specifications.
const (
	encodingPackage = "github.com/seoester/adcl/protocol/encoding"
	maybePackage    = "github.com/seoester/adcl/protocol/maybe"
)

// BasicMapper is a mapper performing straight-forward interpretation for
// all common types.
//
// Types supported:
//     int
//     string
//     float
//     base32
//     ip
var BasicMapper = &Mapper{
	Name: "basic",
	ComposeFieldInfoFunc: func(ctx *Context) *FieldInfo {
		return &FieldInfo{
			FieldName:          jen.Id(ctx.Param.Name),
			FieldType:          basicGolangTypeFromParam(ctx.Param),
			FieldIsMaybe:       !ctx.Param.Required,
			StrFieldName:       jen.Id(toLowerCamelCase(ctx.Param.Name) + "Str"),
			StrIsSingular:      true,
			Multiplicity:       MultiplicityStatic,
			StaticMultiplicity: 1,
		}
	},
	Parser: ParserSpec{
		PositionalParserSpec{
			ModeParserSpecBase: ModeParserSpecBase{
				Available: true,
			},
			ProcessFieldValueFunc: func(ctx *RenderingContext, value jen.Code) jen.Code {
				if ctx.FieldInfo.FieldIsMaybe {
					return jen.Add(ctx.ContentVar).Dot("").Add(ctx.FieldInfo.FieldName).Dot("").Id("Set").Call(jen.Add(value))
				} else {
					return jen.Add(ctx.ContentVar).Dot("").Add(ctx.FieldInfo.FieldName).Op("=").Add(value).Call()
				}
			},
		},
		NamedParserSpec{
			ModeParserSpecBase: ModeParserSpecBase{
				Available: true,
			},
			ParamNameFunc: func(ctx *Context) string {
				return flagNameFromParam(ctx.Param)
			},
			ProcessFieldValueFunc: func(ctx *RenderingContext, value jen.Code) jen.Code {
				if ctx.FieldInfo.FieldIsMaybe {
					return jen.Add(ctx.ContentVar).Dot("").Add(ctx.FieldInfo.FieldName).Dot("").Id("Set").Call(jen.Add(value))
				} else {
					return jen.Add(ctx.ContentVar).Dot("").Add(ctx.FieldInfo.FieldName).Op("=").Add(value).Call()
				}
			},
		},
	},
}

func flagNameFromParam(param *Param) string {
	if len(param.FlagName) > 0 {
		return param.FlagName
	} else {
		return param.Name
	}
}

func basicGolangTypeFromParam(param *Param) jen.Code {
	switch param.Type {
	case "int":
		if param.Required {
			return jen.Int()
		} else {
			return jen.Qual(maybePackage, "Int")
		}
	case "float":
		if param.Required {
			return jen.Float64()
		} else {
			return jen.Qual(maybePackage, "Float64")
		}
	case "string":
		if param.Required {
			return jen.String()
		} else {
			return jen.Qual(maybePackage, "String")
		}
	case "base32":
		if param.Required {
			return jen.Qual(encodingPackage, "*Base32Value")
		} else {
			return jen.Qual(maybePackage, "Base32Value")
		}
	case "ip":
		if param.Required {
			return jen.Qual("net", "IP")
		} else {
			return jen.Qual(maybePackage, "IP")
		}
	default:
		panic(fmt.Sprintf("Parameter type %s not known to basic mapper"))
	}
}

func toLowerCamelCase(name string) string {
	if strings.ToUpper(name) == name {
		return strings.ToLower(name)
	} else {
		return strings.ToLower(name[0:1]) + name[1:]
	}
}
