package goldmark_citeproc

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

type Extender struct {
	Indices *map[string]int
}

func (e *Extender) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(&Parser{Indices: e.Indices}, 199),
		),
	)
}
