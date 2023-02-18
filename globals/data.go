package globals

import "embed"

//go:embed fonts/*
var Fonts embed.FS

var File []byte
