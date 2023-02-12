package conversions

import (
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/goldmark-cite"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
)

func Paragraph(p *ast.Paragraph, centered bool) *spec.Addable {
	if string(p.Text(globals.File)) == "\\fill" {
		var a spec.Addable = &spacing.Filler{}
		return &a
	}
	para := elements.Paragraph{
		Text: spec.Text{
			FontSize:   globals.Cfg.Text.FontSize,
			LineHeight: globals.Cfg.Text.LineHeight,
		},
		Centered: centered,
	}
	for n := p.FirstChild(); n != nil; n = n.NextSibling() {
		//var i spec.ImageObject
		switch n.Kind() {
		case ast.KindText:
			seg := Text(n.(*ast.Text))
			para.Add(&seg)
		case ast.KindCodeSpan:
			seg := CodeSpan(n.(*ast.CodeSpan))
			para.Add(&seg)
		case ast.KindEmphasis:
			seg := Emphasis(n.(*ast.Emphasis))
			para.Add(&seg)
		case goldmark_cite.Kind:
			seg := Citation(n.(*goldmark_cite.Citation))
			para.Add(&seg)
		default:
			continue
		}
	}
	var a spec.Addable = &para
	return &a
}

func Blockquote(bq *ast.Blockquote) (ret []*spec.Addable) {
	for c := bq.FirstChild(); c != nil; c = c.NextSibling() {
		switch c.Kind() {
		case ast.KindParagraph:
			ret = append(ret, Paragraph(c.(*ast.Paragraph), true))
		}
	}
	return
}
