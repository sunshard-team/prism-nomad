package main

import (
	"prism/cmd"
	"prism/internal/service"
	"prism/internal/service/builder"
	"prism/internal/service/output"
	"prism/internal/service/parser"
)

func main() {
	p := parser.NewParser()
	b := builder.NewStructureBuilder()
	o := output.NewOutput()
	s := service.NewService(p, b, o)

	cmd.Execute(s)
}
