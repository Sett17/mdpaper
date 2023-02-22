package spec

import (
	"bytes"
	"fmt"
	"github.com/sett17/mdpaper/v2/globals"
	"strings"
)

type JustifiedText []*TextLine

func (jt JustifiedText) Bytes(fs int) []byte {
	buf := bytes.Buffer{}
	var currFont *Font = nil
	for i, l := range jt {
		lineBuffer := strings.Builder{}
		if l.Offset != 0 {
			buf.WriteString(fmt.Sprintf("%f 0 Td\n", l.Offset))
		}
		buf.WriteString(fmt.Sprintf("%f Tw\n", l.WordSpacing))
		for j := 0; j < len(l.Words); j++ {
			if l.Fonts[j] != currFont {
				if lineBuffer.Len() > 0 {
					buf.WriteString(fmt.Sprintf("("))
					buf.Write(globals.WinAnsiEncode(lineBuffer.String()))
					buf.WriteString(fmt.Sprintf(") Tj\n"))
					lineBuffer.Reset()
				}
				buf.WriteString(fmt.Sprintf("/%s %d Tf\n", l.Fonts[j].Name, fs))
				currFont = l.Fonts[j]
			}
			lineBuffer.WriteString(l.Words[j])
		}
		if lineBuffer.Len() > 0 {
			buf.WriteString(fmt.Sprintf("("))
			buf.Write(globals.WinAnsiEncode(lineBuffer.String()))
			buf.WriteString(fmt.Sprintf(") Tj\n"))
		}
		if i != len(jt)-1 {
			buf.WriteString("T* ")
		}
		if l.Offset != 0 {
			buf.WriteString(fmt.Sprintf("%f 0 Td\n", -l.Offset))
		}
	}
	return buf.Bytes()
}

type TextLine struct {
	Words       []string
	Fonts       []*Font
	WordSpacing float64
	Width       float64
	Offset      float64
}

func escape(str string) (ret string) {
	ret = strings.ReplaceAll(str, "\\", "\\\\")
	ret = strings.ReplaceAll(ret, "(", "\\(")
	ret = strings.ReplaceAll(ret, ")", "\\)")
	return
}

func deEscape(str string) (ret string) {
	ret = strings.ReplaceAll(str, "\\\\", "\\")
	ret = strings.ReplaceAll(ret, "\\(", "(")
	ret = strings.ReplaceAll(ret, "\\)", ")")
	return
}

func (l *TextLine) Add(str string, font *Font) {
	s := escape(str)

	l.Words = append(l.Words, s)
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

func (l *TextLine) Center(fullWidth float64) {
	l.WordSpacing = 0
	diff := fullWidth - l.Width
	l.Offset = diff / 2
}

func (l *TextLine) String() string {
	return strings.Join(l.Words, " ")
}

func ProcessSegments(segs []*Segment, width float64, fontSize int, offset float64) JustifiedText {
	ret := make([]*TextLine, 0)

	width -= offset / 2
	l := &TextLine{WordSpacing: 1.0, Offset: offset}
	for i := 0; i < len(segs); {
		s := segs[i]
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
			if s.Font.WordWidth(w, fontSize) > width {
				if width-l.Width > s.Font.WordWidth("xxx", fontSize) {
					pre, rem := SplitLongString(w, width-l.Width, fontSize, s.Font)
					split = append(split[:j+1], append([]string{rem}, split[j+1:]...)...)
					if pre == "" {
						j++
						continue
					}
					w = pre
					split[j] = w
				}
			}
			if s.Font.WordWidth(w, fontSize)+l.Width <= width {
				suffix := ""
				if j != len(split)-1 {
					suffix = " "
				}
				l.Width += s.Font.WordWidth(w+suffix, fontSize)
				l.Add(w+suffix, s.Font)
				j++
			} else {
				if l.Width != 0 {
					l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
				}
				l.CalculateSpacing(width)
				ret = append(ret, l)
				l = &TextLine{WordSpacing: 1.0, Offset: offset}
			}
		}
		i++
	}
	if len(l.Words) == 0 {
		return ret
	}
	l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
	ret = append(ret, l)
	return ret
}
