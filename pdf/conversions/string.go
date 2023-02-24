package conversions

import (
	"github.com/sett17/mdpaper/v2/globals"
	goldmark_cite "github.com/sett17/mdpaper/v2/goldmark-cite"
	goldmark_figref "github.com/sett17/mdpaper/v2/goldmark-figref"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"strings"
)

func String(s string) *spec.Addable {
	p := goldmark.New(
		goldmark.WithExtensions(
			&goldmark_cite.CitationExtension{},
			&goldmark_figref.FigRefExtension{},
		),
		goldmark.WithParserOptions()).Parser()
	buf := []byte(strings.TrimSpace(s))

	prevFile := globals.File
	globals.File = buf
	parsed := p.Parse(text.NewReader(buf))
	ret := Paragraph(parsed.FirstChild().(*ast.Paragraph), false)
	globals.File = prevFile

	return ret
}
