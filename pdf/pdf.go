package pdf

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
	"mdpaper/pdf/spec"
)

func FromAst(md ast.Node) *spec.PDF {
	pdf := spec.PDF{}
	paper := Paper{}

	catalog := spec.NewDictObject()
	catalog.Set("Type", "/Catalog")
	catalog.Set("PageMode", "/UseOutlines")
	pdf.SetRoot(catalog.Pointer())

	//region add fonts to pdf
	timesRef, timesName := spec.TimesRegular.AddToPDF(&pdf)
	boldRef, boldName := spec.TimesBold.AddToPDF(&pdf)
	italicRef, italicName := spec.TimesItalic.AddToPDF(&pdf)
	boldItalicRef, boldItalicName := spec.TimesBoldItalic.AddToPDF(&pdf)
	courierRef, courierName := spec.CourierMono.AddToPDF(&pdf)
	helveticaRef, helveticaName := spec.HelveticaRegular.AddToPDF(&pdf)

	pageResources := spec.NewDict()
	fonts := spec.NewDict()
	fonts.Set(timesName, timesRef)
	fonts.Set(boldName, boldRef)
	fonts.Set(italicName, italicRef)
	fonts.Set(boldItalicName, boldItalicRef)
	fonts.Set(courierName, courierRef)
	fonts.Set(helveticaName, helveticaRef)
	pageResources.Set("Font", fonts)
	//endregion

	//region convert all nodes to objects and accumulate in paper
	for n := md.FirstChild(); n != nil; n = n.NextSibling() {
		switch n.Kind() {
		case ast.KindHeading:
			h := ConvertHeading(n.(*ast.Heading))
			paper.Add(h)
		case ast.KindParagraph:
			p := ConvertParagraph(n.(*ast.Paragraph))
			paper.Add(p)
		}
	}
	headings := make([]*Heading, 0)
	for _, e := range paper.Elements {
		if h, ok := (*e).(*Heading); ok {
			headings = append(headings, h)
		}
	}
	chapter := GenerateChapterTree(headings)
	for i, c := range chapter.RootNodes() {
		c.(ChapterNode).Heading.Prefix = [6]int{i + 1}
		chapter.GenerateNumbering(c.(ChapterNode))
	}
	//endregion

	//region generate pages and add to pdf
	pages := spec.NewDictObject()
	pages.Set("Type", "/Pages")
	pdf.AddObject(pages.Pointer())
	catalog.Set("Pages", pages.Reference())

	pageNumber := 0
	pagesArray := spec.NewArray()
	for !paper.Finished() {
		pageNumber++
		page := NewPage(&paper, pageNumber)
		page.AddToPdf(&pdf, pageResources, catalog.Reference(), &pagesArray)
	}

	pages.Set("Kids", pagesArray)
	pages.Set("Count", len(pagesArray.Items))
	//endregion

	fmt.Println(chapter)

	return &pdf
}
