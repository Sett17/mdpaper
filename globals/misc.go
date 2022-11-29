package globals

import (
	"embed"
	"fmt"
	"github.com/nickng/bibtex"
)

func InToPt(in float64) float64 {
	return in * 72.0
}

func MmToPt(Mm float64) float64 {
	return InToPt(Mm / 25.4)
}

//go:embed fonts/*
var Fonts embed.FS

var File []byte

var Bibs = make(map[string]*bibtex.BibEntry)
var BibIndices = make(map[string]int)

func IEEE(entry *bibtex.BibEntry) string {
	return fmt.Sprintf(" \"%s\" %s, %s", entry.Fields["title"], entry.Fields["url"], entry.Fields["urldate"])
}
