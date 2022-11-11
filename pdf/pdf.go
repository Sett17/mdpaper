package pdf

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
	"mdpaper/globals"
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
	helveticaBoldRef, helveticaBoldName := spec.HelveticaBold.AddToPDF(&pdf)
	timesMonoRef, timesMonoName := spec.TimesMono.AddToPDF(&pdf)

	pageResources := spec.NewDict()
	fonts := spec.NewDict()
	fonts.Set(timesName, timesRef)
	fonts.Set(boldName, boldRef)
	fonts.Set(italicName, italicRef)
	fonts.Set(boldItalicName, boldItalicRef)
	fonts.Set(courierName, courierRef)
	fonts.Set(helveticaName, helveticaRef)
	fonts.Set(helveticaBoldName, helveticaBoldRef)
	fonts.Set(timesMonoName, timesMonoRef)
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
	//GenerateChapterTree(headings)
	chapters := GenerateChapterTree(headings)
	for i, c := range chapters.Roots() {
		c.Heading.Prefix = [6]int{i + 1}
		chapters.GenerateNumbering(c)
	}
	//endregion

	//region generate pages and add to pdf
	pages := spec.NewDictObject()
	pages.Set("Type", "/Pages")
	pdf.AddObject(pages.Pointer())
	catalog.Set("Pages", pages.Reference())

	displayPageNumber := 0
	realPageNumber := 0
	pagesArray := spec.NewArray()
	var tocPage *Page
	if globals.Cfg.ToC {
		realPageNumber++
		tocPage = NewEmptyPage(0, realPageNumber)
	}
	for !paper.Finished() {
		displayPageNumber++
		realPageNumber++
		page := NewPage(&paper, displayPageNumber, realPageNumber)
		page.AddToPdf(&pdf, pageResources, catalog.Reference(), &pagesArray)
	}
	if globals.Cfg.ToC {
		toc := GenerateTOC(&chapters)
		tocPage.Columns = append(tocPage.Columns, toc.GenerateColumn())
		links := toc.GenerateLinks()
		tocPage.Annots = append(tocPage.Annots, links...)
		tocPage.AddToPdf(&pdf, pageResources, catalog.Reference(), &pagesArray)
	}

	pages.Set("Kids", pagesArray)
	pages.Set("Count", len(pagesArray.Items))
	//endregion

	//region generate outline
	outlines := spec.NewDictObject()
	outlines.Set("Type", "/Outlines")
	outlines.Set("Count", len(chapters))
	pdf.AddObject(outlines.Pointer())
	catalog.Set("Outlines", outlines.Reference())
	outlineItems := chapters.GenerateOutline(&outlines, &pdf)
	for _, item := range outlineItems {
		pdf.AddObject(item.Pointer())
	}
	//endregion

	fmt.Println(chapters.String())

	return &pdf
}
