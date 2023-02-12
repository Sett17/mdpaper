package goldmark_cite

import (
	"github.com/yuin/goldmark/ast"
	"strings"
)

var Kind = ast.NewNodeKind("Citeproc")

var idx = 1

type Citation struct {
	ast.BaseInline
	Keys []string
}

var _ ast.Node = (*Citation)(nil)

// Kind reports the kind of this node.
func (n *Citation) Kind() ast.NodeKind {
	return Kind
}

// Dump dumps the Citation to stdout.
func (n *Citation) Dump(src []byte, level int) {
	ast.DumpHelper(n, src, level, map[string]string{
		"Citekeys": strings.Join(n.Keys, ", "),
	}, nil)
}
