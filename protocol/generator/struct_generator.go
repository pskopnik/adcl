package generator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
)

type paramInfo struct {
	Param     *Param
	Mapper    *Mapper
	Type      *TypeSpec
	FieldInfo *FieldInfo
}

type StructGenerator struct {
	message *Message

	typeName     string
	typeLetter   string
	flagTypeName string

	positionalParams []paramInfo
	namedParams      []paramInfo
}

func NewStructGenerator(message *Message) *StructGenerator {
	return &StructGenerator{
		message: message,
	}
}

func (s *StructGenerator) Generate() error {
	err := s.prepare()
	if err != nil {
		return err
	}

	file := s.generateFile()

	buf := bytes.NewBuffer(nil)
	file.Render(buf)

	fmt.Println(buf.String())

	return nil
}

func (s *StructGenerator) prepare() error {
	var err error

	s.positionalParams, err = s.prepareParams(s.message.PositionalParams)
	if err != nil {
		return err
	}

	s.namedParams, err = s.prepareParams(s.message.NamedParams)
	if err != nil {
		return err
	}

	s.typeName = s.message.Command + "Content"
	s.typeLetter = strings.ToLower(s.typeName[0:1])
	s.flagTypeName = s.message.Command + "Flag"

	return nil
}

func (s *StructGenerator) prepareParams(params []*Param) ([]paramInfo, error) {
	paramInfos := make([]paramInfo, 0, len(params))

	for _, param := range params {
		typeSpec, err := TypeSpecFromName(param.Type)
		if err != nil {
			return nil, errors.Wrapf(err, "type resolution failed for type name %s specified by "+
				"param %s of message %s", param.Name, s.message.Command)
		}
		mapper, err := ResolveMapperFromParam(param)
		if err != nil {
			return nil, errors.Wrapf(err, "mapper resolution failed for param %s of message %s "+
				"with type %s", param.Name, s.message.Command)
		}

		ctx := Context{
			Param:  param,
			Mapper: mapper,
			Type:   typeSpec,
		}

		paramInfos = append(paramInfos, paramInfo{
			Param:     param,
			Mapper:    mapper,
			Type:      typeSpec,
			FieldInfo: mapper.ComposeFieldInfo(&ctx),
		})
	}

	return paramInfos, nil
}

func (s *StructGenerator) generateFile() *jen.File {
	file := jen.NewFile("message")

	file.Comment("Code generated by adcl/protocol/generator. DO NOT EDIT.")

	file.Line()

	file.Type().Id(s.flagTypeName).String()

	if len(s.namedParams) > 0 {
		file.Const().
			DefsFunc(s.generateFlagConstants)
	}

	file.Var().Id("_").Id("ParamAccessor").Op("=").Op("&").Id(s.typeName).Values()

	file.Type().Id(s.typeName).StructFunc(s.generateStructFields)

	file.Func().Params(jen.Id(s.typeLetter).Op("*").Id(s.typeName)).
		Id("Positional").Params().Index().String().
		BlockFunc(s.generatePositional)

	file.Line()

	file.Func().Params(jen.Id(s.typeLetter).Op("*").Id(s.typeName)).
		Id("PosLen").Params().Int().
		BlockFunc(s.generatePosLen)

	file.Line()

	file.Func().Params(jen.Id(s.typeLetter).Op("*").Id(s.typeName)).
		Id("PosAt").Params(jen.Id("i").Int()).String().
		BlockFunc(s.generatePosAt)

	file.Line()

	file.Func().Params(jen.Id(s.typeLetter).Op("*").Id(s.typeName)).
		Id("Named").Params().Map(jen.String()).String().
		BlockFunc(s.generateNamed)

	file.Line()

	file.Func().Params(jen.Id(s.typeLetter).Op("*").Id(s.typeName)).
		Id("NamedGet").Params(jen.Id("key").String()).Params(jen.String(), jen.Bool()).
		BlockFunc(s.generateNamedGet)

	return file
}

func (s *StructGenerator) generateFlagConstants(group *jen.Group) {
	for ind, param := range s.namedParams {
		ctx := s.createContext(param)

		name := param.Mapper.Parser.Named.ParamName(&ctx)

		if ind == 0 {
			group.Id(s.flagTypeName + name).Id(s.flagTypeName).Op("=").Lit(name)
		} else {
			group.Id(s.flagTypeName + name).Op("=").Lit(name)
		}
	}
}

