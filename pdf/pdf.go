package pdf

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
	"mdpaper/globals"
	"mdpaper/pdf/spec"
	"time"
)

func FromAst(md ast.Node) *spec.PDF {
	pdf := spec.PDF{}
	paper := Paper{}

	catalog := spec.NewDictObject()
	catalog.Set("Type", "/Catalog")
	catalog.Set("PageMode", "/UseOutlines")
	pdf.AddObject(catalog.Pointer())
	pdf.Root = catalog.Reference()
	info := spec.NewDictObject()
	info.Set("Producer", "mdpaper")
	info.Set("CreationDate", time.Now().Format("20060102150405-07"))
	info.Set("Title", globals.Cfg.Title)
	info.Set("Author", globals.Cfg.Authors[0])
	pdf.AddObject(info.Pointer())
	pdf.Info = info.Reference()

	//region add fonts to pdf
	tinoRegRef, tinoRegName := spec.SerifRegular.AddToPDF(&pdf)
	tinoBoldRef, tinoBoldName := spec.SerifBold.AddToPDF(&pdf)
	tinoItalicRef, tinoItalicName := spec.SerifItalic.AddToPDF(&pdf)
	latoRegRef, latoRegName := spec.SansRegular.AddToPDF(&pdf)
	latoBoldRef, latoBoldName := spec.SansBold.AddToPDF(&pdf)
	scpRef, scpName := spec.Monospace.AddToPDF(&pdf)

	pageResources := spec.NewDict()
	fonts := spec.NewDict()
	fonts.Set(tinoRegName, tinoRegRef)
	fonts.Set(tinoBoldName, tinoBoldRef)
	fonts.Set(tinoItalicName, tinoItalicRef)
	fonts.Set(latoRegName, latoRegRef)
	fonts.Set(latoBoldName, latoBoldRef)
	fonts.Set(scpName, scpRef)
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
		case ast.KindList:
			l := ConvertList(n.(*ast.List))
			paper.Add(l)
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
		page.AddToPdf(&pdf, pageResources, pages.Reference(), &pagesArray)
	}
	if globals.Cfg.ToC {
		toc := GenerateTOC(&chapters)
		tocPage.Columns = append(tocPage.Columns, toc.GenerateColumn())
		links := toc.GenerateLinks()
		tocPage.Annots = append(tocPage.Annots, links...)
		tocPage.AddToPdf(&pdf, pageResources, pages.Reference(), &pagesArray)
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
