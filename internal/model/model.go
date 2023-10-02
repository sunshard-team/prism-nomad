package model

// Yaml configuration block.
// Any structure in a file configuration that is not a variable
// and contains variables and block structures.
type ConfigBlock struct {
	Name      string                   // "job", "group", "task" etc.
	Parameter []map[string]interface{} // parameter list
	Block     []ConfigBlock            // list of configuration blocks
}

// Structure for creating a nomad configuration template.
type TemplateBlock struct {
	BlockName string                   // "job", "group", "task" etc.
	Name      string                   // job "name", group "name" etc.
	Parameter []map[string]interface{} // parameter list
	Block     []TemplateBlock          // list of configuration blocks
}

type ConfigParameter struct {
	ProjectDir     string
	ProjectDirPath string
	Namespace      string
	Release        string
	Files          []string
}
