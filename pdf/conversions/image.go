package conversions

import (
	"fmt"
	"github.com/sett17/mdpaper/cli"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/conversions/options"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"strconv"
	"strings"
)

func Image(image *ast.Image, node ast.Node) (retO *spec.XObject, retA *spec.Addable, retP *spec.Addable) {
	opt, err := options.Parse(string(image.Text(globals.File)))
	if err != nil {
		cli.Error(fmt.Errorf("error parsing image options: %w", err), false)
	}

	mul := .95
	if f, ok := opt.GetFloat("width"); ok {
		mul = f
	}
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
		Content: string(image.Title),
		Font:    spec.SerifRegular,
	})
	var a spec.Addable = &para
	retP = &a
	return
}
