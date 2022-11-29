package pdf

import (
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
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
	info.Set("Title", globals.Cfg.Paper.Title)
	info.Set("Author", globals.Cfg.Paper.Author)
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
			if n.(*ast.Paragraph).ChildCount() <= 2 && n.FirstChild().Kind() == ast.KindImage {
				xo, i, p := ConvertImage(n.FirstChild().(*ast.Image), n)
				paper.Add(i)
				if xo != nil {
					paper.AddXObject(xo)
				}
				paper.Add(p)
			} else {
				p := ConvertParagraph(n.(*ast.Paragraph), false)
				paper.Add(p)
			}
		case ast.KindList:
			l := ConvertList(n.(*ast.List))
			paper.Add(l)
		case ast.KindBlockquote:
			b := ConvertBlockquote(n.(*ast.Blockquote))
			paper.Add(b...)
		}
	}

	// add elements for citations stuff
	//if globals.Cfg.Citation.Enabled {
	//	seg := spec.Segment{
	//		Content: "Citations",
	//		Font:    spec.SansBold,
	//	}
	//	head := Heading{
	//		Text: spec.Text{
	//			FontSize: int(float64(globals.Cfg.Text.FontSize) * 1.2),
	//			//LineHeight: globals.Cfg.LineHeight * 1.5,
	//			LineHeight: 1.0,
	//			Offset:     0.0,
	//		},
	//		Level: 0,
	//	}
	//	head.Add(&seg)
	//	var out spec.Addable = &head
	//	paper.Add(&out)
	//	para := List{
	//		Text: spec.Text{
	//			FontSize: globals.Cfg.Text.FontSize,
	//			//LineHeight: globals.Cfg.LineHeight * 1.4,
	//			LineHeight: globals.Cfg.Text.ListLineHeight,
	//			Offset:     float64(globals.Cfg.Text.FontSize),
	//		},
	//	}
	//	for key, idx := range globals.BibIndices {
	//		seg := spec.Segment{
	//			Content: fmt.Sprintf("[%d] %s", idx, globals.IEEE(globals.Bibs[key])),
	//			Font:    spec.SerifRegular,
	//		}
	//		para.Add(&seg)
	//	}
	//	var a spec.Addable = &para
	//	paper.Add(&a)
	//}

	headings := make([]*Heading, 0)
	for _, e := range paper.Elements {
		if h, ok := (*e).(*Heading); ok {
			headings = append(headings, h)
		}
	}
	if globals.Cfg.Citation.Enabled {
		headings = append(headings, CitationHeading)
	}
	GenerateChapterTree(headings)
	chapters := GenerateChapterTree(headings)
	for i, c := range chapters.Roots() {
		c.Heading.Prefix = [6]int{i + 1}
		chapters.GenerateNumbering(c)
	}
	//endregion

	//region xobjects
	xobjs := spec.NewDict()
	for _, xo := range paper.XObjects {
		xobjs.Set(xo.Name, xo.Reference())
		pdf.AddObject(xo.Pointer())
	}
	pageResources.Set("XObject", xobjs)
	//endregion

	//region generate pages and add to pdf
	pages := spec.NewDictObject()
	pages.Set("Type", "/Pages")
	pdf.AddObject(pages.Pointer())
	catalog.Set("Pages", pages.Reference())

	displayPageNumber := globals.Cfg.Page.StartPageNumber - 1
	realPageNumber := 0
	pagesArray := spec.NewArray()
	var tocPage *Page
	if globals.Cfg.Toc.Enabled {
		realPageNumber++
		tocPage = NewEmptyPage(0, realPageNumber)
	}
	for !paper.Finished() {
		displayPageNumber++
		realPageNumber++
		page := NewPage(&paper, displayPageNumber, realPageNumber, globals.Cfg.Page.Columns)
		page.AddToPdf(&pdf, pageResources, pages.Reference(), &pagesArray)
	}
	if globals.Cfg.Citation.Enabled {
		cits := Paper{}
		var a spec.Addable = CitationHeading
		cits.Add(&a)
		cits.Add(citationList())
		for !cits.Finished() {
			displayPageNumber++
			realPageNumber++
			page := NewPage(&cits, displayPageNumber, realPageNumber, 1)
			page.AddToPdf(&pdf, pageResources, pages.Reference(), &pagesArray)
		}
	}
	if globals.Cfg.Toc.Enabled {
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

	//fmt.Println(chapters.String())

	return &pdf
}
