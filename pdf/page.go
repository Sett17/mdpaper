package pdf

import (
	"fmt"
	"math"
	"mdpaper/globals"
	"mdpaper/pdf/spec"
)

type Page struct {
	spec.Dictionary
	Number  int
	Columns []*Column
}

func NewPage(paper *Paper, number int) *Page {
	p := &Page{Number: number}
	p.Set("Type", "Page")
	mediaBox := spec.NewArray()
	mediaBox.Add(0, 0, globals.A4Width, globals.A4Height)
	p.Set("MediaBox", string(mediaBox.Bytes()))
	if globals.Cfg.Columns == 1 {
		p.Columns = append(p.Columns, paper.SingleColumn())
	} else if globals.Cfg.Columns == 2 {
		c1, c2 := paper.DoubleColumn()
		p.Columns = append(p.Columns, c1, c2)
	}
	return p
}

func (p *Page) AddToPdf(pdf *spec.PDF, res spec.Dictionary, catalog string, pages *spec.Array) {
	page := spec.NewDictObject()
	page.M = p.M
	p.Set("Parent", catalog)
	p.Set("Resources", string(res.Bytes()))
	c := spec.Array{}
	for _, col := range p.Columns {
		pdf.AddObject(col.Pointer())
		c.Add(col.Reference())
	}
	if p.Number > 0 {
		pN := spec.NewStreamObject()
		seg := spec.Segment{
			Content: fmt.Sprintf("%d", p.Number),
			Font:    &spec.HelveticaRegular,
		}
		para := spec.Text{
			FontSize:   10,
			LineHeight: 1,
		}
		para.Add(&seg)
		para.SetPos(globals.A4Width/2.0-seg.Font.WordWidth(seg.Content, para.FontSize)/2.0, globals.Cfg.Margin/2)
		para.Process(math.MaxFloat64)
		var a spec.Addable = &para
		pN.Add(&a)
		pdf.AddObject(pN.Pointer())
		c.Add(pN.Reference())
	}
	p.Set("Contents", string(c.Bytes()))
	pdf.AddObject(page.Pointer())
	pages.Add(page.Reference())
}
