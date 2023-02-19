package goldmark_figref

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

// FigRefExtension is a custom goldmark extension for parsing citations in the format [@citekey].
type FigRefExtension struct{}

// Extend implements the goldmark.Extender interface.
func (e *FigRefExtension) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewFigRefParser(), 5),
	))
}