func (s *StructGenerator) generateStructFields(group *jen.Group) {
	s.generateParamsStructFields(group, s.positionalParams)
	s.generateParamsStructFields(group, s.namedParams)

	group.Id("Flags").Map(jen.String()).String()

	group.Line()

	if len(s.message.Flags) == 0 {
		group.Comment("No known additional flags.")
	} else {
		for _, flag := range s.message.Flags {
			group.Comment(flag.Comment)
		}
	}
}

func (s *StructGenerator) generateParamsStructFields(group *jen.Group, params []paramInfo) {
	for _, param := range params {
		info := param.FieldInfo

		// TODO: Format and insert comment

		group.Add(info.FieldName).Add(info.FieldType)

		strFieldType := jen.String()
		if !info.StrIsSingular {
			strFieldType = jen.Index().Add(strFieldType)
		}
		group.Add(info.StrFieldName).Add(strFieldType)

		group.Line()
	}
}

func (s *StructGenerator) generatePositional(group *jen.Group) {
	var numStatic int

	for _, param := range s.positionalParams {
		if param.FieldInfo.Multiplicity == MultiplicityStatic {
			numStatic++
		}
	}

	if numStatic == len(s.positionalParams) {
		// All params have static multiplicity, build slice literal.

		group.Return(
			jen.Index().String().ValuesFunc(func(group *jen.Group) {
				for _, param := range s.positionalParams {
					if param.FieldInfo.StrIsSingular {
						group.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName)
					} else {
						for i := 0; i < param.FieldInfo.StaticMultiplicity; i++ {
							group.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName).
								Index(jen.Lit(i))
						}
					}
				}
			}),
		)
	} else if numStatic == 0 && len(s.positionalParams) == 1 {
		// There is only a single, dynamic multiplicity param, return its str
		// field.

		group.Return(
			jen.Id(s.typeLetter).Dot("").Add(s.positionalParams[0].FieldInfo.StrFieldName),
		)
	} else {
		// Params have mixed multiplicity, build slice of positionals
		// manually.

		group.Id("positionals").Op(":=").Make(
			jen.Index().String(),
			jen.Lit(0),
			jen.Id(s.typeLetter).Dot("PosLen").Call(),
		)

		for _, param := range s.positionalParams {
			if param.FieldInfo.StrIsSingular {
				group.Id("positionals").Op("=").Append(
					jen.Id("positionals"),
					jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName),
				)
			} else if param.FieldInfo.Multiplicity == MultiplicityStatic {
				group.Id("positionals").Op("=").AppendFunc(func(group *jen.Group) {
					group.Id("positionals")
					for i := 0; i < param.FieldInfo.StaticMultiplicity; i++ {
						group.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName).
							Index(jen.Lit(i))
					}
				})
			} else {
				group.Id("positionals").Op("=").Append(
					jen.Id("positionals"),
					jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName).Op("..."),
				)
			}
		}

		group.Return(jen.Id("positionals"))
	}
}

func (s *StructGenerator) generatePosLen(group *jen.Group) {
	var staticSum int
	var stmt jen.Statement

	for _, param := range s.positionalParams {
		ctx := s.createRenderingContext(param)

		if param.FieldInfo.Multiplicity == MultiplicityStatic {
			staticSum += param.FieldInfo.StaticMultiplicity
		} else {
			stmt.Op("+").Add(param.FieldInfo.DynamicMultiplicity(&ctx))
		}
	}

	if staticSum > 0 || len(stmt) == 0 {
		group.Return(jen.Lit(staticSum).Add(&stmt))
	} else {
		// Remove first Op("+")
		stmt = stmt[1:]
		group.Return(&stmt)
	}
}

