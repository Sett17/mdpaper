package pdf

import (
	"bytes"
	"fmt"
	"mdpaper/globals"
	"mdpaper/pdf/spec"
	"strings"
)

func GenerateTOC(tree *ChapterTree) Toc {
	m := spec.MmToPt(15)
	col := NewColumn(globals.A4Width-3*m, globals.A4Height-m, 1.5*m, globals.A4Height-.5*m)
	headSeg := spec.Segment{
		Content: "Table of Contents",
		Font:    &spec.HelveticaBold,
	}
	head := Heading{
		Text: spec.Text{
			FontSize:   24,
			LineHeight: 1.5,
		},
		Level: 0,
	}
	head.Add(&headSeg)
	var h spec.Addable = &head
	col.Add(&h)

	ret := make([]*tocEntry, 0)
	for _, e := range *tree {
		te := tocEntry{
			Head:   e.Heading,
			Font:   &spec.TimesRegular,
			Offset: spec.MmToPt(float64((e.Heading.Level - 1) * 10)),
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

func (t Toc) GenerateColumn() *Column {
	m := spec.MmToPt(15)
	col := NewColumn(globals.A4Width-3*m, globals.A4Height-m, 1.5*m, globals.A4Height-.5*m)
	headSeg := spec.Segment{
		Content: "Table of Contents",
		Font:    &spec.HelveticaBold,
	}
	head := Heading{
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
	Head         *Heading
	Pos          [2]float64
	Offset       float64
	numberOffset float64
	dots         int
	Font         *spec.Font
	processed    string
	line         spec.Line
}

func (t *tocEntry) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("BT\n")

	buf.WriteString(fmt.Sprintf("%f %f TD\n", t.Pos[0]+t.Offset, t.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n 0 Tw\n", globals.Cfg.ToCLineHeight*12))

	buf.WriteString(fmt.Sprintf("/%s %d Tf\n", t.Font.Name, 12))
	buf.WriteString(fmt.Sprintf("(%s) Tj\n", t.processed))
	buf.WriteString(fmt.Sprintf("%f %f TD\n", t.numberOffset-t.Offset, 0.0))
	buf.WriteString(fmt.Sprintf("/%s %d Tf\n", spec.TimesMono.Name, 12))
	buf.WriteString(fmt.Sprintf("(%s) Tj\n", fmt.Sprintf("%3d", t.Head.Page)))

	buf.WriteString("ET\n")
	buf.Write(t.line.Bytes())
	return buf.Bytes()
}

func (t *tocEntry) SetPos(x, y float64) {
	t.Pos[0] = x
	t.Pos[1] = y - spec.MmToPt(10)
}

func (t *tocEntry) Height() float64 {
	return 12 * globals.Cfg.ToCLineHeight
}

func (t *tocEntry) GenerateLink() *spec.DictionaryObject {
	d := spec.NewDictObject()
	d.Set("Type", "/Annot")
	d.Set("Subtype", "/Link")
	d.Set("Rect", fmt.Sprintf("[%f %f %f %f]", t.Pos[0], t.Pos[1]+t.Height(), t.numberOffset+t.Pos[0], t.Pos[1]))
	d.Set("Dest", t.Head.Destination())
	return &d
}

func (t *tocEntry) Process(width float64) {
	t.processed = fmt.Sprintf("%s %s", t.Head.Numbering(), t.Head.String())
	wordsEnd := t.Font.WordWidth(strings.TrimSpace(t.processed), 12) + t.Offset
	t.numberOffset = width - spec.TimesMono.WordWidth(" 999", 12)
	t.line = spec.Line{
		PosA:   [2]float64{wordsEnd + t.Pos[0] + spec.MmToPt(1), t.Pos[1]},
		PosB:   [2]float64{t.numberOffset + t.Pos[0] - spec.MmToPt(1), t.Pos[1]},
		Dotted: true,
	}
}
