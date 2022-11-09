package pdf

import (
	"fmt"
	"mdpaper/pdf/spec"
	"strings"
)

type Heading struct {
	spec.Text
	Level  int
	Page   int
	Prefix [6]int
}

func (h *Heading) String() string {
	return fmt.Sprintf("Heading {lvl: %d page: %d content: %s}", h.Level, h.Page, h.Text.String())
}

func (h *Heading) SetPos(x float64, y float64) {
	//h.Text.SetPos(x+spec.MmToPt(0+1.5*float64(h.Level)), y-h.Height()/4)
	h.Text.SetPos(x+spec.MmToPt(0+1.5*float64(h.Level)), y+h.Height()/4)
}

func (h *Heading) Bytes() []byte {
	p := strings.Builder{}
	for _, n := range h.Prefix {
		if n != 0 {
			p.WriteString(fmt.Sprintf("%d.", n))
		}
	}
	h.Text.Processed[0].Words[0] = p.String()[0:len(p.String())-1] + " " + h.Text.Processed[0].Words[0]
	return h.Text.Bytes()
}
