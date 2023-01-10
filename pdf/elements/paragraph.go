package elements

import (
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/spec"
)

type Paragraph struct {
	spec.Text
	Centered bool
}

func (p *Paragraph) Split(percent float64) (r1 spec.Addable, r2 spec.Addable) {
	if len(p.Processed) == 1 {
		return nil, p
	}
	a1, a2 := p.Text.SplitDelegate(percent)
	if a1 != nil {
		r1 = &Paragraph{Text: *a1.(*spec.Text)}
	}
	if a2 != nil {
		r2 = &Paragraph{Text: *a2.(*spec.Text)}
	}
	return
}

func (p *Paragraph) Process(width float64) {
	if p.Centered {
		p.Text.Process(width)
		for _, l := range p.Text.Processed {
			l.Center(width)
		}
	} else {
		p.Text.Process(width)
	}
}

func (p *Paragraph) Height() float64 {
	return (float64(len(p.Processed)))*p.LineHeight*float64(p.FontSize) + globals.MmToPt(globals.Cfg.Spaces.Paragraph)
}
