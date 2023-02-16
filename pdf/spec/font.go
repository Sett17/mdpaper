package spec

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/sett17/goafm"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"strings"
)

type Font struct {
	FilePath   string
	Font       *truetype.Font
	Metrics    *goafm.FontMetric
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
	fontObj.Set("Encoding", "/WinAnsiEncoding")
	fontObj.Set("BaseFont", "/"+f.Metrics.FontName)

	fontObj.Set("FirstChar", 32)
	fontObj.Set("LastChar", 255)

	w := NewArrayObject()
	for i := 32; i < 256; i++ {
		r := rune(i)
		rb := globals.WinAnsiEncode(string(r))[0]
		w.Items = append(w.Items, int(f.CharWidth(rune(rb))*(1000)))
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
	fontFileBuf, err := globals.Fonts.ReadFile("fonts/" + filePath + ".ttf")
	if err != nil {
		cli.Error(fmt.Errorf("could not read font file %s", filePath), true)
	}
	fontFile, err := truetype.Parse(fontFileBuf)
	if err != nil {
		cli.Error(fmt.Errorf("could not parse font file %s", filePath), true)
	}
	afmFileBuf, err := globals.Fonts.ReadFile("fonts/" + filePath + ".afm")
	if err != nil {
		cli.Error(fmt.Errorf("could not read afm file %s", filePath), true)
	}
	f.Metrics, err = goafm.Parse(afmFileBuf)
	if err != nil {
		cli.Error(fmt.Errorf("could not parse afm file %s", filePath), true)
	}

	f.Name = strings.ToLower(strings.ReplaceAll(fontFile.Name(truetype.NameIDPostscriptName), " ", "_"))
	f.Font = fontFile

	ttfStream := NewStreamObject()
	ttfStream.AlwaysDeflate = true
	ttfStream.Set("Length1", len(fontFileBuf))
	ttfStream.Write(fontFileBuf)
	f.Data = &ttfStream

	desc := generateDescriptor(f, flags)
	desc.Set("FontFile2", ttfStream.Reference())
	f.Descriptor = desc

	f.Bold = flags&ForceBold != 0
	f.Mono = flags&FixedPitch != 0

	return
}

func generateDescriptor(f *Font, flags int) Dictionary {
	d := NewDict()
	d.Set("Type", "/FontDescriptor")
	d.Set("FontName", "/"+f.Metrics.FontName)
	d.Set("Name", "/"+f.Name)
	d.Set("Flags", flags)
	d.Set("FontBBox", Array{Items: []interface{}{f.Metrics.FontBBox[0].GetInt(), f.Metrics.FontBBox[1].GetInt(), f.Metrics.FontBBox[2].GetInt(), f.Metrics.FontBBox[3].GetInt()}})
	d.Set("Ascent", int(f.Metrics.Ascender))
	d.Set("Descent", int(f.Metrics.Descender))
	d.Set("ItalicAngle", f.Metrics.ItalicAngle)
	//d.Set("StemV", f.Metrics.StdVW)
	d.Set("StemV", 80)
	d.Set("XHeight", f.Metrics.XHeight)
	d.Set("CapHeight", f.Metrics.CapHeight)
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
	m := f.Metrics.MetricByRune(r)
	if m == nil {
		return f.Metrics.FontBBox[3].GetFloat() / 1e3
	}
	return f.Metrics.MetricByRune(r).WX.GetFloat() / 1e3
}

const (
	FixedPitch  = 0b1
	Serif       = 0b10
	Symbolic    = 0b100
	Script      = 0b1000
	NonSymbolic = 0b100000
	Italic      = 0b1000000
	AllCap      = 0b10000000000000000
	SmallCap    = 0b100000000000000000
	ForceBold   = 0b1000000000000000000
)

var SerifRegular = NewFont("SourceSerifPro-Regular", Serif|Symbolic)
var SerifBold = NewFont("SourceSerifPro-Bold", Serif|Symbolic|ForceBold)
var SerifItalic = NewFont("SourceSerifPro-Italic", Serif|Symbolic|Italic)

var SansRegular = NewFont("SourceSans3-Regular", NonSymbolic)
var SansBold = NewFont("SourceSans3-Bold", Symbolic|ForceBold)

var Monospace = NewFont("SourceCodePro-Regular", FixedPitch|Symbolic)
