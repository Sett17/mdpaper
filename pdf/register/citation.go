package register

import (
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
)

type Citation struct {
	common
}

func (c *Citation) Heading() *elements.Heading {
	if c.heading == nil {
		h := elements.Heading{
			Text: spec.Text{
				FontSize:   21.0,
				LineHeight: 1.0,
			},
			Level: 1}
		seg := spec.Segment{
			Content: globals.Cfg.Citation.Heading,
			Font:    spec.SansBold,
		}
		h.Add(&seg)
		c.heading = &h
	}
	return c.heading
}

func (c *Citation) GenerateEntries() {
	//c.common.GenerateEntries()
	c.paper = &abstracts.Paper{}

	var h spec.Addable = c.Heading()
	c.paper.Add(&h)

	spacer := spacing.NewSpacer(c.Heading().Height() + globals.MmToPt(globals.Cfg.Margins.HeadingTop+globals.Cfg.Margins.HeadingBottom))
	var s spec.Addable = spacer
	c.paper.Add(&s)

	bib, err := globals.Citeproc.MakeBibliography()
	if err != nil {
		cli.Error(err, false)
		return
	}
	for _, bibEntry := range bib {
		entry := Entry{
			Left:       bibEntry,
			Right:      "",
			LineHeight: globals.Cfg.Citation.BibLineHeight,
			FontSize:   globals.Cfg.Citation.BibFontSize,
			Line:       false,
			Font:       spec.SerifRegular,
			LeftAlign:  true,
		}
		var a spec.Addable = &entry
		c.paper.Add(&a)

		spacer := spacing.NewSpacer(globals.MmToPt(2))
		var s spec.Addable = spacer
		c.paper.Add(&s)
	}
}

func NewCitation() *Citation {
	return &Citation{}
}
