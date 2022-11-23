package pdf

import "mdpaper/pdf/spec"

type Paragraph struct {
	spec.Text
}

func (p *Paragraph) Split(percent float64) (spec.Addable, spec.Addable) {
	a1, a2 := p.Text.SplitDelegate(percent)
	var r1 = Paragraph{Text: *a1.(*spec.Text)}
	var r2 = Paragraph{Text: *a2.(*spec.Text)}
	return &r1, &r2
}
