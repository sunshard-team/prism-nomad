package main

import (
	"prism/cmd"
	"prism/internal/service"
	"prism/internal/service/builder"
	"prism/internal/service/parser"
)

func main() {
	p := parser.NewParser()
	b := builder.NewStructureBuilder()
	s := service.NewService(p, b)

	cmd.Execute(s)
}
