package spec

import (
	"github.com/beta/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"mdpaper/globals"
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
	//TODO gen font object
	fontObj := NewDictObject()
	fontObj.Set("Type", "/Font")
	fontObj.Set("Subtype", "/TrueType")
	fontObj.Set("BaseFont", "/"+f.Font.Name(truetype.NameIDPostscriptName))

	fontObj.Set("FirstChar", 32)
	fontObj.Set("LastChar", 255)
	w := NewArrayObject()
	boldMul := 1.0
	if f.Bold {
		boldMul = 1.2
	}
	for i := 32; i < 256; i++ {
		w.Items = append(w.Items, int(f.CharWidth(rune(i))*(1000/boldMul)))
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
		panic(err)
	}
	fontFile, err := truetype.Parse(fontFileBuf)
	if err != nil {
		panic(err)
	}
	f.Name = strings.ToLower(strings.ReplaceAll(fontFile.Name(truetype.NameIDPostscriptName), " ", "_"))
	f.Font = fontFile
	f.Face = truetype.NewFace(fontFile, &truetype.Options{
		Size: float64(globals.Cfg.FontSize),
	})

	ttfStream := NewStreamObject()
	ttfStream.Set("Length1", len(fontFileBuf))
	ttfStream.AddBytes(fontFileBuf)
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
	d.Set("Name", "/"+f.Font.Name(truetype.NameIDPostscriptName))
	d.Set("Flags", flags)
	bbox := f.Font.Bounds(fixed.Int26_6(f.Font.FUnitsPerEm()))
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

	boldMul := 1.0
	if f.Bold {
		boldMul = 1.2
	}
	regularMul := 1.3
	if f.Mono {
		regularMul = .75
	}
	hm := f.Font.HMetric(fixed.Int26_6(f.Font.FUnitsPerEm()), f.Font.Index(r))
	return float64(hm.AdvanceWidth) / float64(f.Font.Bounds(fixed.Int26_6(f.Font.FUnitsPerEm())).Max.X) * regularMul * boldMul
}

var TinosRegular = NewFont("fonts/Tinos-Regular.ttf", 0b00000000000000000000000000000010)
var TinosBold = NewFont("fonts/Tinos-Bold.ttf", 0b00000000000001000000000000000010)
var TinosItalic = NewFont("fonts/Tinos-Italic.ttf", 0b00000000000000000000000010000010)

//var TinosBoldItalic = NewFont("fonts/Tinos-BoldItalic.ttf")

var LatoRegular = NewFont("fonts/Lato-Regular.ttf", 0b00000000000000000000000000000010)
var LatoBold = NewFont("fonts/Lato-Bold.ttf", 0b00000000000001000000000000000000)

//var LatoItalic = NewFont("fonts/Lato-Italic.ttf")
//var LatoBoldItalic = NewFont("fonts/Lato-BoldItalic.ttf")

var SourceCodeProRegular = NewFont("fonts/SourceCodePro-Regular.ttf", 0b00000000000000000000000000000001)
