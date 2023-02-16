package conversions

import (
	"fmt"
	"github.com/goccy/go-graphviz"
	"github.com/golang/freetype/truetype"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/conversions/options"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"golang.org/x/image/font"
	"image"
	"math/rand"
	"strconv"
)

func Dot(fcb *ast.FencedCodeBlock) (retO *spec.XObject, retA *spec.Addable, retP *spec.Addable) {
	optionString := ""
	if fcb.Info != nil {
		optionString = options.Extract(string(fcb.Info.Text(globals.File)))
	}
	opts, err := options.Parse(optionString)
	if err != nil {
		cli.Error(fmt.Errorf("error parsing graphviz options: %w", err), false)
		cli.Warning(optionString)
	}

	mul := 1.0
	if f, ok := opts.GetFloat("width"); ok {
		mul = f
	}
	title := ""
	if t, ok := opts.GetString("title"); ok {
		title = t
	} else if t, ok := opts.GetString("caption"); ok {
		title = t
	} else if t, ok := opts.GetString("label"); ok {
		title = t
	}
	dpi := 331.0
	if d, ok := opts.GetFloat("dpi"); ok {
		dpi = d
	}

	buf := make([]byte, 0)
	for i := 0; i < fcb.Lines().Len(); i++ {
		at := fcb.Lines().At(i)
		buf = append(buf, at.Value(globals.File)...)
	}
	graph, err := graphviz.ParseBytes(buf)
	if err != nil {
		cli.Error(fmt.Errorf("error parsing graphviz code: %w", err), false)
		return
	}

	var img image.Image
	if graph == nil {
		cli.Error(fmt.Errorf("error parsing graphviz code"), false)
		r := spec.FillingRect{
			GraphicRect: spec.GraphicRect{
				Pos:           [2]float64{},
				RoundedBottom: true,
				RoundedTop:    true,
			},
			Ratio: 1.5,
			Mul:   mul,
		}
		var R spec.Addable = &r
		retA = &R
	} else {
		g := graphviz.New()

		//https://github.com/goccy/go-graphviz/issues/63
		g.SetFontFace(func(size float64) (font.Face, error) {
			return truetype.NewFace(spec.SerifRegular.Font, &truetype.Options{Size: size * .91}), nil
		})

		graph.SetDPI(dpi)
		img, err = g.RenderImage(graph)
		if err != nil {
			cli.Error(fmt.Errorf("error rendering graphviz code: %w", err), false)
			return
		}

		io, ia := spec.NewImageObject(img, strconv.Itoa(rand.Int()), mul)
		retO = &io
		retA = &ia
	}

	if title != "" {
		para := elements.Paragraph{
			Text: spec.Text{
				FontSize:   globals.Cfg.Text.FontSize - 1,
				LineHeight: 1.15,
			},
			Centered: true,
		}
		para.Add(&spec.Segment{
			Content: title,
			Font:    spec.SerifRegular,
		})
		var a spec.Addable = &para
		retP = &a
	}

	return
}
