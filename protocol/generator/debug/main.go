package main

import (
	"github.com/seoester/adcl/protocol/generator"
)

func main() {
	g := generator.NewStructGenerator(&resCommand)
	g.Generate()
}
