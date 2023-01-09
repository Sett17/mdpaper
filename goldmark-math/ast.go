package goldmark_math

import (
	"bytes"
	"github.com/yuin/goldmark/ast"
)

type MathBlock struct {
	ast.BaseBlock
	Options string
}

var KindMathBlock = ast.NewNodeKind("MathBlock")

func NewMathBlock() *MathBlock {
	return &MathBlock{}
}

func (n *MathBlock) Dump(source []byte, level int) {
	m := map[string]string{}
	if n.Options != "" {
		m["options"] = n.Options
	}
	ast.DumpHelper(n, source, level, m, nil)
}

func (n *MathBlock) Kind() ast.NodeKind {
	return KindMathBlock
}

func (n *MathBlock) IsRaw() bool {
	return true
}

func (n *MathBlock) Text(source []byte) []byte {
	buf := bytes.Buffer{}
	for i := 0; i < n.Lines().Len(); i++ {
		line := n.Lines().At(i)
		buf.Write(line.Value(source))
	}
	return buf.Bytes()
}
