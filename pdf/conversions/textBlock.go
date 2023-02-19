package conversions

import (
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/goldmark-cite"
	goldmark_figref "github.com/sett17/mdpaper/v2/goldmark-figref"
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
)

func TextBlock(t *ast.TextBlock) *spec.Addable {
	if string(t.Text(globals.File)) == "\\fill" {
		var a spec.Addable = &spacing.Filler{}
		return &a
	}
	txt := spec.Text{
		FontSize:   globals.Cfg.Text.FontSize,
		LineHeight: globals.Cfg.Text.LineHeight,
	}
	for n := t.FirstChild(); n != nil; n = n.NextSibling() {
		switch n.Kind() {
		case ast.KindText:
			seg := Text(n.(*ast.Text))
			txt.Add(&seg)
		case ast.KindCodeSpan:
			seg := CodeSpan(n.(*ast.CodeSpan))
			txt.Add(&seg)
		case ast.KindEmphasis:
			seg := Emphasis(n.(*ast.Emphasis))
			txt.Add(&seg)
		case goldmark_cite.Kind:
			seg := Citation(n.(*goldmark_cite.Citation))
			txt.Add(&seg)
		case goldmark_figref.Kind:
			seg := FigRef(n.(*goldmark_figref.FigRef))
			txt.Add(&seg)
		default:
			continue
		}
	}
	var a spec.Addable = &txt
	return &a
}
