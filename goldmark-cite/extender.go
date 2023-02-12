package goldmark_cite

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

// CitationExtension is a custom goldmark extension for parsing citations in the format [@citekey].
type CitationExtension struct{}

// Extend implements the goldmark.Extender interface.
func (e *CitationExtension) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewCitationParser(), 5),
	))
}
