package spec

import (
	"bytes"
	"fmt"
)

type GraphicRect struct {
	Pos           [2]float64
	W             float64
	H             float64
	Color         [3]float64
	BorderColor   [3]float64
	Filled        bool
	RoundedTop    bool
	RoundedBottom bool
}

const radius = 8

func (r *GraphicRect) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("q\n")
	if r.RoundedTop && r.RoundedBottom {
		buf.WriteString(fmt.Sprintf("%f %f m ", r.Pos[0]+radius, r.Pos[1]))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+r.W-radius, r.Pos[1]))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0]+r.W, r.Pos[1], r.Pos[0]+r.W, r.Pos[1], r.Pos[0]+r.W, r.Pos[1]-radius))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+r.W, r.Pos[1]-r.H+radius))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0]+r.W, r.Pos[1]-r.H, r.Pos[0]+r.W, r.Pos[1]-r.H, r.Pos[0]+r.W-radius, r.Pos[1]-r.H))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+radius, r.Pos[1]-r.H))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0], r.Pos[1]-r.H, r.Pos[0], r.Pos[1]-r.H, r.Pos[0], r.Pos[1]-r.H+radius))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0], r.Pos[1]-radius))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c\n", r.Pos[0], r.Pos[1], r.Pos[0], r.Pos[1], r.Pos[0]+radius, r.Pos[1]))
	} else if r.RoundedTop {
		buf.WriteString(fmt.Sprintf("%f %f m ", r.Pos[0]+radius, r.Pos[1]))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+r.W-radius, r.Pos[1]))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0]+r.W, r.Pos[1], r.Pos[0]+r.W, r.Pos[1], r.Pos[0]+r.W, r.Pos[1]-radius))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+r.W, r.Pos[1]-r.H))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0], r.Pos[1]-r.H))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0], r.Pos[1]-radius))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c\n", r.Pos[0], r.Pos[1], r.Pos[0], r.Pos[1], r.Pos[0]+radius, r.Pos[1]))
	} else if r.RoundedBottom {
		buf.WriteString(fmt.Sprintf("%f %f m ", r.Pos[0]+radius, r.Pos[1]))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+r.W, r.Pos[1]))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+r.W, r.Pos[1]-r.H+radius))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0]+r.W, r.Pos[1]-r.H, r.Pos[0]+r.W, r.Pos[1]-r.H, r.Pos[0]+r.W-radius, r.Pos[1]-r.H))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0]+radius, r.Pos[1]-r.H))
		buf.WriteString(fmt.Sprintf("%f %f %f %f %f %f c ", r.Pos[0], r.Pos[1]-r.H, r.Pos[0], r.Pos[1]-r.H, r.Pos[0], r.Pos[1]-r.H+radius))
		buf.WriteString(fmt.Sprintf("%f %f l ", r.Pos[0], r.Pos[1]))
	} else {
		buf.WriteString(fmt.Sprintf("%f %f %f %f re\n", r.Pos[0], r.Pos[1], r.W, -r.H))
	}

	buf.WriteString(fmt.Sprintf("%f %f %f rg %f %f %f RG\n", r.Color[0], r.Color[1], r.Color[2], r.BorderColor[0], r.BorderColor[1], r.BorderColor[2]))
	if r.Filled {
		buf.WriteString("b\n")
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

type FillingRect struct {
	GraphicRect
	Ratio float64
	Mul   float64
}

func (f *FillingRect) Height() float64 {
	return f.H
}

func (f *FillingRect) Process(width float64) {
	if f.Ratio == 0 {
		f.Ratio = 2
	}
	f.W = width * f.Mul
	f.H = f.W / f.Ratio
}
