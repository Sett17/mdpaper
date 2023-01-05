package conversions

import (
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
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
	if h.Level <= 2 {
		fs = int(float64(fs) * 1.2)
	}
	para := elements.Heading{
		Text: spec.Text{
			FontSize: fs,
			//LineHeight: globals.Cfg.LineHeight * 1.5,
			LineHeight: 1.0,
			Offset:     globals.MmToPt(0 + 1.5*float64(h.Level)),
		},
		Level: h.Level,
	}
	para.Add(&seg)
	var out spec.Addable = &para
	return &out
}
