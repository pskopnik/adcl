package main

import (
	"github.com/seoester/adcl/protocol/generator"
)

var sidCommand = generator.Message{
	Command: "SID",
	PositionalParams: []*generator.Param{
		&generator.Param{
			Mode:     generator.ParamModePositional,
			Name:     "SID",
			Type:     "base32",
			Required: true,
		},
	},
}

var resCommand = generator.Message{
	Command: "RES",
	NamedParams: []*generator.Param{
		&generator.Param{
			Mode:     generator.ParamModeNamed,
			Name:     "FN",
			Type:     "string",
			Required: true,
		},
		&generator.Param{
			Mode:     generator.ParamModeNamed,
			Name:     "SI",
			Type:     "int",
			Required: true,
		},
		&generator.Param{
			Mode:     generator.ParamModeNamed,
			Name:     "SL",
			Type:     "int",
			Required: false,
		},
		&generator.Param{
			Mode:     generator.ParamModeNamed,
			Name:     "TO",
			Type:     "string",
			Required: true,
		},
		&generator.Param{
			Mode:     generator.ParamModeNamed,
			Name:     "TR",
			Type:     "base32",
			Required: false,
		},
		&generator.Param{
			Mode:     generator.ParamModeNamed,
			Name:     "TD",
			Type:     "int",
			Required: false,
		},
	},
	Flags: []*generator.Flag{
		&generator.Flag{
			Comment: "FI, FO, DA; EXT § 3.27 ASCH - Extended searching capability (EXT v1.0.8)",
		},
	},
}
