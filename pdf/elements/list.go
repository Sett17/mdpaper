package elements

import (
	"github.com/sett17/mdpaper/pdf/spec"
	"math"
	"strings"
)

type List struct {
	spec.Text
}

func (p *List) Split(percent float64) (spec.Addable, spec.Addable) {
	segCutoff := int(math.Max(float64(len(p.Segments))*percent, 1))
	segsAfter := p.Segments[segCutoff-1:]
	a1 := &List{
		Text: spec.Text{
			Segments:   p.Segments[:segCutoff-1],
			Pos:        p.Pos,
			FontSize:   p.FontSize,
			LineHeight: p.LineHeight,
			Processed:  make([]*spec.TextLine, 0),
			Offset:     p.Offset},
	}
	a1.Process(p.Width)
	a2 := &List{
		Text: spec.Text{
			Segments:   segsAfter,
			FontSize:   p.FontSize,
			LineHeight: p.LineHeight,
			Processed:  make([]*spec.TextLine, 0),
			Offset:     p.Offset},
	}
	return a1, a2
}

func (p *List) Process(maxWidth float64) {
	p.Processed = make([]*spec.TextLine, 0)
	maxWidth -= p.Offset
	p.Width = maxWidth
	l := &spec.TextLine{WordSpacing: 1.0}
	for _, s := range p.Segments {
		if len(s.Content) == 0 {
			continue
		}
		splitSmall := strings.Split(s.Content, " ")
		split := make([]string, 0)
		for _, s := range splitSmall {
			if len(s) > 0 {
				split = append(split, strings.SplitAfter(s, "/")...)
			}
		}
		for j := 0; j < len(split); {
			w := split[j]
			if j != 0 && strings.TrimSpace(w) == "" {
				j++
				continue
			}
			if s.Font.WordWidth(w, p.FontSize)+l.Width <= maxWidth {
				suffix := ""
				if j != len(split)-1 && !strings.HasSuffix(w, "/") {
					suffix = " "
				}
				l.Width += s.Font.WordWidth(w+suffix, p.FontSize)
				l.Add(w+suffix, s.Font)
				j++
			} else {
				if l.Width != 0 {
					l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
				}
				//l.CalculateSpacing(maxWidth)
				p.Processed = append(p.Processed, l)
				l = &spec.TextLine{WordSpacing: 1.0}
			}
		}
		p.Processed = append(p.Processed, l)
		l = &spec.TextLine{WordSpacing: 1.0}
	}
}
