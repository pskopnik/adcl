package main

import (
	"github.com/seoester/adcl/protocol/generator"
)

func main() {
	g := generator.NewStructGenerator(&resCommand)
	err := g.Generate()
	if err != nil {
		panic(err)
	}
}
