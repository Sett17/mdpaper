package conversions

import (
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"strconv"
)

func List(list *ast.List) (adds []*spec.Addable) {
	i := 1
	for n := list.FirstChild(); n != nil; n = n.NextSibling() {
		// do not support nested lists for now
		switch n.Kind() {
		case ast.KindListItem:
			number := -1
			if list.IsOrdered() {
				number = i
			}
			adds = append(adds, ListItem(n.(*ast.ListItem), number, string(list.Marker)))
		}
	}
	return
}

func ListItem(item *ast.ListItem, number int, marker string) *spec.Addable {
	prefix := marker + " "
	if number >= 0 {
		prefix = strconv.Itoa(number) + prefix
	}
	ret := elements.ListItem{
		Text: spec.Text{
			FontSize:   globals.Cfg.Text.FontSize,
			LineHeight: globals.Cfg.Text.ListLineHeight,
			Offset:     spec.SerifRegular.WordWidth(prefix, globals.Cfg.Text.FontSize) * 1.2,
			Margin:     (globals.Cfg.Text.ListLineHeight - 1) * float64(globals.Cfg.Text.FontSize) * .2,
		},
		Prefix: prefix,
	}
	child := item.FirstChild()
	var txt *spec.Addable
	if block, ok := child.(*ast.TextBlock); ok {
		txt = TextBlock(block)
	}
	if txt != nil {
		ret.Add((*txt).(*spec.Text).Segments...)
	}
	var a spec.Addable = &ret
	return &a
}
