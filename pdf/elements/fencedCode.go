package elements

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/spec"
	"math"
	"strings"
)

type codeLine struct {
	Words  []string
	Colors [][3]float64
}

func (l *codeLine) Add(str string, color [3]float64) {
	s := strings.ReplaceAll(str, "\\", "\\\\")
	s = strings.ReplaceAll(s, "(", "\\(")
	s = strings.ReplaceAll(s, ")", "\\)")
	l.Words = append(l.Words, s)
	l.Colors = append(l.Colors, color)
}

type FencedCode struct {
	Pos         [2]float64
	Tokens      chroma.Iterator
	Processed   []*codeLine
	w           float64
	Style       *chroma.Style
	LineNumbers bool
	StartNumber int
	FontSize    int
	FlatTop     bool
	FlatBottom  bool
}

func (f *FencedCode) Split(percent float64) (spec.Addable, spec.Addable) {
	cutoffLine := int(math.Round(float64(len(f.Processed)) * percent))

	cutoffBlockLines := f.Processed[cutoffLine:]
	cutoffBlock := FencedCode{
		Processed:   cutoffBlockLines,
		w:           f.w,
		Style:       f.Style,
		LineNumbers: f.LineNumbers,
		StartNumber: cutoffLine + f.StartNumber,
		FontSize:    f.FontSize,
		FlatTop:     true,
		FlatBottom:  f.FlatBottom,
	}

	f.Processed = f.Processed[:cutoffLine]
	f.FlatBottom = true

	return f, &cutoffBlock
}

func (f *FencedCode) Bytes() []byte {
	buf := bytes.Buffer{}

	backcol := f.Style.Get(chroma.Text).Background
	rectColor := [3]float64{math.Min(float64(backcol.Red())/255, 1), math.Min(float64(backcol.Green())/255, 1), math.Min(float64(backcol.Blue())/255, 1)}
	borderColor := rectColor
	if backcol.Brightness() > .9 {
		borderColor = [3]float64{math.Abs(rectColor[0] - 1), math.Abs(rectColor[1] - 1), math.Abs(rectColor[2] - 1)}
	}
	r := spec.GraphicRect{
		Pos:           [2]float64{f.Pos[0], f.Pos[1] - globals.MmToPt(2)},
		W:             f.w,
		H:             f.Height() - globals.Cfg.Spaces.Code,
		Color:         rectColor,
		BorderColor:   borderColor,
		Filled:        true,
		RoundedTop:    !f.FlatTop,
		RoundedBottom: !f.FlatBottom,
	}
	buf.Write(r.Bytes())

	buf.WriteString("BT\n")
	buf.WriteString("q\n")
	buf.WriteString(fmt.Sprintf("%f %f Td\n", f.Pos[0]+globals.MmToPt(1), f.Pos[1]-float64(f.FontSize)-globals.MmToPt(2+1)))
	buf.WriteString(fmt.Sprintf("%f TL\n", float64(f.FontSize)))
	buf.WriteString(fmt.Sprintf("%f Tc\n", globals.Cfg.Code.CharacterSpacing))
	buf.WriteString(fmt.Sprintf("/%s %d Tf\n", spec.Monospace.Name, f.FontSize))

	lineNumberLength := len(fmt.Sprintf("%d", len(f.Processed)))
	for k, line := range f.Processed {
		buf.WriteString(".5 .5 .5 rg\n")
		if f.LineNumbers {
			buf.WriteString(fmt.Sprintf("(%*d ) Tj\n", lineNumberLength, f.StartNumber+k+1))
		}
		for i := 0; i < len(line.Words); i++ {
			buf.WriteString(fmt.Sprintf("%f %f %f rg\n", line.Colors[i][0], line.Colors[i][1], line.Colors[i][2]))
			buf.WriteString(fmt.Sprintf("("))
			buf.Write(globals.PDFEncode(line.Words[i]))
			buf.WriteString(fmt.Sprintf(") Tj\n"))
		}
		buf.WriteString("T*\n")
	}

	buf.WriteString("Q\n")
	buf.WriteString("ET\n")

	return buf.Bytes()
}

func (f *FencedCode) SetPos(x, y float64) {
	f.Pos = [2]float64{x, y + globals.MmToPt(2)}
}

func (f *FencedCode) Height() float64 {
	return float64(len(f.Processed))*float64(f.FontSize) + globals.MmToPt(4) + globals.Cfg.Spaces.Code
}

type token struct {
	Value string
	Color [3]float64
}

func processSingleToken(t chroma.Token, style *chroma.Style) (tokens []token) {
	if strings.HasPrefix(t.Value, "\n") {
		tokens = append(tokens, token{Value: "\n", Color: [3]float64{0, 0, 0}})
		if strings.TrimPrefix(t.Value, "\n") != "" {
			col := style.Get(t.Type).Colour
			tokens = append(tokens, token{Value: strings.TrimPrefix(t.Value, "\n"), Color: [3]float64{math.Min(float64(col.Red())/255, 1), math.Min(float64(col.Green())/255, 1), math.Min(float64(col.Blue())/255, 1)}})
		}
		return
	}
	if strings.HasSuffix(t.Value, "\n") {
		if strings.TrimSuffix(t.Value, "\n") != "" {
			col := style.Get(t.Type).Colour
			tokens = append(tokens, token{Value: strings.TrimSuffix(t.Value, "\n"), Color: [3]float64{math.Min(float64(col.Red())/255, 1), math.Min(float64(col.Green())/255, 1), math.Min(float64(col.Blue())/255, 1)}})
		}
		tokens = append(tokens, token{Value: "\n", Color: [3]float64{0, 0, 0}})
		return
	}
	if t.Value == "" {
		return
	}
	col := style.Get(t.Type).Colour
	tokens = append(tokens, token{Value: t.Value, Color: [3]float64{math.Min(float64(col.Red())/255, 1), math.Min(float64(col.Green())/255, 1), math.Min(float64(col.Blue())/255, 1)}})
	return
}

func (f *FencedCode) Process(width float64) {
	//disable processing multiple times
	if f.Processed != nil {
		return
	}

	f.w = width
	tokens := make([]token, 0)
	for t := f.Tokens(); t != chroma.EOF; t = f.Tokens() {

		split := strings.SplitAfter(t.Value, "\n")
		for _, s := range split {
			tokens = append(tokens, processSingleToken(chroma.Token{Value: s, Type: t.Type}, f.Style)...)
		}

	}
	if len(tokens) > 0 && tokens[len(tokens)-1].Value == "\n" {
		tokens = tokens[:len(tokens)-1]
	}

	lineBuffer := &codeLine{}
	for _, t := range tokens {
		if t.Value == "\n" {
			f.Processed = append(f.Processed, lineBuffer)
			lineBuffer = &codeLine{}
			continue
		}
		lineBuffer.Add(t.Value, t.Color)
	}
	if len(lineBuffer.Words) > 0 {
		f.Processed = append(f.Processed, lineBuffer)
	}
}
