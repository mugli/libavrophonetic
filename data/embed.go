package data

import "embed"

//go:embed *.gob
var DataFiles embed.FS
