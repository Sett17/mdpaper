package abstracts

import (
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/conversions"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
)

func GenerateCover() *Column {
	col := NewColumn(
		globals.A4Width-(globals.MmToPt(40)*2),
		globals.A4Height-(globals.MmToPt(30*2)),
		globals.MmToPt(40),
		globals.A4Height-globals.MmToPt(30),
	)

	//spacer1 := spacing.NewSpacer(globals.MmToPt(10))

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

	abstract := (*conversions.String(globals.Cfg.Cover.Abstract)).(*elements.Paragraph)
	abstract.FontSize = 11
	abstract.LineHeight = 1.1

	var abs spec.Addable = abstract

	spacer3 := spacing.NewSpacer(globals.MmToPt(10))

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

	contentHeight := head.Height() + sub.Height() + abstract.Height() + author.Height()
	spacerHeight := (col.MaxHeight - contentHeight) / 4
	//spacer1.H = spacerHeight
	spacer2.H = spacerHeight
	spacer3.H = spacerHeight

	//var s1 spec.Addable = spacer1
	//col.Add(&s1)
	col.Add(&h)
	col.Add(&s)
	var s2 spec.Addable = spacer2
	col.Add(&s2)
	col.Add(&abs)
	var s3 spec.Addable = spacer3
	col.Add(&s3)
	col.Add(&a)

	return col
}
