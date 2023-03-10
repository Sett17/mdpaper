package goldmark_figref

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// NewFigRefParser returns a new inline parser for citations.
func NewFigRefParser() parser.InlineParser {
	return &citationParser{}
}

type citationParser struct{}

// Trigger returns characters that trigger this parser.
func (p *citationParser) Trigger() []byte {
	return []byte("[")
}

func (p *citationParser) Parse(_ ast.Node, block text.Reader, _ parser.Context) ast.Node {
	var key string

	if block.Peek() == '[' {
		var buf []byte
		block.Advance(1)
		if block.Peek() == '!' {
			block.Advance(1)
			for block.Peek() != ']' && block.Peek() != 0 {
				buf = append(buf, block.Peek())
				block.Advance(1)
			}
			block.Advance(1)
		} else {
			return nil
		}
		key = string(buf)
	}

	return &FigRef{
		FigureKey: key,
	}
}
