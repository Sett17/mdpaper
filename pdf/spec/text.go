package spec

import (
	"bytes"
	"fmt"
	"math"
	"mdpaper/globals"
	"strings"
)

//type line struct {
//	Text        string
//	WordSpacing float64
//	width       float64
//}
//
//func (l *line) escape() {
//	l.Text = strings.NewReplacer("(", "\\(", ")", "\\)", "\\", "\\\\").Replace(l.Text)
//}

type Segment struct {
	Content string
	Font    *Font
}

type TextLine struct {
	Words       []string
	Fonts       []*Font
	WordSpacing float64
	Width       float64
}

func (l *TextLine) Add(str string, font *Font) {
	//s := strings.NewReplacer("(", "\\(", ")", "\\)", "\\", "\\\\").Replace(str)
	s := strings.ReplaceAll(str, "\\", "\\\\")
	s = strings.ReplaceAll(s, "(", "\\(")
	s = strings.ReplaceAll(s, ")", "\\)")
	l.Words = append(l.Words, s)
	//l.Words = append(l.Words, str)
	l.Fonts = append(l.Fonts, font)
}

func (l *TextLine) CalculateSpacing(maxWidth float64) {
	buf := strings.Builder{}
	for _, w := range l.Words {
		buf.WriteString(w)
	}
	spaces := float64(strings.Count(buf.String(), " "))
	if spaces == 0 {
		l.WordSpacing = 1.0
		return
	}
	diff := maxWidth - l.Width + globals.MmToPt(.7)
	l.WordSpacing = diff / spaces
}

type Text struct {
	Segments   []*Segment
	Pos        [2]float64
	FontSize   int
	LineHeight float64
	Processed  []TextLine
	Offset     float64
	Width      float64
}

func (p *Text) String() string {
	str := strings.Builder{}
	for _, segment := range p.Segments {
		str.WriteString(segment.Content)
	}
	return str.String()
}

func (p *Text) SplitDelegate(percent float64) (Addable, Addable) {
	procCutoff := int(math.Floor(float64(len(p.Processed)) * percent))
	cutoffText := strings.Join(p.Processed[procCutoff].Words, "")
	var segsAfter []*Segment
	var leftoverSegs []*Segment
	for i, segment := range p.Segments {
		if strings.Contains(segment.Content, cutoffText) {
			// Split the segment
			split := strings.Split(segment.Content, cutoffText)
			split[1] = cutoffText + split[1]
			p.Segments[i].Content = split[0]
			segsAfter = p.Segments[i+1:]
			p.Segments = p.Segments[:i+1]
			leftoverSegs = append(leftoverSegs, &Segment{Content: split[1], Font: segment.Font})
			leftoverSegs = append(leftoverSegs, segsAfter...)
			break
		}
	}
	a1 := &Text{
		Segments:   p.Segments,
		Pos:        p.Pos,
		FontSize:   p.FontSize,
		LineHeight: p.LineHeight,
		Processed:  p.Processed[:procCutoff],
		Offset:     p.Offset,
	}
	a2 := &Text{
		Segments:   leftoverSegs,
		FontSize:   p.FontSize,
		LineHeight: p.LineHeight,
		Processed:  []TextLine{},
		Offset:     p.Offset,
	}
	return a1, a2
}

func (p *Text) SetPos(x, y float64) {
	p.Pos = [2]float64{x, y}
}

func (p *Text) Height() float64 {
	if len(p.Processed) == 1 {
		return float64(p.FontSize) * p.LineHeight
	}
	return (float64(len(p.Processed)) + .5) * p.LineHeight * float64(p.FontSize)
}

func (p *Text) Process(maxWidth float64) {
	p.Processed = make([]TextLine, 0)

	//if len(p.Segments) == 1 {
	//	if p.Segments[0].Font.WordWidth(p.Segments[0].Content, p.FontSize) <= maxWidth {
	//		p.SingleLine = true
	//	}
	//}
	//if p.SingleLine {
	//	l := line{WordSpacing: 1.0}
	//	for _, s := range p.Segments {
	//		for i, w := range strings.Split(s.Content, " ") {
	//			if i != 0 {
	//				l.Add(" ", s.Font)
	//			}
	//			l.Add(w, s.Font)
	//		}
	//		p.Processed = append(p.Processed, l)
	//		l = line{WordSpacing: 1.0}
	//	}
	//	return
	//}

	//for _, s := range p.Segments {
	//for {
	maxWidth -= p.Offset
	p.Width = maxWidth
	l := TextLine{WordSpacing: 1.0}
	for i := 0; i < len(p.Segments); {
		s := p.Segments[i]
		if len(s.Content) == 0 {
			i++
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
				l = TextLine{WordSpacing: 1.0}
			}
		}
		i++
	}
	if len(l.Words) == 0 {
		return
	}
	l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
	//l.CalculateSpacing(maxWidth)
	p.Processed = append(p.Processed, l)
	//}
}

func (p *Text) Add(a *Segment) {
	p.Segments = append(p.Segments, a)
}

func (p *Text) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("BT\n")

	buf.WriteString(fmt.Sprintf("%f %f TD\n", p.Pos[0]+p.Offset, p.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n", p.LineHeight*float64(p.FontSize)))

	// we can assume that paragraph has been processed

	buf.WriteString("T*\n")

	var currFont *Font = nil
	for i, l := range p.Processed {
		lineBuffer := strings.Builder{}
		buf.WriteString(fmt.Sprintf("%f Tw\n", l.WordSpacing))
		for j := 0; j < len(l.Words); j++ {
			if l.Fonts[j] != currFont {
				if lineBuffer.Len() > 0 {
					buf.WriteString(fmt.Sprintf("(%s) Tj\n", lineBuffer.String()))
					lineBuffer.Reset()
				}
				buf.WriteString(fmt.Sprintf("/%s %d Tf\n", l.Fonts[j].Name, p.FontSize))
				currFont = l.Fonts[j]
			}
			lineBuffer.WriteString(l.Words[j])
		}
		if lineBuffer.Len() > 0 {
			buf.WriteString(fmt.Sprintf("(%s) Tj\n", lineBuffer.String()))
		}
		if i != len(p.Processed)-1 {
			buf.WriteString("T* ")
		}
	}

	buf.WriteString("ET\n")
	if globals.Cfg.Debug {
		rect := GraphicRect{
			Pos:   [2]float64{p.Pos[0] + p.Offset, p.Pos[1]},
			W:     p.Width,
			H:     p.Height(),
			Color: [3]float64{0.5, 0.5, 0.0},
		}
		if rect.W == 0 {
			fmt.Println("dikka")
		}
		buf.WriteString("\n")
		buf.Write(rect.Bytes())
	}
	return buf.Bytes()
}
