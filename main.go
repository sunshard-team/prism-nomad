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
	bb := builder.NewBlockBuilder()
	sb := builder.NewStructureBuilder(*bb)
	o := output.NewOutput()
	c := builder.NewChanges()
	s := service.NewService(p, bb, sb, c, o)

	cmd.Execute(s)
}
