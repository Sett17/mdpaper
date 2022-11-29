package pdf

import (
	"github.com/sett17/mdpaper/pdf/spec"
	"strings"
)

type List struct {
	spec.Text
}

//TODO make splittable
//func (p *List) Split(percent float64) (spec.Addable, spec.Addable) {
//
//}

func (p *List) Process(maxWidth float64) {
	p.Processed = make([]*spec.TextLine, 0)
	maxWidth -= p.Offset
	p.Width = maxWidth
	l := &spec.TextLine{WordSpacing: 1.0}
	for _, s := range p.Segments {
		if len(s.Content) == 0 {
			continue
		}
		split := strings.Split(s.Content, " ")
		for j := 0; j < len(split); {
			w := split[j]
			if j != 0 && strings.TrimSpace(w) == "" {
				j++
				continue
			}
			if s.Font.WordWidth(w, p.FontSize)+l.Width <= maxWidth {
				suffix := ""
				if j != len(split)-1 {
					suffix = " "
				}
				l.Width += s.Font.WordWidth(w+suffix, p.FontSize)
				l.Add(w+suffix, s.Font)
				j++
			} else {
				if l.Width != 0 {
					l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
				}
				l.CalculateSpacing(maxWidth)
				p.Processed = append(p.Processed, l)
				l = &spec.TextLine{WordSpacing: 1.0}
			}
		}
		p.Processed = append(p.Processed, l)
		l = &spec.TextLine{WordSpacing: 1.0}
	}
}
