// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
