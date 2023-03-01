package register

import (
	"fmt"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"sort"
)

type figures struct {
	common
}

func (f *figures) Heading() *elements.Heading {
	if f.heading == nil {
		h := elements.Heading{
			Text: spec.Text{
				FontSize:   21.0,
				LineHeight: 1.0,
			},
			Level: 1}
		seg := spec.Segment{
			Content: globals.Cfg.Tof.Heading,
			Font:    spec.SansBold,
		}
		h.Add(&seg)
		f.heading = &h
	}
	return f.heading
}

func (f *figures) GenerateEntries() {
	f.paper = &abstracts.Paper{}

	var h spec.Addable = f.Heading()
	f.paper.Add(&h)

	spacer := spacing.NewSpacer(f.Heading().Height() + globals.MmToPt(globals.Cfg.Margins.HeadingTop+globals.Cfg.Margins.HeadingBottom))
	var s spec.Addable = spacer
	f.paper.Add(&s)

	figs := make([]*globals.FigureInformation, 0, len(globals.Figures))
	for _, fig := range globals.Figures {
		figs = append(figs, fig)
	}
	sort.Slice(figs, func(i, j int) bool {
		return figs[i].Number < figs[j].Number
	})

	for _, fig := range figs {
		leftSegs := make([]*spec.Segment, 0)
		leftSegs = append(leftSegs, &spec.Segment{
			Content: fmt.Sprintf("%d: %s", fig.Number, fig.Title),
			Font:    spec.SansRegular,
		})
		rightSegs := make([]*spec.Segment, 0)
		rightSegs = append(rightSegs, &spec.Segment{
			Content: fig.Key,
			Font:    spec.SansRegular,
		})
		entry := Entry{
			Left:       leftSegs,
			Right:      rightSegs, // keep this here to replace later... seems very hacky
			LineHeight: globals.Cfg.Tof.LineHeight,
			FontSize:   globals.Cfg.Tof.FontSize,
			Line:       true,
			LeftAlign:  true,
		}
		var a spec.Addable = &entry
		f.paper.Add(&a)
	}
}

func (f *figures) GeneratePages() {
	f.common.GeneratePages()
}

func (f *figures) InsertPageNumbers() {
	for _, page := range f.Pages {
		for _, col := range page.Columns {
			for _, el := range col.Content {
				if entry, ok := (*el).(*Entry); ok {
					kSegs := entry.Right
					entry.Right = make([]*spec.Segment, 0)
					k := kSegs[0].Content // no checks here, because only set above to key; dangerous?
					if fig, ok := globals.Figures[k]; ok {
						for i, p := range fig.Used {
							seg := spec.Segment{Font: spec.SansRegular}
							if i < len(fig.Used)-1 {
								seg.Content = fmt.Sprintf("%d, ", p)
							} else {
								seg.Content = fmt.Sprintf("%d", p)
							}
							entry.Right = append(entry.Right, &seg)
						}
					}
					entry.Process(entry.width)
				}
			}
		}
	}
}

var Figures = &figures{}