func (s *StructGenerator) generatePosAt(group *jen.Group) {
	var numStatic int

	if len(s.positionalParams) == 0 {
		group.Panic(jen.Lit("index out of range"))
		return
	}

	for _, param := range s.positionalParams {
		if param.FieldInfo.Multiplicity == MultiplicityStatic {
			numStatic++
		}
	}

	if numStatic == len(s.positionalParams) {
		// All params have static multiplicity, build switch statement.

		var runningIndex int

		group.Switch(jen.Id("i")).BlockFunc(func(group *jen.Group) {
			for _, param := range s.positionalParams {
				if param.FieldInfo.StrIsSingular {
					group.Case(jen.Lit(runningIndex)).Block(
						jen.Return(
							jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName),
						),
					)
					runningIndex++
				} else {
					for i := 0; i < param.FieldInfo.StaticMultiplicity; i++ {
						group.Case(jen.Lit(runningIndex)).Block(
							jen.Return(
								jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName).
									Index(jen.Lit(i)),
							),
						)
						runningIndex++
					}
				}
			}

			group.Default().Block(
				jen.Panic(jen.Lit("index out of range")),
			)
		})
	} else if numStatic == 0 && len(s.positionalParams) == 1 {
		// There is only a single, dynamic multiplicity param, return the ith
		// element of its str field.

		group.Return(
			jen.Id(s.typeLetter).Dot("").Add(s.positionalParams[0].FieldInfo.StrFieldName).
				Index(jen.Id("i")),
		)
	} else {
		// Params have mixed multiplicity, build conditional switch statement
		// manually.

		var runningStaticIndex int
		var runningDynamicLens []jen.Code

		runningLenStmt := func(op string, dynamics ...jen.Code) *jen.Statement {
			var stmt jen.Statement

			if runningStaticIndex > 0 || len(runningDynamicLens) == 0 {
				stmt.Lit(runningStaticIndex).Op(op)
			}

			for _, dynamic := range runningDynamicLens {
				stmt.Add(dynamic).Op(op)
			}

			for _, dynamic := range dynamics {
				stmt.Add(dynamic).Op(op)
			}

			// Remove last Op(op)
			stmt = stmt[:len(stmt)-1]

			return &stmt
		}

		group.Switch().BlockFunc(func(group *jen.Group) {
			for _, param := range s.positionalParams {
				if param.FieldInfo.StrIsSingular {
					group.Case(
						jen.Id("i").Op("==").Add(runningLenStmt("+")),
					).Block(
						jen.Return(
							jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName),
						),
					)
					runningStaticIndex++
				} else if param.FieldInfo.Multiplicity == MultiplicityStatic {
					for i := 0; i < param.FieldInfo.StaticMultiplicity; i++ {
						group.Case(
							jen.Id("i").Op("==").Add(runningLenStmt("+")),
						).Block(
							jen.Return(
								jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName).
									Index(jen.Lit(i).Op("-").Add(runningLenStmt("-"))),
							),
						)
						runningStaticIndex++
					}
				} else {
					ctx := s.createRenderingContext(param)
					dynamicLen := param.FieldInfo.DynamicMultiplicity(&ctx)

					group.Case(
						jen.Id("i").Op("<").Add(runningLenStmt("+", dynamicLen)),
					).Block(
						jen.Return(
							jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName).
								// TODO: Omit -0 from [i-0]
								Index(jen.Lit("i").Op("-").Add(runningLenStmt("-"))),
						),
					)
					runningDynamicLens = append(runningDynamicLens, dynamicLen)
				}
			}

			group.Default().Block(
				jen.Panic(jen.Lit("index out of range")),
			)
		})
	}
}

