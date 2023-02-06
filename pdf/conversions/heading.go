package conversions

import (
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"math"
	"strings"
)

func Heading(h *ast.Heading) *spec.Addable {
	buf := strings.Builder{}
	buf.Write(h.Text(globals.File))
	seg := spec.Segment{
		Content: buf.String(),
		Font:    spec.SansBold,
	}
	fs := globals.Cfg.Text.FontSize
	fs = int(math.Max(float64(fs+1), float64(fs)*(1+(float64(6-h.Level)/10))))
	para := elements.Heading{
		Text: spec.Text{
			FontSize:   fs,
			LineHeight: 1.0,
			Offset:     globals.MmToPt(0 + 1.5*float64(h.Level)),
		},
		Level: h.Level,
	}
	para.Add(&seg)
	var out spec.Addable = &para
	return &out
}
