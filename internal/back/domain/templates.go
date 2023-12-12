package domain

import "embed"

//go:embed templates/badge-long.svg.gotmpl
//go:embed templates/badge-short.svg.gotmpl
var templates embed.FS
