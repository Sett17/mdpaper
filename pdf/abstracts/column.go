package abstracts

import (
	"github.com/sett17/mdpaper/v2/pdf/spacing"
	"github.com/sett17/mdpaper/v2/pdf/spec"
)

type Column struct {
	spec.StreamObject
	Width     float64
	MaxHeight float64
	Pos       [2]float64
	bottom    float64
}

func NewColumn(width, maxHeight, x, y float64) *Column {
	c := Column{Width: width, MaxHeight: maxHeight, Pos: [2]float64{x, y}, StreamObject: spec.NewStreamObject(), bottom: y}
	//rect := spec.GraphicRect{
	//	Pos: [2]float64{x, y - maxHeight},
	//	W:   width,
	//	H:   -maxHeight,
	//}
	//var a spec.Addable = &rect
	//c.StreamObject.Add(&a)
	return &c
}

func (c *Column) Add(a *spec.Addable) (leftover *spec.Addable, full bool) {
	if a == nil || (*a) == nil {
		return nil, false
	}
	A := *a
	if _, ok := A.(*spacing.Filler); ok {
		c.bottom = c.Pos[1] - c.MaxHeight
		return nil, true
	}
	A.SetPos(c.Pos[0], c.bottom)
	A.Process(c.Width)
	h := A.Height()
	if c.bottom-h < c.Pos[1]-c.MaxHeight {
		availSpace := c.bottom - (c.Pos[1] - c.MaxHeight)
		if s, ok := A.(spec.Splittable); ok {
			fitting, leftoverS := s.Split(availSpace / h)

			if fitting != nil {
				c.StreamObject.Add(&fitting)
			}
			return &leftoverS, true
		} else {
			return a, true
		}
	}
	c.StreamObject.Add(a)
	c.bottom -= h
	return nil, false
}
