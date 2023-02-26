package elements

import (
	"bytes"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/spec"
)

type Table struct {
	Header []*TableCell
	Rows   [][]*TableCell

	Pos     [2]float64
	Padding float64

	lines []*spec.GraphicLine

	cumHeight float64
}

func (t *Table) Bytes() []byte {
	buf := bytes.Buffer{}

	for _, line := range t.lines {
		buf.Write(line.Bytes())
	}

	for _, cell := range t.Header {
		buf.Write(cell.Bytes())
	}

	for _, row := range t.Rows {
		for _, cell := range row {
			buf.Write(cell.Bytes())
		}
	}

	return buf.Bytes()
}

func (t *Table) SetPos(x, y float64) {
	t.Pos = [2]float64{x, y}
}

func (t *Table) Height() float64 {
	return t.cumHeight
}

func (t *Table) Process(width float64) {
	cellWidthNoPadding := width / float64(len(t.Header))
	verticalLines := make([]*spec.GraphicLine, len(t.Header)-1)
	for i, _ := range verticalLines {
		verticalLines[i] = &spec.GraphicLine{
			PosA:      [2]float64{t.Pos[0] + float64(i+1)*cellWidthNoPadding, t.Pos[1]},
			PosB:      [2]float64{},
			Thickness: globals.MmToPt(.25),
		}
	}

	cellWidth := cellWidthNoPadding - 2*t.Padding

	horizontalLinesOffsets := make([]float64, len(t.Rows))

	maxCellHeight := 0.0
	for i, cell := range t.Header {
		cell.Process(cellWidth)
		cell.SetPos(t.Pos[0]+float64(i)*cellWidthNoPadding+t.Padding, t.Pos[1]-t.Padding)
		if cell.Height() > maxCellHeight {
			maxCellHeight = cell.Height()
		}
		for _, l := range cell.Processed {
			l.WordSpacing = 0
		}
	}
	t.cumHeight = maxCellHeight + 2*t.Padding
	horizontalLinesOffsets[0] = t.cumHeight

	for i, row := range t.Rows {
		maxCellHeight := 0.0
		for j, cell := range row {
			cell.Process(cellWidth)
			cell.SetPos(t.Pos[0]+float64(j)*cellWidthNoPadding+t.Padding, t.Pos[1]-t.cumHeight-t.Padding)
			if cell.Height() > maxCellHeight {
				maxCellHeight = cell.Height()
			}
			for _, l := range cell.Processed {
				l.WordSpacing = 0
			}
		}
		t.cumHeight += maxCellHeight + 2*t.Padding
		if i < len(t.Rows)-1 {
			horizontalLinesOffsets[i+1] = t.cumHeight
		}
	}

	for _, l := range verticalLines {
		l.PosB = [2]float64{l.PosA[0], l.PosA[1] - t.cumHeight}
	}

	for i := 0; i < len(t.Rows); i++ {
		t.lines = append(t.lines, &spec.GraphicLine{
			PosA:      [2]float64{t.Pos[0], t.Pos[1] - horizontalLinesOffsets[i] - .25*(globals.Cfg.Text.LineHeight*float64(globals.Cfg.Text.FontSize))},
			PosB:      [2]float64{t.Pos[0] + width, t.Pos[1] - horizontalLinesOffsets[i] - .25*(globals.Cfg.Text.LineHeight*float64(globals.Cfg.Text.FontSize))},
			Thickness: globals.MmToPt(.25),
		})
	}

	t.lines = append(t.lines, verticalLines...)
}
