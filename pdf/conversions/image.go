package conversions

import (
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"strconv"
	"strings"
)

func Image(image *ast.Image, node ast.Node) (retO *spec.XObject, retA *spec.Addable, retP *spec.Addable) {
	mul := .95
	if node.ChildCount() == 2 {
		mul, _ = strconv.ParseFloat(strings.TrimSpace(string(node.FirstChild().NextSibling().Text(globals.File))), 64)
	}
	io, ia := spec.NewImageObjectFromFile(string(image.Destination), mul)
	retO = &io
	retA = &ia
	para := elements.Paragraph{
		Text: spec.Text{
			FontSize:   globals.Cfg.Text.FontSize - 1,
			LineHeight: 1.0,
		},
		Centered: true,
	}
	para.Add(&spec.Segment{
		Content: string(image.Text(globals.File)),
		Font:    spec.SerifRegular,
	})
	var a spec.Addable = &para
	retP = &a
	return
}
