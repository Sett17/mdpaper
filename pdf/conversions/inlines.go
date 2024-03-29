package conversions

import (
	"bytes"
	"github.com/sett17/citeproc-js-go/csljson"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	goldmark_cite "github.com/sett17/mdpaper/v2/goldmark-cite"
	goldmark_figref "github.com/sett17/mdpaper/v2/goldmark-figref"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"strconv"
	"strings"
)

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

func FigRef(ref *goldmark_figref.FigRef) spec.Segment {
	fig := globals.Figures[ref.FigureKey]
	if fig == nil {
		//return spec.Segment{}
		return spec.Segment{
			Content: globals.Cfg.Text.FigureText + " " + ref.FigureKey + " not found",
			Font:    spec.SerifItalic,
		}
	}
	return spec.Segment{
		Content: globals.Cfg.Text.FigureText + " " + strconv.Itoa(fig.Number),
		Font:    spec.SerifItalic,
	}
}
