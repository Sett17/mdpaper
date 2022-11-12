package pdf

import (
	"mdpaper/globals"
	"mdpaper/pdf/spec"
)

type Column struct {
	spec.StreamObject
	Width     float64
	MaxHeight float64
	Pos       [2]float64
	Full      bool
	bottom    float64
}

func NewColumn(width, maxHeight, x, y float64) *Column {
	c := Column{Width: width, MaxHeight: maxHeight, Pos: [2]float64{x, y}, StreamObject: spec.NewStreamObject(), bottom: y}

	if globals.Cfg.Debug {
		r := spec.GraphicRect{
			Pos:   [2]float64{c.Pos[0], c.Pos[1]},
			W:     c.Width,
			H:     c.MaxHeight,
			Color: [3]float64{0, 0, 0},
		}
		var rA spec.Addable = &r
		c.Add(&rA)
	}

	return &c
}

func (c *Column) Add(a *spec.Addable) (leftover *spec.Addable) {
	if a == nil {
		return nil
	}
	A := *a
	if _, ok := A.(*Filler); ok {
		c.bottom = c.Pos[1] - c.MaxHeight
		c.Full = true
		return nil
	}
	A.SetPos(c.Pos[0], c.bottom)
	A.Process(c.Width)
	h := A.Height()
	if c.bottom-h < c.Pos[1]-c.MaxHeight {
		availSpace := c.bottom - (c.Pos[1] - c.MaxHeight)
		if s, ok := A.(spec.Splittable); ok {
			fitting, leftoverS := s.Split(availSpace / h)
			//fitting.Process(c.Width)
			c.StreamObject.Add(&fitting)
			c.Full = true
			return &leftoverS
		} else {
			c.Full = true
			return a
		}
	}
	c.StreamObject.Add(a)
	c.bottom -= h
	return nil
}
