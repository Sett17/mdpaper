package conversions

import (
	"bytes"
	"github.com/sett17/citeproc-js-go/csljson"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/goldmark-cite"
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"strings"
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
		default:
			continue
		}
	}
	var a spec.Addable = &txt
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

func Citation(cite *goldmark_cite.Citation) spec.Segment {
	if !globals.Cfg.Citation.Enabled {
		return spec.Segment{
			Content: "[" + strings.Join(cite.Keys, ", ") + "]",
			Font:    spec.SerifRegular,
		}
	}
	items := make([]csljson.Item, len(cite.Keys))

	for i, key := range cite.Keys {
		item, ok := globals.Citations[key]
		if !ok {
			//cli.Warning("Citation %s not found\n", key)
		} else {
			items[i] = item
		}
	}

	citationInsert, err := globals.Citeproc.ProcessCitationCluster(items...)
	if err != nil {
		cli.Warning("Citation cluster %+v could not be processed (%s)", cite.Keys, err)
		return spec.Segment{}
	}

	citationInsert = strings.TrimSpace(citationInsert) + " "

	//this is temporary, until I can really properly figure out all the font stuff + this bothers me for ieee
	citationInsert = strings.ReplaceAll(citationInsert, "–", "-")

	return spec.Segment{
		Content: citationInsert,
		Font:    spec.SerifRegular,
	}
}
