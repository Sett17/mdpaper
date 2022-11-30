package pdf

import (
	"bytes"
	"fmt"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/goldmark-cite"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"strconv"
	"strings"
)

func ConvertHeading(h *ast.Heading) *spec.Addable {
	buf := strings.Builder{}
	buf.Write(h.Text(globals.File))
	seg := spec.Segment{
		Content: buf.String(),
		Font:    spec.SansBold,
	}
	fs := globals.Cfg.Text.FontSize
	if h.Level <= 2 {
		fs = int(float64(fs) * 1.2)
	}
	para := Heading{
		Text: spec.Text{
			FontSize: fs,
			//LineHeight: globals.Cfg.LineHeight * 1.5,
			LineHeight: 1.0,
			Offset:     globals.MmToPt(0 + 1.5*float64(h.Level)),
		},
		Level: h.Level,
	}
	para.Add(&seg)
	var out spec.Addable = &para
	return &out
}

func ConvertParagraph(p *ast.Paragraph, centered bool) *spec.Addable {
	if string(p.Text(globals.File)) == "\\fill" {
		var a spec.Addable = &Filler{}
		return &a
	}
	para := Paragraph{
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
			seg := ConvertText(n.(*ast.Text))
			para.Add(&seg)
		case ast.KindCodeSpan:
			seg := ConvertCodeSpan(n.(*ast.CodeSpan))
			para.Add(&seg)
		case ast.KindEmphasis:
			seg := ConvertEmphasis(n.(*ast.Emphasis))
			para.Add(&seg)
		case goldmark_cite.Kind:
			seg := ConvertCiteProc(n.(*goldmark_cite.Node))
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

func ConvertText(text *ast.Text) spec.Segment {
	return spec.Segment{
		Content: string(text.Text(globals.File)),
		Font:    spec.SerifRegular,
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
		Font:    spec.Monospace,
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
		t.Font = spec.SerifItalic
	case 2:
		t.Font = spec.SerifBold
	}
	return t
}

func ConvertCiteProc(cite *goldmark_cite.Node) spec.Segment {
	return spec.Segment{
		Content: fmt.Sprintf("[%d]", globals.BibIndices[cite.Key]),
		Font:    spec.SerifRegular,
	}
}

func ConvertList(list *ast.List) *spec.Addable {
	para := List{
		Text: spec.Text{
			FontSize: globals.Cfg.Text.FontSize,
			//LineHeight: globals.Cfg.LineHeight * 1.4,
			LineHeight: globals.Cfg.Text.ListLineHeight,
			Offset:     float64(globals.Cfg.Text.FontSize),
		},
	}
	i := 1
	for n := list.FirstChild(); n != nil; n = n.NextSibling() {
		// do not support nested lists for now
		switch n.Kind() {
		case ast.KindListItem:
			prefix := ""
			if list.IsOrdered() {
				prefix = strconv.Itoa(i)
				i++
			}
			seg := ConvertListItem(n.(*ast.ListItem), prefix, string(list.Marker))
			para.Add(&seg)
		}
	}
	var a spec.Addable = &para
	return &a
}

func ConvertListItem(item *ast.ListItem, prefix string, marker string) spec.Segment {
	seg := spec.Segment{
		Content: prefix + marker + " " + string(item.Text(globals.File)),
		Font:    spec.SerifRegular,
	}
	return seg
}

func ConvertImage(image *ast.Image, node ast.Node) (retO *spec.XObject, retA *spec.Addable, retP *spec.Addable) {
	mul := .95
	if node.ChildCount() == 2 {
		mul, _ = strconv.ParseFloat(strings.TrimSpace(string(node.FirstChild().NextSibling().Text(globals.File))), 64)
	}
	io, ia := spec.NewImageObject(string(image.Destination), mul)
	retO = &io
	retA = &ia
	para := Paragraph{
		Text: spec.Text{
			FontSize:   globals.Cfg.Text.FontSize - 1,
			LineHeight: 1.0,
		},
		Centered: true,
	}
	para.Add(&spec.Segment{
		Content: string(image.Text(globals.File)),
		Font:    spec.SerifRegular,
	})
	var a spec.Addable = &para
	retP = &a
	return
}

func ConvertBlockquote(bq *ast.Blockquote) (ret []*spec.Addable) {
	for c := bq.FirstChild(); c != nil; c = c.NextSibling() {
		switch c.Kind() {
		case ast.KindParagraph:
			ret = append(ret, ConvertParagraph(c.(*ast.Paragraph), true))
		}
	}
	return
}
