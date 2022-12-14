package spec

import (
	"bytes"
	"fmt"
)

type GraphicLine struct {
	PosA      [2]float64
	PosB      [2]float64
	Thickness float64
	Dotted    bool
}

func (l *GraphicLine) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("q\n")
	buf.WriteString(fmt.Sprintf("%f w\n", l.Thickness))
	if l.Dotted {
		buf.WriteString("[1 2] 0 d\n")
	}
	buf.WriteString(fmt.Sprintf("%f %f m ", l.PosA[0], l.PosA[1]))
	buf.WriteString(fmt.Sprintf("%f %f l S\n", l.PosB[0], l.PosB[1]))
	buf.WriteString("Q\n")
	return buf.Bytes()
}

func (l *GraphicLine) SetPos(x, y float64) {
	l.PosA = [2]float64{x, y}
}

func (l *GraphicLine) Height() float64 {
	return l.Thickness
}

func (l *GraphicLine) Process(_ float64) {
}
