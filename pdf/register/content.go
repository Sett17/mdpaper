package register

import (
	"fmt"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/sett17/mdpaper/v2/pdf/toc"
)

type content struct {
	common
	Tree    toc.ChapterTree
	entries map[*Entry]*elements.Heading
}

func (c *content) Heading() *elements.Heading {
	if c.heading == nil {
		head := elements.Heading{
			Text: spec.Text{
				FontSize:   21,
				LineHeight: 1,
			},
			Level: 0,
		}
		headSeg := spec.Segment{
			Content: globals.Cfg.Toc.Heading,
			Font:    spec.SansBold,
		}
		head.Add(&headSeg)
		c.heading = &head
	}
	return c.heading
}

func (c *content) GenerateEntries() {
	c.paper = &abstracts.Paper{}

	var h spec.Addable = c.Heading()
	c.paper.Add(&h)

	spacer := spacing.NewSpacer(c.Heading().Height() + globals.MmToPt(globals.Cfg.Margins.HeadingTop+globals.Cfg.Margins.HeadingBottom))
	var s spec.Addable = spacer
	c.paper.Add(&s)

	for _, chapter := range c.Tree {
		entry := Entry{
			Left:       chapter.Heading.String(),
			Right:      "X",
			FontSize:   globals.Cfg.Toc.FontSize,
			LineHeight: globals.Cfg.Toc.LineHeight,
			Line:       true,
			Font:       spec.SansRegular,
			Offset:     globals.MmToPt(float64((chapter.Heading.Level - 1) * 10)),
			LeftAlign:  true,
		}
		c.entries[&entry] = chapter.Heading
		var a spec.Addable = &entry
		c.paper.Add(&a)
	}
}

func (c *content) InsertPageNumbers() {
	for _, page := range c.Pages {
		for _, col := range page.Columns {
			for _, el := range col.Content {
				if entry, ok := (*el).(*Entry); ok {
					if h, ok := c.entries[entry]; ok {
						entry.Right = fmt.Sprintf("%d", h.Page)
						entry.Process(entry.width)
					}
				}
			}
		}
	}
}

func (c *content) InsertLinks() {
	for _, page := range c.Pages {
		for _, col := range page.Columns {
			for _, el := range col.Content {
				if entry, ok := (*el).(*Entry); ok {
					if h, ok := c.entries[entry]; ok {
						annot := spec.NewDictObject()
						annot.Set("Type", "/Annot")
						annot.Set("Subtype", "/Link")
						annot.Set("Rect", fmt.Sprintf("[%f %f %f %f]", entry.Pos[0]+entry.Offset, entry.Pos[1]-entry.HeightOffset(), entry.rightOffset+entry.Pos[0], entry.Pos[1]+entry.Height()-entry.HeightOffset()))
						annot.Set("Border", "[0 0 0]")
						annot.Set("Dest", h.Destination())
						page.Annots = append(page.Annots, &annot)
					}
				}
			}
		}
	}
}

var Content = &content{entries: make(map[*Entry]*elements.Heading)}
