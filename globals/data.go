package globals

import (
	"embed"
	"sync"
)

//go:embed fonts/*
var Fonts embed.FS

var File []byte

var ImageSync sync.WaitGroup
