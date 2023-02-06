package pdf

import (
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	goldmark_math "github.com/sett17/mdpaper/v2/goldmark-math"
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/conversions"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/references"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/sett17/mdpaper/v2/pdf/toc"
	"github.com/yuin/goldmark/ast"
	"time"
)

func FromAst(md ast.Node) *spec.PDF {
	pdf := spec.PDF{}
	paper := abstracts.Paper{}

	catalog := spec.NewDictObject()
	catalog.Set("Type", "/Catalog")
	catalog.Set("PageMode", "/UseOutlines")
	pdf.AddObject(catalog.Pointer())
	pdf.Root = catalog.Reference()
	info := spec.NewDictObject()
	info.Set("Producer", "mdpaper")
	info.Set("CreationDate", time.Now().Format("20060102150405-07"))
	info.Set("Title", globals.Cfg.Paper.Title)
	info.Set("Author", globals.Cfg.Paper.Author)
	pdf.AddObject(info.Pointer())
	pdf.Info = info.Reference()

	//region FONTS
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

	//region CONVERSIONS
	cli.Output("Addding Elements\n")
	i := 0
	for n := md.FirstChild(); n != nil; n = n.NextSibling() {
		switch n.Kind() {
		case ast.KindHeading:
			h := conversions.Heading(n.(*ast.Heading))
			paper.Add(h)
		case ast.KindParagraph:
			if n.(*ast.Paragraph).ChildCount() <= 2 && n.FirstChild().Kind() == ast.KindImage {
				xo, i, p := conversions.Image(n.FirstChild().(*ast.Image), n)
				paper.Add(i)
				if xo != nil {
					paper.AddXObject(xo)
				}
				paper.Add(p)
			} else {
				p := conversions.Paragraph(n.(*ast.Paragraph), false)
				paper.Add(p)
			}
		case ast.KindList:
			l := conversions.List(n.(*ast.List))
			paper.Add(l...)
		case ast.KindBlockquote:
			b := conversions.Blockquote(n.(*ast.Blockquote))
			paper.Add(b...)
		case ast.KindFencedCodeBlock:
			lang := string(n.(*ast.FencedCodeBlock).Language(globals.File))
			if globals.Cfg.Code.Mermaid && lang == "mermaid" {
				xo, i := conversions.Mermaid(n.(*ast.FencedCodeBlock))
				paper.Add(i)
				if xo != nil {
					paper.AddXObject(xo)
				}
				continue
			}
			c := conversions.Code(n.(*ast.FencedCodeBlock))
			paper.Add(c)
		case goldmark_math.KindMathBlock:
			xo, i := conversions.Math(n.(*goldmark_math.MathBlock))
			paper.Add(i)
			if xo != nil {
				paper.AddXObject(xo)
			}
		}
		cli.Other(".")
		i++
		if i%80 == 0 {
			cli.Other("\n")
		}
	}
	cli.Other("\n")
	cli.Other("Added %v elements\n", len(paper.Elements))

	headings := make([]*elements.Heading, 0)
	for _, e := range paper.Elements {
		if h, ok := (*e).(*elements.Heading); ok {
			headings = append(headings, h)
		}
	}
	if globals.Cfg.Citation.Enabled {
		references.GenerateCitationHeading()
		headings = append(headings, references.CitationHeading)
	}
	toc.GenerateChapterTree(headings)
	chapters := toc.GenerateChapterTree(headings)
	for i, c := range chapters.Roots() {
		c.Heading.SetPrefix([6]int{i + 1})
		chapters.GenerateNumbering(c)
	}
	//endregion

	//region XOBJECTS
	xobjs := spec.NewDict()
	for _, xo := range paper.XObjects {
		xobjs.Set(xo.Name, xo.Reference())
		pdf.AddObject(xo.Pointer())
	}
	pageResources.Set("XObject", xobjs)
	//endregion

	//region PAGES
	pages := spec.NewDictObject()
	pages.Set("Type", "/Pages")
	pdf.AddObject(pages.Pointer())
	catalog.Set("Pages", pages.Reference())

	displayPageNumber := globals.Cfg.Page.StartPageNumber - 1
	realPageNumber := 0
	pagesArray := spec.NewArray()
	var coverPage *abstracts.Page
	if globals.Cfg.Cover.Enabled {
		realPageNumber++
		coverPage = abstracts.NewEmptyPage(-1, realPageNumber)
	}
	var tocPage *abstracts.Page
	if globals.Cfg.Toc.Enabled {
		realPageNumber++
		tocPage = abstracts.NewEmptyPage(0, realPageNumber)
	}
	for !paper.Finished() {
		displayPageNumber++
		realPageNumber++
		page := abstracts.NewPage(&paper, displayPageNumber, realPageNumber, globals.Cfg.Page.Columns)
		page.AddToPdf(&pdf, pageResources, pages.Reference(), &pagesArray)
	}
	if globals.Cfg.Citation.Enabled {
		cits := abstracts.Paper{}
		var a spec.Addable = references.CitationHeading
		cits.Add(&a)
		cits.Add(references.Citations()...)
		for !cits.Finished() {
			displayPageNumber++
			realPageNumber++
			page := abstracts.NewPage(&cits, displayPageNumber, realPageNumber, 1)
			page.AddToPdf(&pdf, pageResources, pages.Reference(), &pagesArray)
		}
	}
	if globals.Cfg.Cover.Enabled {
		cover := abstracts.GenerateCover()
		coverPage.Columns = append(coverPage.Columns, cover)
		coverPage.AddToPdf(&pdf, pageResources, pages.Reference(), &pagesArray)
	}
	if globals.Cfg.Toc.Enabled {
		tableOfContents := toc.GenerateTOC(&chapters)
		tocPage.Columns = append(tocPage.Columns, tableOfContents.GenerateColumn())
		links := tableOfContents.GenerateLinks()
		tocPage.Annots = append(tocPage.Annots, links...)
		tocPage.AddToPdf(&pdf, pageResources, pages.Reference(), &pagesArray)
	}

	pages.Set("Kids", pagesArray)
	pages.Set("Count", len(pagesArray.Items))
	//endregion

	//region OUTLINE
	outlines := spec.NewDictObject()
	outlines.Set("Type", "/Outlines")
	outlines.Set("Count", len(chapters))
	pdf.AddObject(outlines.Pointer())
	catalog.Set("Outlines", outlines.Reference())
	outlineItems := chapters.GenerateOutline(&outlines)
	for _, item := range outlineItems {
		pdf.AddObject(item.Pointer())
	}
	//endregion

	//fmt.Println(chapters.String())

	return &pdf
}
