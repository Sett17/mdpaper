package spec

import (
	"bytes"
	"fmt"
)

type GraphicRect struct {
	Pos   [2]float64
	W     float64
	H     float64
	Color [3]float64
}

func (r *GraphicRect) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("q\n")
	buf.WriteString(fmt.Sprintf("%f %f %f RG\n", r.Color[0], r.Color[1], r.Color[2]))
	buf.WriteString(fmt.Sprintf("%f %f %f %f re s\n", r.Pos[0], r.Pos[1], r.W, -r.H))
	buf.WriteString("Q\n")
	return buf.Bytes()
}

func (r *GraphicRect) SetPos(x, y float64) {
	r.Pos = [2]float64{x, y}
}

func (r *GraphicRect) Height() float64 {
	return 0
}

func (r *GraphicRect) Process(width float64) {
	r.W = width
}
