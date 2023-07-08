package templates

import "embed"

//go:embed *.tmpl
var TemplateFile embed.FS
