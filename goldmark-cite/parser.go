package goldmark_cite

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// NewCitationParser returns a new inline parser for citations.
func NewCitationParser() parser.InlineParser {
	return &citationParser{}
}

type citationParser struct{}

// Trigger returns characters that trigger this parser.
func (p *citationParser) Trigger() []byte {
	return []byte{'['}
}

func (p *citationParser) Parse(_ ast.Node, block text.Reader, _ parser.Context) ast.Node {
	var keys []string

	for block.Peek() == '[' {
		var buf []byte
		block.Advance(1)
		if block.Peek() != '@' {
			break
		}
		block.Advance(1)

		for block.Peek() != ']' && block.Peek() != 0 {
			buf = append(buf, block.Peek())
			block.Advance(1)
		}
		if block.Peek() == 0 {
			break
		}
		block.Advance(1)

		_, _, ok := block.SkipSpaces()
		if !ok {
			keys = append(keys, string(buf))
			break
		}

		keys = append(keys, string(buf))
	}

	if len(keys) == 0 {
		return nil
	}

	return &Citation{
		Keys: keys,
	}
}

//func (p *citationParser) Parse(_ ast.Node, block text.Reader, _ parser.Context) ast.Node {
//	line, seg := block.PeekLine()
//	if !bytes.HasPrefix(line, []byte("[@")) {
//		return nil
//	}
//
//	stop := bytes.Index(line, []byte("]"))
//	if stop < 0 {
//		return nil // must close on the same ine
//	}
//	seg = text.NewSegment(seg.Start+2, seg.Start+stop)
//
//	n := &Citation{
//		Keys: []string{string(block.Value(seg))},
//	}
//	//if (*p.Indices)[n.Key] == 0 {
//	//	(*p.Indices)[n.Key] = idx
//	//	idx++
//	//}
//
//	if len(n.Keys) == 0 || seg.Len() == 0 {
//		return nil // key and label must not be empty
//	}
//
//	n.AppendChild(n, ast.NewTextSegment(seg))
//	block.Advance(stop + 1)
//	return n
//}
