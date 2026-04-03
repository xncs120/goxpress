package views

import (
	"embed"
)

//go:embed *.html
var HtmlFiles embed.FS

//go:embed docs/*.yaml
var YamlFiles []byte
