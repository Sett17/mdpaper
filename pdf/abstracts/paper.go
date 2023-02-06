package abstracts

import (
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/spec"
)

type Paper struct {
	Elements []*spec.Addable
	XObjects []*spec.XObject
}

func (p *Paper) Add(add ...*spec.Addable) {
	p.Elements = append(p.Elements, add...)
}

func (p *Paper) AddXObject(x ...*spec.XObject) {
	p.XObjects = append(p.XObjects, x...)
}

func (p *Paper) Finished() bool {
	return len(p.Elements) == 0
}

func (p *Paper) SingleColumn() *Column {
	return p.nextColumn(
		globals.A4Width-(globals.MmToPt(globals.Cfg.Page.MarginHori)*2),
		globals.A4Height-(globals.MmToPt(globals.Cfg.Page.MarginTop)+globals.MmToPt(globals.Cfg.Page.MarginBottom)),
		globals.MmToPt(globals.Cfg.Page.MarginHori),
		globals.A4Height-globals.MmToPt(globals.Cfg.Page.MarginTop),
	)
}

func (p *Paper) DoubleColumn() (c1, c2 *Column) {
	width := globals.A4Width/2 - globals.MmToPt(globals.Cfg.Page.MarginHori+globals.Cfg.Page.ColumnGap/2)
	height := globals.A4Height - globals.MmToPt(globals.Cfg.Page.MarginTop+globals.Cfg.Page.MarginBottom)

	top := globals.A4Height - globals.MmToPt(globals.Cfg.Page.MarginTop)
	c1 = p.nextColumn(
		width,
		height,
		globals.MmToPt(globals.Cfg.Page.MarginHori),
		top,
	)
	c2 = p.nextColumn(
		width,
		height,
		globals.A4Width/2+globals.MmToPt(globals.Cfg.Page.ColumnGap/2),
		top,
	)

	return
}

func (p *Paper) nextColumn(width, height, x, y float64) (c *Column) {
	c = NewColumn(width, height, x, y)

	for i := 0; i < len(p.Elements); {
		e := p.Elements[i]
		if spill, full := c.Add(e); spill != nil || full {
			p.Elements = p.Elements[i+1:]
			z := p.Elements
			p.Elements = make([]*spec.Addable, 0)
			p.Elements = append(p.Elements, spill)
			p.Elements = append(p.Elements, z...)
			break
		}
		if i == len(p.Elements)-1 {
			p.Elements = nil
		}
		i++
	}

	return c
}
