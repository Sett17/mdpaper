package pdf

import (
	"bytes"
	"github.com/yuin/goldmark/ast"
	"mdpaper/globals"
	"mdpaper/pdf/spec"
	"strings"
)

func ConvertHeading(h *ast.Heading) *spec.Addable {
	buf := strings.Builder{}
	buf.Write(h.Text(globals.File))
	seg := spec.Segment{
		Content: buf.String(),
		Font:    &spec.TimesBold,
	}
	fs := globals.Cfg.FontSize
	if h.Level <= 2 {
		fs = int(float64(fs) * 1.2)
	}
	para := Heading{
		Text: spec.Text{
			FontSize:   fs,
			LineHeight: globals.Cfg.LineHeight * 1.5,
			SingleLine: true,
		},
		Level: h.Level,
	}
	para.Add(&seg)
	var out spec.Addable = &para
	return &out
}

func ConvertParagraph(p *ast.Paragraph) *spec.Addable {
	if string(p.Text(globals.File)) == "\\fill" {
		var a spec.Addable
		a = &Filler{}
		return &a
	}
	para := Paragraph{
		Text: spec.Text{
			FontSize:   globals.Cfg.FontSize,
			LineHeight: globals.Cfg.LineHeight,
		},
	}
	for n := p.FirstChild(); n != nil; n = n.NextSibling() {
		//var i spec.ImageObject
		switch n.Kind() {
		case ast.KindText:
			seg := ConvertText(n.(*ast.Text))
			para.Add(&seg)
		case ast.KindCodeSpan:
			seg := ConvertCodeSpan(n.(*ast.CodeSpan))
			para.Add(&seg)
		case ast.KindEmphasis:
			seg := ConvertEmphasis(n.(*ast.Emphasis))
			para.Add(&seg)
		default:
			continue
			//case ast.KindImage:
			//	//TODO support images
			//	//i = ConvertImage(n.(*ast.Image))
		}
		//if first {
		//	t.NewParagraph = true
		//	first = false
		//}
	}
	var a spec.Addable
	a = &para
	return &a
}

func ConvertText(text *ast.Text) spec.Segment {
	return spec.Segment{
		Content: string(text.Text(globals.File)),
		Font:    &spec.TimesRegular,
	}
}

func ConvertCodeSpan(span *ast.CodeSpan) spec.Segment {
	buf := bytes.Buffer{}
	for c := span.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindText && c.IsRaw() {
			buf.Write(c.(*ast.Text).Text(globals.File))
		}
	}
	return spec.Segment{
		Content: buf.String(),
		Font:    &spec.CourierMono,
	}
}

func ConvertEmphasis(span *ast.Emphasis) spec.Segment {
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
		t.Font = &spec.TimesItalic
	case 2:
		t.Font = &spec.TimesBold
	}
	return t
}
