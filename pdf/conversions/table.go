package conversions

import (
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
	east "github.com/yuin/goldmark/extension/ast"
)

func Table(t *east.Table) *spec.Addable {
	table := &elements.Table{}
	for c := t.FirstChild(); c != nil; c = c.NextSibling() {
		switch r := c.(type) {
		case *east.TableHeader:
			converted := east.TableRow{
				BaseBlock:  r.BaseBlock,
				Alignments: r.Alignments,
			}
			table.Header = convertRow(&converted)
		case *east.TableRow:
			table.Rows = append(table.Rows, convertRow(r))
		}
	}

	table.Padding = globals.MmToPt(globals.Cfg.Table.Padding)

	var addable spec.Addable = table
	return &addable
}

func convertRow(r *east.TableRow) []*elements.TableCell {
	var cells []*elements.TableCell
	for c := r.FirstChild(); c != nil; c = c.NextSibling() {
		switch c := c.(type) {
		case *east.TableCell:
			tb := ast.TextBlock{BaseBlock: c.BaseBlock}
			text := (*TextBlock(&tb)).(*spec.Text)
			cells = append(cells,
				&elements.TableCell{
					Text:     *text,
					Centered: c.Alignment == east.AlignCenter,
				},
			)
		}
	}
	return cells
}
