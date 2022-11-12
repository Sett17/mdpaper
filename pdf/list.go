package pdf

import (
	"mdpaper/pdf/spec"
	"strings"
)

type List struct {
	spec.Text
}

func (p *List) Process(maxWidth float64) {
	maxWidth -= p.Offset
	l := spec.TextLine{WordSpacing: 1.0}
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
				l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
				l.CalculateSpacing(maxWidth)
				p.Processed = append(p.Processed, l)
				l = spec.TextLine{WordSpacing: 1.0}
			}
		}
		p.Processed = append(p.Processed, l)
		l = spec.TextLine{WordSpacing: 1.0}
	}
}
