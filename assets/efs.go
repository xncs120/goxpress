package assets

import (
	"embed"
)

//go:embed "statics" "templates"
var EmbeddedFiles embed.FS
