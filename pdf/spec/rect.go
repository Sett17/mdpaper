package spec

import (
	"bytes"
	"fmt"
)

type Rect struct {
	Pos [2]float64
	W   float64
	H   float64
}

func (r *Rect) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("q\n")
	buf.WriteString(fmt.Sprintf("%f %f %f %f re s\n", r.Pos[0], r.Pos[1], r.W, -r.H))
	buf.WriteString("Q\n")
	return buf.Bytes()
}

func (r *Rect) SetPos(x, y float64) {
	r.Pos = [2]float64{x, y}
}

func (r *Rect) Height() float64 {
	return 0
}

func (r *Rect) Process(width float64) {
	r.W = width
}
