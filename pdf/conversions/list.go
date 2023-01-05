package conversions

import (
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"strconv"
)

func List(list *ast.List) *spec.Addable {
	para := elements.List{
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
			seg := ListItem(n.(*ast.ListItem), prefix, string(list.Marker))
			para.Add(&seg)
		}
	}
	var a spec.Addable = &para
	return &a
}

func ListItem(item *ast.ListItem, prefix string, marker string) spec.Segment {
	seg := spec.Segment{
		Content: prefix + marker + " " + string(item.Text(globals.File)),
		Font:    spec.SerifRegular,
	}
	return seg
}
