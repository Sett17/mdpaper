package conversions

import (
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
)

func Blockquote(bq *ast.Blockquote) (ret []*spec.Addable) {
	for c := bq.FirstChild(); c != nil; c = c.NextSibling() {
		switch c.Kind() {
		case ast.KindParagraph:
			ret = append(ret, Paragraph(c.(*ast.Paragraph), true))
		}
	}
	return
}
