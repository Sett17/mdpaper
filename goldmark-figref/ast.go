package goldmark_figref

import (
	"github.com/yuin/goldmark/ast"
)

var Kind = ast.NewNodeKind("FigRef")

type FigRef struct {
	ast.BaseInline
	FigureKey string
}

var _ ast.Node = (*FigRef)(nil)

// Kind reports the kind of this node.
func (n *FigRef) Kind() ast.NodeKind {
	return Kind
}

// Dump dumps the FigRef to stdout.
func (n *FigRef) Dump(src []byte, level int) {
	ast.DumpHelper(n, src, level, map[string]string{
		"FigureKey": n.FigureKey,
	}, nil)
}
