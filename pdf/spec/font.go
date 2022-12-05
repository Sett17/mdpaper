package spec

import (
	"fmt"
	"github.com/beta/freetype/truetype"
	"github.com/sett17/mdpaper/cli"
	"github.com/sett17/mdpaper/globals"
	"golang.org/x/image/font"
	"os"
	"strings"
)

type Font struct {
	Width      map[int]int
	FilePath   string
	Font       *truetype.Font
	Face       font.Face
	Data       *StreamObject
	Name       string
	Descriptor Dictionary
	Bold       bool
	Mono       bool
}

func (f *Font) AddToPDF(p *PDF) (ref string, name string) {
	fontObj := NewDictObject()
	fontObj.Set("Type", "/Font")
	fontObj.Set("Subtype", "/TrueType")
	fontObj.Set("BaseFont", "/"+f.Font.Name(truetype.NameIDPostscriptName))

	fontObj.Set("FirstChar", 32)
	fontObj.Set("LastChar", 255)
	w := NewArrayObject()
	for i := 32; i < 256; i++ {
		w.Items = append(w.Items, int(f.CharWidth(rune(i))*(1000)))
	}
	p.AddObject(w.Pointer())
	fontObj.Set("Widths", w.Reference())

	p.AddObject(fontObj.Pointer())
	descObj := NewDictObject()
	descObj.M = f.Descriptor.M
	p.AddObject(descObj.Pointer())
	fontObj.Set("FontDescriptor", descObj.Reference())

	p.AddObject(f.Data.Pointer())

	return fontObj.Reference(), f.Name
}

func NewFont(filePath string, flags int) (f *Font) {
	f = &Font{}
	fontFileBuf, err := os.ReadFile(filePath)
	if err != nil {
		fontFileBuf, err = globals.Fonts.ReadFile(filePath) // look in embedded fonts if not found
		if err != nil {
			cli.Error(fmt.Errorf("could not read font file %s", filePath), true)
		}
	}
	fontFile, err := truetype.Parse(fontFileBuf)
	if err != nil {
		cli.Error(fmt.Errorf("could not parse font file %s", filePath), true)
	}
	f.Name = strings.ToLower(strings.ReplaceAll(fontFile.Name(truetype.NameIDPostscriptName), " ", "_"))
	f.Font = fontFile
	f.Face = truetype.NewFace(fontFile, &truetype.Options{
		Size: float64(globals.Cfg.Text.FontSize),
	})

	ttfStream := NewStreamObject()
	ttfStream.AlwaysDeflate = true
	ttfStream.Set("Length1", len(fontFileBuf))
	ttfStream.Write(fontFileBuf)
	f.Data = &ttfStream

	desc := generateDescriptor(f, flags)
	desc.Set("FontFile2", ttfStream.Reference())
	f.Descriptor = desc

	f.Bold = flags&0b1000000000000000000 != 0
	f.Mono = flags&0b1 != 0

	return
}

func generateDescriptor(f *Font, flags int) Dictionary {
	d := NewDict()
	d.Set("Type", "/FontDescriptor")
	d.Set("FontName", "/"+f.Font.Name(truetype.NameIDPostscriptName))
	d.Set("Name", "/"+f.Name)
	d.Set("Flags", flags)
	bbox := f.Font.Bounds(1000)
	d.Set("FontBBox", Array{Items: []interface{}{
		int(bbox.Min.X), int(bbox.Min.Y), int(bbox.Max.X), int(bbox.Max.Y),
	}})
	d.Set("Ascent", int(f.Face.Metrics().Ascent))
	d.Set("Descent", int(f.Face.Metrics().Descent))
	slope := f.Face.Metrics().CaretSlope
	angle := 0
	if !(slope.X == 0 && slope.Y == 0) {
		angle = int(float64(slope.X) / float64(slope.Y))
	}
	d.Set("ItalicAngle", angle)
	d.Set("StemV", 80)
	d.Set("CapHeight", int(f.Face.Metrics().CapHeight))
	return d
}

func (f *Font) WordWidth(w string, fs int) float64 {
	var width float64
	for _, r := range w {
		width += f.CharWidth(r) * float64(fs)
	}
	return width
}

func (f *Font) CharWidth(r rune) float64 {
	//if adv, ok := f.Face.GlyphAdvance(r); ok {
	//	return float64(adv) / 64
	//}
	//adv, _ := f.Face.GlyphAdvance(' ')
	//return float64(adv) / 64

	regularMul := 1.3
	if f.Mono {
		regularMul = .75
	}
	hm := f.Font.HMetric(1000, f.Font.Index(r))
	return float64(hm.AdvanceWidth) / float64(f.Font.Bounds(1000).Max.X) * regularMul
}

var SerifRegular = NewFont("fonts/Tinos-Regular.ttf", 0b00000000000000000000000000000110)
var SerifBold = NewFont("fonts/Tinos-Bold.ttf", 0b00000000000001000000000000000110)
var SerifItalic = NewFont("fonts/Tinos-Italic.ttf", 0b00000000000000000000000001000110)

//var TinosBoldItalic = NewFont("fonts/Tinos-BoldItalic.ttf")

var SansRegular = NewFont("fonts/Lato-Regular.ttf", 0b00000000000000000000000000000100)
var SansBold = NewFont("fonts/Lato-Bold.ttf", 0b00000000000001000000000000000100)

//var LatoItalic = NewFont("fonts/Lato-Italic.ttf")
//var LatoBoldItalic = NewFont("fonts/Lato-BoldItalic.ttf")

var Monospace = NewFont("fonts/CutiveMono-Regular.ttf", 0b00000000000000000000000000000101)
