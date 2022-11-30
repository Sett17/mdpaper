package goldmark_cite

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Parser struct {
	Indices *map[string]int
}

var _ parser.InlineParser = (*Parser)(nil)

var (
	_open  = []byte("[@")
	_close = []byte("]")
)

// Trigger returns characters that trigger this parser.
func (p *Parser) Trigger() []byte {
	return []byte{'['}
}

func (p *Parser) Parse(_ ast.Node, block text.Reader, _ parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if !bytes.HasPrefix(line, _open) {
		return nil
	}

	stop := bytes.Index(line, _close)
	if stop < 0 {
		return nil // must close on the same ine
	}
	seg = text.NewSegment(seg.Start+2, seg.Start+stop)

	n := &Node{Key: string(block.Value(seg))}
	if (*p.Indices)[n.Key] == 0 {
		(*p.Indices)[n.Key] = idx
		idx++
	}

	if len(n.Key) == 0 || seg.Len() == 0 {
		return nil // key and label must not be empty
	}

	n.AppendChild(n, ast.NewTextSegment(seg))
	block.Advance(stop + 1)
	return n
}
