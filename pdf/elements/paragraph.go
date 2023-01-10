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
	p.Text.Process(width)
	if p.Centered {
		for _, l := range p.Text.Processed {
			l.Center(width)
		}
	}
}

func (p *Paragraph) SetPos(x, y float64) {
	p.Pos = [2]float64{x, y - globals.MmToPt(globals.Cfg.Margins.Paragraph)}
}

func (p *Paragraph) Height() float64 {
	return (float64(len(p.Processed)))*p.LineHeight*float64(p.FontSize) + globals.MmToPt(globals.Cfg.Margins.Paragraph*2)
}
