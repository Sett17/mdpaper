package pdf

import "mdpaper/pdf/spec"

type Heading struct {
	spec.Text
	Level int
}

func (h *Heading) SetPos(x float64, y float64) {
	//h.Text.SetPos(x+spec.MmToPt(0+1.5*float64(h.Level)), y-h.Height()/4)
	h.Text.SetPos(x+spec.MmToPt(0+1.5*float64(h.Level)), y+h.Height()/4)
}

func (h *Heading) Bytes() []byte {
	return h.Text.Bytes()
}
