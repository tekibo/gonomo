package runtime

import (
	"embed"
)

//go:embed all:*
var TemplatesFS embed.FS
