package register

import (
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/conversions"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"strings"
)

type citation struct {
	common
}

func (c *citation) Heading() *elements.Heading {
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

func (c *citation) GenerateEntries() {
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
		leftSegs := make([]*spec.Segment, 0)
		//bib text can hold inline html styling (i've only seen <i>)
		//for now just handle italic and bold
		bibEntry = strings.ReplaceAll(bibEntry, "<i>", "_")
		bibEntry = strings.ReplaceAll(bibEntry, "</i>", "_")
		bibEntry = strings.ReplaceAll(bibEntry, "<em>", "_")
		bibEntry = strings.ReplaceAll(bibEntry, "</em>", "_")
		bibEntry = strings.ReplaceAll(bibEntry, "<b>", "**")
		bibEntry = strings.ReplaceAll(bibEntry, "</b>", "**")
		bibEntry = strings.ReplaceAll(bibEntry, "<strong>", "**")
		bibEntry = strings.ReplaceAll(bibEntry, "</strong>", "**")
		leftSegs = (*conversions.String(bibEntry)).(*elements.Paragraph).Segments

		entry := Entry{
			Left:       leftSegs,
			Right:      make([]*spec.Segment, 0),
			LineHeight: globals.Cfg.Citation.BibLineHeight,
			FontSize:   globals.Cfg.Citation.BibFontSize,
			Line:       false,
			LeftAlign:  true,
		}
		var a spec.Addable = &entry
		c.paper.Add(&a)

		spacer := spacing.NewSpacer(globals.MmToPt(2))
		var s spec.Addable = spacer
		c.paper.Add(&s)
	}
}

var Citation = &citation{}
