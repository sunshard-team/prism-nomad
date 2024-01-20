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

type Pack struct {
	Name          string           `yaml:"name"`
	Description   string           `yaml:"description"`
	Maintainers   []string         `yaml:"maintainers"`
	Type          string           `yaml:"type"`
	Sources       []string         `yaml:"sources"`
	DeployVersion string           `yaml:"deploy_version"`
	PackVersion   string           `yaml:"pack_version"`
	NomadVersion  string           `yaml:"nomad_version"`
	Dependencies  []PackDependency `yaml:"dependencies"`
}

type PackDependency struct {
	Name        string   `yaml:"name"`
	PackVersion string   `yaml:"pack_version"`
	Path        string   `yaml:"path"`
	Files       []string `yaml:"files"`
}

// Necessary data for building the job configuration structure.
type BuildStructure struct {
	Config       ConfigBlock
	FilesDirPath string
}

// Job deployment data.
type ConfigParameter struct {
	ProjectDirPath string
	Namespace      string
	Release        string
	Files          []string
	EnvFilePath    string
	EnvVars        map[string]string
}

type CheckNamespace struct {
	Client          *api.Client
	Namespace       string
	CreateNamespace bool
}

type Changes struct {
	Release     string
	Namespace   string
	Files       []TemplateBlock
	Pack        Pack
	EnvFilePath string
	EnvVars     map[string]string
}

type BlockChanges struct {
	Release   string
	Namespace string
	File      TemplateBlock
	Pack      Pack
}

type Deployment struct {
	Client    *api.Client
	JobName   string
	Namespace string
	Config    string
	WaitTime  int
}
