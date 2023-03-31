package elements

import (
	"bytes"
	"fmt"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"strings"
)

type ListItem struct {
	spec.Text
	Prefix string
}

func (p *ListItem) Process(maxWidth float64) {
	newSeg := spec.Segment{
		Content: p.Prefix,
		Font:    spec.SerifRegular,
	}
	p.Segments = append([]*spec.Segment{&newSeg}, p.Segments...)
	p.Text.Process(maxWidth)
}

func (p *ListItem) Bytes() []byte {
	buf := bytes.Buffer{}

	buf.WriteString("BT\n")
	buf.WriteString(fmt.Sprintf("%f %f TD\n", p.Pos[0], p.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n", float64(p.FontSize)*p.LineHeight))

	// we can assume that paragraph has been processed

	buf.WriteString("T*\n")

	var currFont *spec.Font = nil

	lineBuffer := strings.Builder{}
	l := spec.TextLine{}
	//l.Add(p.Prefix, spec.SerifRegular)
	if l.Offset != 0 {
		buf.WriteString(fmt.Sprintf("%f 0 Td\n", l.Offset))
	}
	buf.WriteString(fmt.Sprintf("%f Tw\n", 0.0))
	for j := 0; j < len(l.Words); j++ {
		if l.Fonts[j] != currFont {
			if lineBuffer.Len() > 0 {
				buf.WriteString(fmt.Sprintf("("))
				buf.Write(globals.WinAnsiEncode(lineBuffer.String()))
				buf.WriteString(fmt.Sprintf(") Tj\n"))
				lineBuffer.Reset()
			}
			buf.WriteString(fmt.Sprintf("/%s %d Tf\n", l.Fonts[j].Name, p.FontSize))
			currFont = l.Fonts[j]
		}
		lineBuffer.WriteString(l.Words[j])
	}
	if lineBuffer.Len() > 0 {
		buf.WriteString(fmt.Sprintf("("))
		buf.Write(globals.WinAnsiEncode(lineBuffer.String()))
		buf.WriteString(fmt.Sprintf(") Tj\n"))
	}

	for i, l := range p.Processed {
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
				buf.WriteString(fmt.Sprintf("/%s %d Tf\n", l.Fonts[j].Name, p.FontSize))
				currFont = l.Fonts[j]
			}
			lineBuffer.WriteString(l.Words[j])
		}
		if lineBuffer.Len() > 0 {
			buf.WriteString(fmt.Sprintf("("))
			buf.Write(globals.WinAnsiEncode(lineBuffer.String()))
			buf.WriteString(fmt.Sprintf(") Tj\n"))
		}
		if i != len(p.Processed)-1 {
			buf.WriteString("T* ")
		}
		if l.Offset != 0 {
			buf.WriteString(fmt.Sprintf("%f 0 Td\n", -l.Offset))
		}
	}

	buf.WriteString("ET\n")

	return buf.Bytes()

}
