package goldmark_citeproc

import (
	"github.com/yuin/goldmark/ast"
)

var Kind = ast.NewNodeKind("Citeproc")

var idx = 1

type Node struct {
	ast.BaseInline

	Key string
}

var _ ast.Node = (*Node)(nil)

// Kind reports the kind of this node.
func (n *Node) Kind() ast.NodeKind {
	return Kind
}

// Dump dumps the Node to stdout.
func (n *Node) Dump(src []byte, level int) {
	ast.DumpHelper(n, src, level, map[string]string{
		"Citekey": n.Key,
	}, nil)
}
