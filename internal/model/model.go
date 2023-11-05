// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

import "github.com/hashicorp/nomad/api"

// Yaml configuration block.
// Any structure in a file configuration that is not a variable
// and contains variables and block structures.
type ConfigBlock struct {
	Type      string                   // "job", "group", "task" etc.
	Parameter []map[string]interface{} // parameter list
	Block     []ConfigBlock            // list of configuration blocks
}

// Structure for creating a nomad configuration template.
type TemplateBlock struct {
	Type      string                   // "job", "group", "task" etc.
	Label     string                   // job "name", group "name" etc.
	Parameter []map[string]interface{} // parameter list
	Block     []TemplateBlock          // list of configuration blocks
}

// Necessary data for building the job configuration structure.
type BuildStructure struct {
	Config       ConfigBlock
	FilesDirPath string
}

// Job deployment data.
type ConfigParameter struct {
	ProjectDir     string
	ProjectDirPath string
	Namespace      string
	Release        string
	Files          []string
}

type CheckNamespace struct {
	Client          *api.Client
	Namespace       string
	CreateNamespace bool
}

type Changes struct {
	Release   string
	Namespace string
	Files     []TemplateBlock
	Pack      ConfigBlock
}

type BlockChanges struct {
	Release   string
	Namespace string
	File      TemplateBlock
	Pack      ConfigBlock
}
