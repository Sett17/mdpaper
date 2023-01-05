package conversions

import (
	"bytes"
	"fmt"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/goldmark-cite"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spacing"
	"github.com/sett17/mdpaper/pdf/spec"
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
			seg := CiteProc(n.(*goldmark_cite.Node))
			para.Add(&seg)
		default:
			continue
		}
		//if first {
		//	t.NewParagraph = true
		//	first = false
		//}
	}
	var a spec.Addable = &para
	return &a
}

func Text(text *ast.Text) spec.Segment {
	suffix := ""
	if text.SoftLineBreak() {
		suffix = " "
	}
	return spec.Segment{
		Content: string(text.Text(globals.File)) + suffix,
		Font:    spec.SerifRegular,
	}
}

func CodeSpan(span *ast.CodeSpan) spec.Segment {
	buf := bytes.Buffer{}
	for c := span.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindText && c.IsRaw() {
			buf.Write(c.(*ast.Text).Text(globals.File))
		}
	}
	return spec.Segment{
		Content: buf.String(),
		Font:    spec.Monospace,
	}
}

func Emphasis(span *ast.Emphasis) spec.Segment {
	buf := bytes.Buffer{}
	// TODO maybe? anything inside an emphasis that isnt text is ignored and text can not be parsed further...
	for c := span.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindText {
			buf.Write(c.(*ast.Text).Text(globals.File))
		}
	}
	t := spec.Segment{
		Content: buf.String(),
	}
	switch span.Level {
	case 1:
		t.Font = spec.SerifItalic
	case 2:
		t.Font = spec.SerifBold
	}
	return t
}

func CiteProc(cite *goldmark_cite.Node) spec.Segment {
	return spec.Segment{
		Content: fmt.Sprintf("[%d]", globals.BibIndices[cite.Key]),
		Font:    spec.SerifRegular,
	}
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
