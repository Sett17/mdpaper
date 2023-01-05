package abstracts

import (
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spacing"
	"github.com/sett17/mdpaper/pdf/spec"
)

func GenerateCover() *Column {
	col := NewColumn(
		globals.A4Width-(globals.MmToPt(globals.Cfg.Page.MarginHori)*2),
		globals.A4Height-(globals.MmToPt(globals.Cfg.Page.MarginTop)+globals.MmToPt(globals.Cfg.Page.MarginBottom)),
		globals.MmToPt(globals.Cfg.Page.MarginHori),
		globals.A4Height-globals.MmToPt(globals.Cfg.Page.MarginTop),
	)

	spacer1 := spacing.NewSpacer(globals.MmToPt(10))

	headSeg := spec.Segment{
		Content: globals.Cfg.Paper.Title,
		Font:    spec.SansBold,
	}
	head := elements.Paragraph{
		Text: spec.Text{
			FontSize:   24,
			LineHeight: 1.4,
		},
		Centered: true,
	}
	head.Add(&headSeg)
	var h spec.Addable = &head

	subSeg := spec.Segment{
		Content: globals.Cfg.Cover.Subtitle,
		Font:    spec.SansRegular,
	}
	sub := elements.Paragraph{
		Text: spec.Text{
			FontSize:   16,
			LineHeight: 1.3,
		},
		Centered: true,
	}
	sub.Add(&subSeg)
	var s spec.Addable = &sub

	spacer2 := spacing.NewSpacer(globals.MmToPt(10))

	authorSeg := spec.Segment{
		Content: globals.Cfg.Paper.Author,
		Font:    spec.SansRegular,
	}
	author := elements.Paragraph{
		Text: spec.Text{
			FontSize:   16,
			LineHeight: 1.3,
		},
		Centered: true,
	}
	author.Add(&authorSeg)
	var a spec.Addable = &author

	contentHeight := head.Height() + sub.Height() + author.Height()
	spacerHeight := (col.MaxHeight - contentHeight) / 3.5
	spacer1.H = spacerHeight
	spacer2.H = spacerHeight

	var s1 spec.Addable = spacer1
	col.Add(&s1)
	col.Add(&h)
	col.Add(&s)
	var s2 spec.Addable = spacer2
	col.Add(&s2)
	col.Add(&a)

	return col
}
