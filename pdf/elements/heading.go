package elements

import (
	"fmt"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/spec"
	"strings"
)

type Heading struct {
	spec.Text
	Level       int
	Page        int
	DisplayPage int
	Prefix      [6]int
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

func (h *Heading) SetPrefix(prefix [6]int) {
	h.Prefix = prefix
	if h.Prefix != [6]int{} {
		h.Text.Segments[0].Content = h.Numbering() + " " + h.Text.Segments[0].Content
	}
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
	h.Text.SetPos(x, y-globals.MmToPt(globals.Cfg.Spaces.Heading)/1.5) //bit below center off whitespace
}

func (h *Heading) Height() float64 {
	return h.Text.Height() + globals.MmToPt(globals.Cfg.Spaces.Heading)
}
