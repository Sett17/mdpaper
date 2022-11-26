package spec

import (
	"bytes"
	"fmt"
)

type GraphicRect struct {
	Pos     [2]float64
	W       float64
	H       float64
	Color   [3]float64
	Filled  bool
	Rounded bool
}

const radius = 8

func (r *GraphicRect) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("q\n")
	buf.WriteString(fmt.Sprintf("%f %f %f RG\n", r.Color[0], r.Color[1], r.Color[2]))
	if r.Rounded {
		buf.WriteString(fmt.Sprintf("%f %f m ", r.Pos[0]+radius, r.Pos[1]))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+r.W-radius, r.Pos[1]))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0]+r.W, r.Pos[1], r.Pos[0]+r.W, r.Pos[1], r.Pos[0]+r.W, r.Pos[1]-radius))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+r.W, r.Pos[1]-r.H+radius))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0]+r.W, r.Pos[1]-r.H, r.Pos[0]+r.W, r.Pos[1]-r.H, r.Pos[0]+r.W-radius, r.Pos[1]-r.H))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+radius, r.Pos[1]-r.H))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0], r.Pos[1]-r.H, r.Pos[0], r.Pos[1]-r.H, r.Pos[0], r.Pos[1]-r.H+radius))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0], r.Pos[1]-radius))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c\n", r.Pos[0], r.Pos[1], r.Pos[0], r.Pos[1], r.Pos[0]+radius, r.Pos[1]))
	} else {
		buf.WriteString(fmt.Sprintf("%f %f %f %f re\n", r.Pos[0], r.Pos[1], r.W, -r.H))
	}

	if r.Filled {
		buf.WriteString("f\n")
	} else {
		buf.WriteString("s\n")
	}
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
