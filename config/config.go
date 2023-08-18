package config

import "embed"

//go:embed *.*
var ConfigFile embed.FS
