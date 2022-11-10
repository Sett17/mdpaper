package pdf

import (
	"fmt"
	"mdpaper/globals"
	"mdpaper/pdf/spec"
	"strings"
)

type Heading struct {
	spec.Text
	Level  int
	Page   int
	Prefix [6]int
}

func (h *Heading) Destination() string {
	X := 0.0
	if h.Text.Pos[0] > globals.A4Width/2 {
		X = globals.A4Width / 2
	}
	return fmt.Sprintf("[%d /XYZ %f %f 0]", h.Page-1, X, h.Text.Pos[1])
}

func (h *Heading) String() string {
	return h.Text.String()
}

func (h *Heading) Numbering() string {
	if h.Prefix != [6]int{} {
		p := strings.Builder{}
		for _, n := range h.Prefix {
			if n != 0 {
				p.WriteString(fmt.Sprintf("%d.", n))
			}
		}
		return p.String()[0 : len(p.String())-1]
	}
	return "x.x.x"
}

func (h *Heading) SetPos(x float64, y float64) {
	//h.Text.SetPos(x+spec.MmToPt(0+1.5*float64(h.Level)), y-h.Height()/4)
	h.Text.SetPos(x+spec.MmToPt(0+1.5*float64(h.Level)), y+h.Height()/4)
}

func (h *Heading) Bytes() []byte {
	if h.Prefix != [6]int{} {
		h.Text.Processed[0].Words[0] = h.Numbering() + " " + h.Text.Processed[0].Words[0]
	}
	return h.Text.Bytes()
}
