package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Multiplicity int

const (
	MultiplicityStatic Multiplicity = iota
	MultiplicityDynamic
)

func (m Multiplicity) String() string {
	switch m {
	case MultiplicityStatic:
		return "static"
	case MultiplicityDynamic:
		return "dynamic"
	default:
		return fmt.Sprintf("Multiplicity(%d)", m)
	}
}

type FieldInfo struct {
	FieldName    jen.Code
	FieldType    jen.Code
	FieldIsMaybe bool
	StrFieldName jen.Code
	// StrIsSingular is true if the type of the Str field is string, not
	// []string. This should always be the case when the field has a static
	// multiplicity of 1.
	StrIsSingular bool
	// Multiplicity is the multiplicity type of the field.
	Multiplicity Multiplicity
	// StaticMultiplicity is the static multiplicity of the field.
	// This is only set if Multiplicity is equal to MultiplicityStatic.
	StaticMultiplicity int
}

// DynamicMultiplicity returns code which evaluates to the multiplicity of the field.
func (p *FieldInfo) DynamicMultiplicity(ctx *RenderingContext) jen.Code {
	switch p.Multiplicity {
	case MultiplicityStatic:
		return jen.Lit(p.StaticMultiplicity)
	case MultiplicityDynamic:
		return jen.Len(jen.Add(ctx.ContentVar).Dot("").Add(ctx.FieldInfo.StrFieldName))
	default:
		panic(
			fmt.Sprintf("Unknown multiplicity type Multiplicity(%d)", p.Multiplicity),
		)
	}
}

type Context struct {
	Param  *Param
	Mapper *Mapper
	Type   *TypeSpec
}

type RenderingContext struct {
	Context

	// ContentVar contains code to access the content struct the field is
	// contained in.
	ContentVar jen.Code
	// FieldInfo contains a FieldInfo instance created prior by the
	// ComposeFieldInfoFunc of the Mapper.
	FieldInfo *FieldInfo
	// ErrorVar contains the code to access an existing error variable. If
	// none exists, this field is nil.
	ErrorVar jen.Code
	// ErrorCheckFunc is a function generating code to check an error
	// variable. Code to access the error variable has to be passed in.
	ErrorCheckFunc func(errorVar jen.Code) jen.Code
}

func (r *RenderingContext) ErrorCheck(errorVar jen.Code) jen.Code {
	if r.ErrorCheckFunc == nil {
		panic(
			fmt.Sprintf("No ErrorCheckFunc defined"),
		)
	}

	return r.ErrorCheckFunc(errorVar)
}

type TypeSpec struct {
	Name          string
	Mappers       []*Mapper
	DefaultMapper *Mapper
}

type Mapper struct {
	Name string

	ComposeFieldInfoFunc func(ctx *Context) *FieldInfo

	Parser  ParserSpec
	Builder BuilderSpec
}

func (m Mapper) ComposeFieldInfo(ctx *Context) *FieldInfo {
	return m.ComposeFieldInfoFunc(ctx)
}

type ParserSpec struct {
	Positional PositionalParserSpec
	Named      NamedParserSpec
}

type ModeParserSpecBase struct {
	Available           bool
	InitialiseFieldFunc func(ctx *RenderingContext) jen.Code
	FinaliseFieldFunc   func(ctx *RenderingContext) jen.Code
}

func (m *ModeParserSpecBase) checkAvailable(ctx *Context) {
	if !m.Available {
		panic(
			fmt.Sprintf(
				"The mapper %s has no spec available for parsing in mode %s",
				ctx.Mapper.Name,
				ctx.Param.Mode,
			),
		)
	}
}

func (m ModeParserSpecBase) InitialiseField(ctx *RenderingContext) jen.Code {
	m.checkAvailable(&ctx.Context)

	if m.InitialiseFieldFunc == nil {
		return nil
	}

	return m.InitialiseFieldFunc(ctx)
}

func (m ModeParserSpecBase) FinaliseField(ctx *RenderingContext) jen.Code {
	m.checkAvailable(&ctx.Context)

	if m.FinaliseFieldFunc == nil {
		return nil
	}

	return m.FinaliseFieldFunc(ctx)
}

type PositionalParserSpec struct {
	ModeParserSpecBase

	DynamicProcessCondFunc        func(ctx *RenderingContext, value jen.Code) jen.Code
	DynamicProcessCondPrepareFunc func(ctx *RenderingContext, value jen.Code) jen.Code
	DynamicProcessCondCleanupFunc func(ctx *RenderingContext, value jen.Code) jen.Code

	ProcessFieldValueFunc func(ctx *RenderingContext, value jen.Code) jen.Code
}

func (p PositionalParserSpec) DynamicProcessCond(ctx *RenderingContext, value jen.Code) jen.Code {
	p.checkAvailable(&ctx.Context)

	if p.DynamicProcessCondFunc == nil {
		panic(
			fmt.Sprintf(
				"The mapper %s has no spec available for parsing with dynamic multiplicity",
				ctx.Mapper.Name,
			),
		)
	}

	return p.DynamicProcessCondFunc(ctx, value)
}

func (p PositionalParserSpec) DynamicProcessCondPrepare(ctx *RenderingContext, value jen.Code) jen.Code {
	p.checkAvailable(&ctx.Context)

	if p.DynamicProcessCondPrepareFunc == nil {
		return nil
	}

	return p.DynamicProcessCondPrepareFunc(ctx, value)
}

func (p PositionalParserSpec) DynamicProcessCondCleanup(ctx *RenderingContext, value jen.Code) jen.Code {
	p.checkAvailable(&ctx.Context)

	if p.DynamicProcessCondCleanupFunc == nil {
		return nil
	}

	return p.DynamicProcessCondCleanupFunc(ctx, value)
}

func (p PositionalParserSpec) ProcessFieldValue(ctx *RenderingContext, value jen.Code) jen.Code {
	p.checkAvailable(&ctx.Context)

	if p.ProcessFieldValueFunc == nil {
		return nil
	}

	return p.ProcessFieldValueFunc(ctx, value)
}

type NamedParserSpec struct {
	ModeParserSpecBase

	ParamNameFunc func(ctx *Context) string

	ProcessFieldValueFunc func(ctx *RenderingContext, value jen.Code) jen.Code
}

func (n NamedParserSpec) ParamName(ctx *Context) string {
	n.checkAvailable(ctx)

	if n.ParamNameFunc == nil {
		panic(
			fmt.Sprintf(
				"The mapper %s has an incomplete spec for parsing in named mode: ParamNameFunc missing",
				ctx.Mapper.Name,
			),
		)
	}

	return n.ParamNameFunc(ctx)
}

func (n NamedParserSpec) ProcessFieldValue(ctx *RenderingContext, value jen.Code) jen.Code {
	n.checkAvailable(&ctx.Context)

	if n.ProcessFieldValueFunc == nil {
		return nil
	}

	return n.ProcessFieldValueFunc(ctx, value)
}

type BuilderSpec struct{}
