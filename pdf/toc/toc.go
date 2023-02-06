package toc

import (
	"bytes"
	"fmt"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"strings"
)

func GenerateTOC(tree *ChapterTree) Toc {
	ret := make([]*tocEntry, 0)
	for _, e := range *tree {
		te := tocEntry{
			Head:   e.Heading,
			Font:   spec.SansRegular,
			Offset: globals.MmToPt(float64((e.Heading.Level - 1) * 10)),
		}
		ret = append(ret, &te)
	}
	return ret
}

type Toc []*tocEntry

func (t Toc) GenerateLinks() []*spec.DictionaryObject {
	ret := make([]*spec.DictionaryObject, 0)
	for _, e := range t {
		d := e.GenerateLink()
		ret = append(ret, d)
	}
	return ret
}

func (t Toc) GenerateColumn() *abstracts.Column {
	col := abstracts.NewColumn(
		globals.A4Width-(globals.MmToPt(globals.Cfg.Page.MarginHori)*2),
		globals.A4Height-(globals.MmToPt(globals.Cfg.Page.MarginTop)+globals.MmToPt(globals.Cfg.Page.MarginBottom)),
		globals.MmToPt(globals.Cfg.Page.MarginHori),
		globals.A4Height-globals.MmToPt(globals.Cfg.Page.MarginTop),
	)
	headSeg := spec.Segment{
		Content: globals.Cfg.Toc.Heading,
		Font:    spec.SansBold,
	}
	head := elements.Heading{
		Text: spec.Text{
			FontSize:   24,
			LineHeight: 1.5,
		},
		Level: 0,
	}
	head.Add(&headSeg)
	var h spec.Addable = &head
	col.Add(&h)
	for _, e := range t {
		var a spec.Addable = e
		col.Add(&a)
	}
	return col
}

type tocEntry struct {
	Head         *elements.Heading
	Pos          [2]float64
	Offset       float64
	numberOffset float64
	Font         *spec.Font
	//processed    string
	processed []*spec.TextLine
	line      spec.GraphicLine
}

func (t *tocEntry) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("BT\n")

	buf.WriteString(fmt.Sprintf("%f %f TD\n", t.Pos[0]+t.Offset, t.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n", globals.Cfg.Toc.LineHeight*float64(globals.Cfg.Toc.FontSize)))

	buf.WriteString(fmt.Sprintf("/%s %d Tf\n", t.Font.Name, globals.Cfg.Toc.FontSize))
	for _, l := range t.processed {
		buf.WriteString(fmt.Sprintf("("))
		buf.Write(globals.PDFEncode(l.String()))
		buf.WriteString(fmt.Sprintf(") Tj T*\n"))
	}
	buf.WriteString(fmt.Sprintf("%f %f TD\n", t.numberOffset-t.Offset, t.Height()-t.HeightOffset()))
	buf.WriteString(fmt.Sprintf("/%s %d Tf\n", t.Font.Name, globals.Cfg.Toc.FontSize))
	buf.WriteString(fmt.Sprintf("(%s) Tj\n", fmt.Sprintf("%3d", t.Head.DisplayPage)))

	buf.WriteString("ET\n")

	//rect := spec.GraphicRect{
	//	Pos:   [2]float64{t.Pos[0] + t.Offset, t.Pos[1] - t.HeightOffset()},
	//	H:     -t.Height(),
	//	W:     t.numberOffset - t.Offset,
	//	Color: [3]float64{0.0, 0.5, 0.5},
	//}
	//buf.Write(rect.Bytes())

	buf.Write(t.line.Bytes())
	return buf.Bytes()
}

func (t *tocEntry) SetPos(x, y float64) {
	t.Pos[0] = x
	t.Pos[1] = y - globals.MmToPt(10)
}

func (t *tocEntry) Height() float64 {
	return globals.Cfg.Toc.LineHeight * float64(globals.Cfg.Toc.FontSize) * float64(len(t.processed))
}

func (t *tocEntry) HeightOffset() float64 {
	return globals.Cfg.Toc.LineHeight * float64(globals.Cfg.Toc.FontSize) * float64(len(t.processed)-1)
}

func (t *tocEntry) GenerateLink() *spec.DictionaryObject {
	d := spec.NewDictObject()
	d.Set("Type", "/Annot")
	d.Set("Subtype", "/Link")
	// [llx lly urx ury]
	d.Set("Rect", fmt.Sprintf("[%f %f %f %f]", t.Pos[0]+t.Offset, t.Pos[1]-t.HeightOffset(), t.numberOffset+t.Pos[0], t.Pos[1]+t.Height()-t.HeightOffset()))
	d.Set("Border", "[0 0 0]")
	d.Set("Dest", t.Head.Destination())
	return &d
}

func (t *tocEntry) Process(width float64) {
	t.numberOffset = width - t.Font.WordWidth(" 999", globals.Cfg.Toc.FontSize)

	// this is awful
	prevWidth := t.Head.Text.Width
	t.Head.Process(t.numberOffset - t.Offset - globals.MmToPt(2.5))
	t.processed = t.Head.Processed
	t.Head.Process(prevWidth)

	wordsEnd := t.Font.WordWidth(strings.TrimSpace(t.processed[len(t.processed)-1].String()), globals.Cfg.Toc.FontSize) + t.Offset
	lineY := globals.Cfg.Toc.LineHeight * float64(globals.Cfg.Toc.FontSize) * float64(len(t.processed)-1)
	t.line = spec.GraphicLine{
		PosA:   [2]float64{wordsEnd + t.Pos[0] + globals.MmToPt(1), t.Pos[1] - lineY},
		PosB:   [2]float64{t.numberOffset + t.Pos[0] - globals.MmToPt(1), t.Pos[1] - lineY},
		Dotted: true,
	}
}
