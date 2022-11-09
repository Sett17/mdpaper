package spec

import (
	"bytes"
	"fmt"
	"math"
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

type line struct {
	Words       []string
	Fonts       []*Font
	WordSpacing float64
	width       float64
}

func (l *line) Add(str string, font *Font) {
	//s := strings.NewReplacer("(", "\\(", ")", "\\)", "\\", "\\\\").Replace(str)
	s := strings.ReplaceAll(str, "\\", "\\\\")
	s = strings.ReplaceAll(s, "(", "\\(")
	s = strings.ReplaceAll(s, ")", "\\)")
	l.Words = append(l.Words, s)
	//l.Words = append(l.Words, str)
	l.Fonts = append(l.Fonts, font)
}

func (l *line) CalculateSpacing(maxWidth float64) {
	buf := strings.Builder{}
	for _, w := range l.Words {
		buf.WriteString(w)
	}
	spaces := float64(strings.Count(buf.String(), " "))
	if spaces == 0 {
		l.WordSpacing = 1.0
		return
	}
	diff := maxWidth - l.width
	l.WordSpacing = diff / spaces
}

type Text struct {
	Segments   []*Segment
	Pos        [2]float64
	FontSize   int
	LineHeight float64
	Processed  []line
	SingleLine bool
}

func (p *Text) Split(percent float64) (Addable, Addable) {
	procCutoff := int(math.Floor(float64(len(p.Processed)) * percent))
	cutoffText := strings.Join(p.Processed[procCutoff].Words, "")
	var segsAfter []*Segment
	var leftoverSegs []*Segment
	for i, segment := range p.Segments {
		if strings.Contains(segment.Content, cutoffText) {
			// Split the segment
			split := strings.Split(segment.Content, cutoffText)
			split[1] = strings.TrimLeft(split[1], " ")
			p.Segments[i].Content = split[0] + cutoffText
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
		SingleLine: p.SingleLine,
	}
	a2 := &Text{
		Segments:   leftoverSegs,
		FontSize:   p.FontSize,
		LineHeight: p.LineHeight,
		Processed:  []line{},
		SingleLine: p.SingleLine,
	}
	return a1, a2
}

func (p *Text) SetPos(x, y float64) {
	p.Pos = [2]float64{x, y}
}

func (p *Text) Height() float64 {
	if p.SingleLine {
		return float64(p.FontSize) * p.LineHeight
	}
	return (float64(len(p.Processed)) + .5) * p.LineHeight * float64(p.FontSize)
}

func (p *Text) Process(maxWidth float64) {
	p.Processed = make([]line, 0)
	if len(p.Segments) == 1 {
		if p.Segments[0].Font.WordWidth(p.Segments[0].Content, p.FontSize) <= maxWidth {
			p.SingleLine = true
		}
	}
	if p.SingleLine {
		l := line{WordSpacing: 1.0}
		for _, s := range p.Segments {
			for i, w := range strings.Split(s.Content, " ") {
				if i != 0 {
					l.Add(" ", s.Font)
				}
				l.Add(w, s.Font)
			}
		}
		p.Processed = append(p.Processed, l)
		return
	}

	//for _, s := range p.Segments {
	//for {
	l := line{WordSpacing: 1.0}
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
			if s.Font.WordWidth(w, p.FontSize)+l.width <= maxWidth {
				suffix := ""
				if j != len(split)-1 {
					suffix = " "
				}
				l.width += s.Font.WordWidth(w+suffix, p.FontSize)
				l.Add(w+suffix, s.Font)
				j++
			} else {
				l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
				l.CalculateSpacing(maxWidth)
				p.Processed = append(p.Processed, l)
				l = line{WordSpacing: 1.0}
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

	buf.WriteString(fmt.Sprintf("%f %f TD\n", p.Pos[0], p.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n", p.LineHeight*float64(p.FontSize)))

	// we can assume that paragraph has been processed

	buf.WriteString("T*\n")

	for i, l := range p.Processed {
		var currFont *Font = nil
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
			buf.WriteString("T*\n")
		}
	}

	buf.WriteString("ET\n")
	return buf.Bytes()
}

//type Text struct {
//	FontSize     int
//	Content      string
//	LineHeight   float64
//	Pos          [2]float64
//	Font         *Font
//	Processed    []line
//	NewParagraph bool
//	SingleLine   bool
//}
//
//func (t *Text) SetPos(x, y float64) {
//	t.Pos = [2]float64{x, y}
//}
//
//func (t *Text) Bytes() []byte {
//	buf := bytes.Buffer{}
//
//	buf.WriteString("BT\n")
//
//	buf.WriteString(fmt.Sprintf("/%s %d Tf\n", t.Font.Name, t.FontSize))
//	buf.WriteString(fmt.Sprintf("%f %f TD\n", t.Pos[0], t.Pos[1]))
//
//	//buf.WriteString(fmt.Sprintf("%f Tc\n", global.CharacterSpacing))
//
//	buf.WriteString(fmt.Sprintf("%f TL\n", t.LineHeight*float64(t.FontSize)))
//
//	//t.Process()
//
//	for _, l := range t.Processed {
//		l.escape()
//		buf.WriteString(fmt.Sprintf("T*\n%f Tw\n(%s) Tj\n", l.WordSpacing, l.Text))
//	}
//
//	buf.WriteString("ET")
//	return buf.Bytes()
//}
//
//func (t *Text) Process(width, offset float64) {
//	split := strings.Split(t.Content, " ")
//	if t.SingleLine || len(split) == 1 {
//		t.Processed = []line{{Text: t.Content, WordSpacing: 0}}
//		return
//	}
//
//	//greedy algorithm for now
//	t.Processed = []line{}
//	var l line
//	for i := 0; i < len(split); {
//		s := split[i]
//		w := width
//		if i == 0 {
//			w -= offset
//		}
//		if l.width+t.WordWidth(s) < w {
//			l.Text += s + " "
//			l.width = t.WordWidth(l.Text)
//			i++
//		} else {
//			if len(l.Text) > 0 {
//				l.Text = l.Text[:len(l.Text)-1]
//			}
//			l.width = t.WordWidth(l.Text)
//
//			spaces := float64(strings.Count(l.Text, " "))
//			diff := w - l.width
//			l.WordSpacing = diff / spaces
//
//			t.Processed = append(t.Processed, l)
//			l = line{}
//			//
//			i++
//		}
//	}
//	if len(l.Text) > 0 {
//		l.Text = l.Text[:len(l.Text)-1]
//	}
//	l.width = t.WordWidth(l.Text)
//	t.Processed = append(t.Processed, l)
//	l = line{}
//
//	//global.Log("Processed text into %d lines", len(t.Processed))
//}
//
//func (t *Text) WordWidth(w string) (width float64) {
//	for _, r := range w {
//		width += t.Font.CharWidth(r) / 1000 * float64(t.FontSize)
//	}
//	return
//}
//
//func (t *Text) Height() float64 {
//	if t.SingleLine {
//		return float64(t.FontSize) * t.LineHeight
//	}
//	return (float64(len(t.Processed)) + .5) * t.LineHeight * float64(t.FontSize)
//}
//
//func (t *Text) WidthPercent(width float64) float64 {
//	if t.SingleLine {
//		return 1
//	}
//	return t.Processed[len(t.Processed)-1].width / width
//}
//
//func (t *Text) AddToPdf(pdf *PDF, pageContent *Array, _ *Dictionary) {
//	z := NewStreamObject()
//	//z.Deflate = true
//	//z.Write(t.Bytes())
//	pdf.AddObject(z.Pointer())
//	pageContent.Add(z.Reference())
//}
