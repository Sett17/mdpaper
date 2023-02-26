package elements

import "github.com/sett17/mdpaper/v2/pdf/spec"

type TableCell struct {
	spec.Text
	Centered bool
}

func (c *TableCell) Process(width float64) {
	c.Text.Process(width)
	if c.Centered {
		for _, l := range c.Text.Processed {
			l.Center(width)
		}
	}
}
