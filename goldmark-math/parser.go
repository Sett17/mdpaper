package goldmark_math

import (
	"bytes"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"strings"
)

type mathBlockParser struct {
}

func (m *mathBlockParser) Trigger() []byte {
	return []byte{'$'}
}

func (m *mathBlockParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	lineBytes, _ := reader.PeekLine()
	line := strings.TrimSpace(string(lineBytes))
	if strings.HasPrefix(line, "$$") {
		options := strings.TrimSpace(strings.TrimLeft(line, "$"))
		node := NewMathBlock()
		node.Options = options
		return node, parser.NoChildren
	}
	return nil, parser.NoChildren
}

func (m *mathBlockParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	lineBytes, segment := reader.PeekLine()
	line := strings.TrimSpace(string(lineBytes))
	if line == "$$" {
		reader.Advance(segment.Len())
		return parser.Close
	}
	node.Lines().Append(segment)
	return parser.Continue | parser.NoChildren
}

func (m *mathBlockParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	lines := node.Lines()
	var buf bytes.Buffer
	for i := 0; i < lines.Len(); i++ {
		segment := lines.At(i)
		buf.Write(segment.Value(reader.Source()))
	}
}

func (m *mathBlockParser) CanInterruptParagraph() bool {
	return false
}

func (m *mathBlockParser) CanAcceptIndentedLine() bool {
	return true
}

var defaultMathBlockParser = &mathBlockParser{}

func NewMathBlockParser() parser.BlockParser {
	return defaultMathBlockParser
}
