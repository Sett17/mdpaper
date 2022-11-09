package pdf

import (
	"fmt"
	"mdpaper/globals"
	"mdpaper/pdf/spec"
	"strings"
)

type Paper struct {
	Elements []*spec.Addable
}

//func (p *Paper) String() string {
//	buf := bytes.Buffer{}
//
//	for _, e := range p.Elements {
//		if t, ok := (*e).(*spec.Text); ok {
//			buf.WriteString(fmt.Sprintf("Text: len: %d, font: %s\n", len(t.Content), t.Font.FontName))
//		}
//	}
//
//	return buf.String()
//}

func (p *Paper) Add(add ...*spec.Addable) {
	p.Elements = append(p.Elements, add...)
}

func (p *Paper) Finished() bool {
	return len(p.Elements) == 0
}

func (p *Paper) SingleColumn() *Column {
	return p.nextColumn(
		globals.A4Height-2*globals.Cfg.Margin,
		globals.A4Height-2*globals.Cfg.Margin,
		globals.Cfg.Margin,
		globals.A4Height-globals.Cfg.Margin,
	)
}

func (p *Paper) DoubleColumn() (c1, c2 *Column) {
	width := globals.A4Width/2 - 1.25*globals.Cfg.Margin
	height := globals.A4Height - 3*globals.Cfg.Margin

	top := globals.A4Height - globals.Cfg.Margin*1.5

	c1 = p.nextColumn(width, height, globals.Cfg.Margin, top)
	c2 = p.nextColumn(width, height, globals.A4Width/2+globals.Cfg.Margin/4, top)

	return
}

func (p *Paper) nextColumn(width, height, x, y float64) (c *Column) {
	c = NewColumn(width, height, x, y)

	//for i, e := range p.Elements {
	for i := 0; i < len(p.Elements); {
		e := p.Elements[i]
		if spill := c.Add(e); spill != nil {
			p.Elements = p.Elements[i:]
			z := p.Elements
			p.Elements = make([]*spec.Addable, 0)
			p.Elements = append(p.Elements, spill)
			p.Elements = append(p.Elements, z...)
			break
		}
		if i == len(p.Elements)-1 {
			p.Elements = nil
		}
		i++
	}

	return c
}

func (p *Paper) GenerateHeading() {
	levels := []int{0, 0, 0, 0, 0, 0}
	cur := 0
	for _, e := range p.Elements {
		if h, ok := (*e).(*Heading); ok {
			if h.Level-1 > cur {
				if cur == 6 {
					panic("can't have more than 6 heading levels")
				}
				cur = h.Level - 1
			} else if h.Level-1 < cur {
				if cur == 0 {
					panic("can't go under level 0")
				}
				cur = h.Level - 1
			}
			buf := strings.Builder{}
			levels[cur]++
			for i := 0; i <= cur; i++ {
				buf.WriteString(fmt.Sprintf("%d", levels[i]))
				if i != cur {
					buf.WriteString(".")
				}
			}
			buf.WriteString(" ")
			seg := spec.Segment{Content: buf.String(), Font: h.Segments[0].Font}
			segs := h.Segments
			h.Segments = make([]*spec.Segment, 0)
			h.Segments = append(h.Segments, &seg)
			h.Segments = append(h.Segments, segs...)
			h.Process(0)
		}
	}
}