func (s *StructGenerator) generateNamed(group *jen.Group) {
	if len(s.namedParams) == 0 {
		// There are no named parameters, return the Flags map.
		group.Return(
			jen.Id(s.typeLetter).Dot("Flags"),
		)

		return
	}

	group.Id("params").Op(":=").Make(jen.Map(jen.String()).String())

	group.Line()

	group.For(
		jen.List(jen.Id("key"), jen.Id("val")).
			Op(":=").
			Range().Id(s.typeLetter).Dot("Flags"),
	).Block(
		jen.Id("params").Index(jen.Id("key")).Op("=").Id("val"),
	)

	group.Line()

	for _, param := range s.namedParams {
		strStmt := jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName)
		if !param.FieldInfo.StrIsSingular {
			strStmt.Index(jen.Lit(0))
		}

		setStmt := jen.Id("params").Index(
			jen.Add(strStmt).
				Index(
					jen.Empty(), jen.Lit(2),
				),
		).
			Op("=").
			Add(strStmt).
			Index(
				jen.Lit(2), jen.Empty(),
			)

		if param.FieldInfo.FieldIsMaybe || !param.FieldInfo.StrIsSingular {
			var condStmts []jen.Code

			if param.FieldInfo.FieldIsMaybe {
				condStmts = append(
					condStmts,
					jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.FieldName).
						Dot("IsSet"),
				)
			}
			if !param.FieldInfo.StrIsSingular {
				condStmts = append(
					condStmts,
					jen.Len(
						jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName),
					).
						Op(">").Lit(0),
				)
			}

			group.If(
				s.opJoin("&&", condStmts...)...,
			).
				Block(
					setStmt,
				)
		} else {
			group.Add(
				setStmt,
			)
		}
	}

	group.Line()

	group.Return(jen.Id("params"))
}

func (s *StructGenerator) generateNamedGet(group *jen.Group) {
	// TODO: Should the switch block be wrapped in the follwing if?
	// group.If(jen.Len(jen.Id("key")).Op("==").Lit(2)).Block()

	if len(s.namedParams) > 0 {
		group.Switch(jen.Id(s.flagTypeName).Parens(jen.Id("key"))).
			BlockFunc(func(group *jen.Group) {
				for _, param := range s.namedParams {
					ctx := s.createContext(param)
					name := param.Mapper.Parser.Named.ParamName(&ctx)

					returnStrStmt := jen.Id(s.typeLetter).Dot("").
						Add(param.FieldInfo.StrFieldName)
					if !param.FieldInfo.StrIsSingular {
						returnStrStmt.Index(jen.Lit(0))
					}
					returnStrStmt.Index(
						jen.Lit(2), jen.Empty(),
					)

					if param.FieldInfo.FieldIsMaybe || !param.FieldInfo.StrIsSingular {
						var condStmts []jen.Code

						if param.FieldInfo.FieldIsMaybe {
							condStmts = append(
								condStmts,
								jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.FieldName).
									Dot("IsSet"),
							)
						}
						if !param.FieldInfo.StrIsSingular {
							condStmts = append(
								condStmts,
								jen.Len(
									jen.Id(s.typeLetter).Dot("").Add(param.FieldInfo.StrFieldName),
								).
									Op(">").Lit(0),
							)
						}

						group.Case(jen.Id(s.flagTypeName + name)).
							If(
								s.opJoin("&&", condStmts...)...,
							).
							Block(
								jen.Return(jen.List(
									returnStrStmt,
									jen.Lit(true),
								)),
							).Else().
							Block(
								jen.Return(jen.List(
									jen.Lit(""),
									jen.Lit(false),
								)),
							)
					} else {
						group.Case(jen.Id(s.flagTypeName + name)).
							Return(jen.List(
								returnStrStmt,
								jen.Lit(true),
							))
					}

				}

			})

		group.Line()
	}

	group.
		List(
			jen.Id("key"), jen.Id("val"),
		).
		Op(":=").
		Id(s.typeLetter).Dot("Flags").Index(jen.Id("key"))

	group.Return(jen.List(
		jen.Id("key"), jen.Id("val"),
	))
}

func (s *StructGenerator) opJoin(op string, codes ...jen.Code) jen.Statement {
	var stmt jen.Statement

	for _, code := range codes {
		stmt.Add(code).Op(op)
	}

	if len(stmt) > 0 {
		// Remove last Op(op)
		stmt = stmt[:len(stmt)-1]
	}

	// TODO: Return Empty() when stmt is empty?

	return stmt
}

func (s *StructGenerator) createRenderingContext(param paramInfo) RenderingContext {
	return RenderingContext{
		Context:    s.createContext(param),
		ContentVar: jen.Id(s.typeLetter),
		FieldInfo:  param.FieldInfo,
	}
}

func (s *StructGenerator) createContext(param paramInfo) Context {
	return Context{
		Param:  param.Param,
		Mapper: param.Mapper,
		Type:   param.Type,
	}
}
