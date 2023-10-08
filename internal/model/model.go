package model

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

type BuildStructure struct {
	Config       ConfigBlock
	FilesDirPath string
}

type ConfigParameter struct {
	ProjectDir     string
	ProjectDirPath string
	Namespace      string
	Release        string
	Files          []string
}

type Changes struct {
	Release   string
	Namespace string
	Files     []TemplateBlock
	Chart     ConfigBlock
}

type BlockChanges struct {
	Release   string
	Namespace string
	File      TemplateBlock
	Chart     ConfigBlock
}
