package pdf

import "mdpaper/pdf/spec"

type Paragraph struct {
	spec.Text
	Centered bool
}

func (p *Paragraph) Split(percent float64) (spec.Addable, spec.Addable) {
	a1, a2 := p.Text.SplitDelegate(percent)
	var r1 = Paragraph{Text: *a1.(*spec.Text)}
	var r2 = Paragraph{Text: *a2.(*spec.Text)}
	return &r1, &r2
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
