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
	if entry == nil {
		return ""
	}
	title := strings.ReplaceAll(entry.Fields["title"].String(), "{", "")
	title = strings.ReplaceAll(title, "}", "")
	url := ""
	if entry.Fields["url"] != nil {
		url = entry.Fields["url"].String()
	}
	date := ""
	if entry.Fields["urldate"] != nil {
		date = entry.Fields["urldate"].String()
	}

	return fmt.Sprintf(" \"%s\" %s, %s", title, url, date)
}

}
