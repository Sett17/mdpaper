package register

import (
	"bytes"
	"fmt"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"math"
)

type Entry struct {
	Left           string
	Right          string
	Pos            [2]float64
	Font           *spec.Font
	FontSize       int
	LineHeight     float64
	Line           bool
	Offset         float64
	LeftAlign      bool
	rightOffset    float64
	leftJustified  spec.JustifiedText
	rightJustified spec.JustifiedText
	line           spec.GraphicLine
	width          float64
}

func (e *Entry) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("BT\n")
	buf.WriteString(fmt.Sprintf("%f %f TD\n", e.Pos[0]+e.Offset, e.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n", e.LineHeight*float64(e.FontSize)))
	buf.Write(e.leftJustified.Bytes(e.FontSize))
	buf.WriteString("ET\n")

	buf.WriteString("BT\n")
	buf.WriteString(fmt.Sprintf("%f %f TD\n", e.Pos[0]+e.rightOffset+e.Offset, e.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n", e.LineHeight*float64(e.FontSize)))
	for i := 0; i < len(e.leftJustified)-1; i++ {
		buf.WriteString("T*\n")
	}
	buf.Write(e.rightJustified.Bytes(e.FontSize))
	buf.WriteString("ET\n")

	if e.Line {
		buf.Write(e.line.Bytes())
	}

	//rect := spec.GraphicRect{
	//	Pos:   [2]float64{e.Pos[0] + e.Offset, e.Pos[1] - e.HeightOffset()},
	//	H:     -e.Height(),
	//	W:     e.rightOffset,
	//	Color: [3]float64{0.0, 0.5, 0.5},
	//}
	//buf.Write(rect.Bytes())

	return buf.Bytes()
}

func (e *Entry) SetPos(x, y float64) {
	e.Pos = [2]float64{x, y}
}

func (e *Entry) Height() float64 {
	return e.LineHeight * float64(e.FontSize) * math.Max(float64(len(e.leftJustified)), float64(len(e.rightJustified)))
}

func (e *Entry) HeightOffset() float64 {
	return e.LineHeight * float64(e.FontSize) * (math.Max(float64(len(e.leftJustified)), float64(len(e.rightJustified))) - 1)
}

func (e *Entry) Process(width float64) {
	e.width = width
	width -= e.Offset

	e.rightOffset = math.Max(width-e.Font.WordWidth(e.Right, e.FontSize)-globals.MmToPt(3), width*.66)

	e.leftJustified = spec.ProcessSegments(
		[]*spec.Segment{{Content: e.Left, Font: e.Font}},
		e.rightOffset-globals.MmToPt(3),
		e.FontSize,
		0,
	)
	e.rightJustified = spec.ProcessSegments(
		[]*spec.Segment{{Content: e.Right, Font: e.Font}},
		width-e.rightOffset,
		e.FontSize,
		0)

	if e.LeftAlign {
		for _, l := range e.leftJustified {
			l.WordSpacing = 0
		}
		for _, l := range e.rightJustified {
			l.WordSpacing = 0
		}
	}

	if e.Line {
		lineY := e.LineHeight * float64(e.FontSize) * float64(len(e.leftJustified)-1)
		e.line = spec.GraphicLine{
			PosA:   [2]float64{e.Font.WordWidth(e.leftJustified[len(e.leftJustified)-1].String(), e.FontSize) + e.Offset + e.Pos[0] + globals.MmToPt(1), e.Pos[1] - lineY},
			PosB:   [2]float64{e.rightOffset + e.Offset + e.Pos[0] - globals.MmToPt(1), e.Pos[1] - lineY},
			Dotted: true,
		}
	}
}
