package pdf

import (
	"crypto/md5"
	"fmt"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	goldmark_math "github.com/sett17/mdpaper/v2/goldmark-math"
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/conversions"
	"github.com/sett17/mdpaper/v2/pdf/conversions/options"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/register"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/sett17/mdpaper/v2/pdf/toc"
	"github.com/yuin/goldmark/ast"
	"path"
	"time"
)

func FromAst(md ast.Node) *spec.PDF {
	start := time.Now()
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

	//region EXTRACT INFO FOR REGISTERS
	for n := md.FirstChild(); n != nil; n = n.NextSibling() {
		switch n.Kind() {
		case ast.KindParagraph:
			if n.(*ast.Paragraph).ChildCount() <= 2 && n.FirstChild().Kind() == ast.KindImage {
				image := n.FirstChild().(*ast.Image)

				opt, err := options.Parse(string(image.Text(globals.File)))
				if err != nil {
					cli.Error(fmt.Errorf("error parsing image options: %w", err), false)
				}

				var id string
				_, id = path.Split(string(image.Destination))
				id = globals.NameEncode(id)
				optId, ok := opt.GetString("id")
				if ok {
					id = optId
				}
				figInfo := globals.NewFigureInformation(string(image.Title), id)
				globals.Figures[id] = figInfo
			}
		case ast.KindFencedCodeBlock:
			lang := string(n.(*ast.FencedCodeBlock).Language(globals.File))
			if globals.Cfg.Code.Dot && lang == "dot" {
				fcb := n.(*ast.FencedCodeBlock)
				optionString := ""
				if fcb.Info != nil {
					optionString = options.Extract(string(fcb.Info.Text(globals.File)))
				}
				opts, err := options.Parse(optionString)
				if err != nil {
					cli.Error(fmt.Errorf("error parsing code options: %w", err), false)
				}

				title := ""
				if t, ok := opts.GetString("title"); ok {
					title = t
				} else if t, ok := opts.GetString("caption"); ok {
					title = t
				} else if t, ok := opts.GetString("label"); ok {
					title = t
				}

				//couldn't think of another 'nice' way to get a unique id
				hashBuf := make([]byte, 0)
				for i := 0; i < fcb.Lines().Len(); i++ {
					at := fcb.Lines().At(i)
					hashBuf = append(hashBuf, at.Value(globals.File)...)
				}
				startByte := fcb.Lines().At(0).Start
				hashBuf = append(hashBuf, (byte)(startByte>>24), (byte)(startByte>>16), (byte)(startByte>>8), (byte)(startByte)) //include this to make it unique to this special block
				id := fmt.Sprintf("%x", md5.Sum(hashBuf))

				optId, ok := opts.GetString("id")
				if ok {
					id = optId
				}
				figInfo := globals.NewFigureInformation(title, id)
				globals.Figures[id] = figInfo
			}
		}
	}
	if len(globals.Figures) == 0 {
		globals.Cfg.Tof.Enabled = false
	}
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
			if globals.Cfg.Code.Dot && lang == "dot" {
				xo, i, p := conversions.Dot(n.(*ast.FencedCodeBlock))
				paper.Add(i)
				if p != nil {
					paper.Add(p)
				}
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
	parsed := time.Now()
	globals.ImageSync.Wait()
	cli.Other("Added %v elements in %v\n", i, parsed.Sub(start))

	headings := make([]*elements.Heading, 0)
	for _, e := range paper.Elements {
		if h, ok := (*e).(*elements.Heading); ok {
			headings = append(headings, h)
		}
	}
	if globals.Cfg.Citation.Enabled {
		headings = append(headings, register.Citation.Heading())
	}
	if globals.Cfg.Tof.Enabled {
		headings = append(headings, register.Figures.Heading())
	}
	//TODO add other registry headings
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
	pagesDict := spec.NewDictObject()
	pagesDict.Set("Type", "/Pages")
	pdf.AddObject(pagesDict.Pointer())
	catalog.Set("Pages", pagesDict.Reference())

	pages := make([]*abstracts.Page, 0)

	pagesArray := spec.NewArray()
	for !paper.Finished() {
		page := abstracts.NewPage(&paper, globals.Cfg.Page.Columns)
		pages = append(pages, page)
	}
	if globals.Cfg.Toc.Enabled {
		register.Content.Tree = toc.GenerateChapterTree(headings)
		register.Content.GenerateEntries()
		register.Content.GeneratePages()
		pages = append(register.Content.Pages, pages...) //prepend
	}
	if globals.Cfg.Cover.Enabled {
		cover := abstracts.GenerateCover()
		coverPage := abstracts.NewEmptyPage()
		coverPage.Columns = append(coverPage.Columns, cover)
		pages = append([]*abstracts.Page{coverPage}, pages...) //prepend
	}
	if globals.Cfg.Tof.Enabled {
		register.Figures.GenerateEntries()
		register.Figures.GeneratePages()
		pages = append(pages, register.Figures.Pages...)
	}
	if globals.Cfg.Citation.Enabled {
		register.Citation.GenerateEntries()
		register.Citation.GeneratePages()
		pages = append(pages, register.Citation.Pages...)
	}

	for i, page := range pages {
		page.Number = i + 1
		for _, col := range page.Columns {
			for _, e := range col.Content {
				if h, ok := (*e).(*elements.Heading); ok {
					h.Page = page.Number
				} else if i, ok := (*e).(*spec.ImageAddable); ok {
					if f, ok := globals.Figures[i.Id]; ok {
						f.Used = append(f.Used, page.Number)
					}
				}
			}
		}
	}

	//basically backtracking to insert page numbers where needed
	if globals.Cfg.Tof.Enabled {
		register.Figures.InsertPageNumbers()
	}
	if globals.Cfg.Toc.Enabled {
		register.Content.InsertPageNumbers()
		register.Content.InsertLinks()
	}

	for _, page := range pages {
		page.AddToPdf(&pdf, pageResources, pagesDict.Reference(), &pagesArray)
	}

	pagesDict.Set("Kids", pagesArray)
	pagesDict.Set("Count", len(pagesArray.Items))
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

	cli.Output("PDF generated in %v\n", time.Since(parsed))
	return &pdf
}
