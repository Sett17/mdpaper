package pdf

import (
	"fmt"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/spec"
)

type Page struct {
	spec.DictionaryObject
	DisplayNumber int
	RealNumber    int
	Columns       []*Column
	Annots        []*spec.DictionaryObject
}

func NewPage(paper *Paper, displayNumber int, realNumber int) *Page {
	p := &Page{DisplayNumber: displayNumber, RealNumber: realNumber, DictionaryObject: spec.NewDictObject()}
	p.Set("Type", "/Page")
	mediaBox := spec.NewArray()
	mediaBox.Add(0, 0, globals.A4Width, globals.A4Height)
	p.Set("MediaBox", mediaBox)
	if globals.Cfg.Page.Columns == 1 {
		p.Columns = append(p.Columns, paper.SingleColumn())
	} else if globals.Cfg.Page.Columns >= 2 {
		c1, c2 := paper.DoubleColumn()
		p.Columns = append(p.Columns, c1, c2)
	}
	for _, col := range p.Columns {
		for _, a := range col.Content {
			if h, ok := (*a).(*Heading); ok {
				h.Page = realNumber
				h.DisplayPage = displayNumber
			}
		}
	}
	return p
}

func NewEmptyPage(displayNumber int, realnumber int) *Page {
	p := &Page{DisplayNumber: displayNumber, RealNumber: realnumber, DictionaryObject: spec.NewDictObject()}
	p.Set("Type", "/Page")
	mediaBox := spec.NewArray()
	mediaBox.Add(0, 0, globals.A4Width, globals.A4Height)
	p.Set("MediaBox", string(mediaBox.Bytes()))
	return p
}

func (p *Page) AddToPdf(pdf *spec.PDF, res spec.Dictionary, pagesRef string, pages *spec.Array) {
	p.Set("Parent", pagesRef)
	p.Set("Resources", res)
	annotsArray := spec.NewArray()
	for _, a := range p.Annots {
		pdf.AddObject(a.Pointer())
		annotsArray.Add(a.Reference())
	}
	p.Set("Annots", annotsArray)
	c := spec.Array{}
	for _, col := range p.Columns {
		pdf.AddObject(col.Pointer())
		c.Add(col.Reference())
	}
	if p.DisplayNumber > 0 && globals.Cfg.Page.PageNumbers {
		pN := spec.NewStreamObject()
		//pN.Deflate = true
		seg := spec.Segment{
			Content: fmt.Sprintf("%d", p.DisplayNumber),
			Font:    spec.SansRegular,
		}
		para := spec.Text{
			FontSize:   10,
			LineHeight: 1,
		}
		para.Add(&seg)
		para.SetPos(globals.A4Width/2.0-seg.Font.WordWidth(seg.Content, para.FontSize)/2.0, globals.Cfg.Page.MarginBottom*1.5)
		para.Process(seg.Font.WordWidth(seg.Content, para.FontSize))
		var a spec.Addable = &para
		pN.Add(&a)
		pdf.AddObject(pN.Pointer())
		c.Add(pN.Reference())
	}
	p.Set("Contents", c)
	pdf.AddObject(p.DictionaryObject.Pointer())
	if p.RealNumber > len(pages.Items) {
		pages.Add(p.DictionaryObject.Reference())
	} else {
		pagesOld := pages.Items
		pages.Items = make([]interface{}, 0)
		pages.Items = append(pages.Items, pagesOld[:p.RealNumber-1]...)
		pages.Items = append(pages.Items, p.DictionaryObject.Reference())
		pages.Items = append(pages.Items, pagesOld[p.RealNumber-1:]...)
	}
}
