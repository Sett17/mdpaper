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
	Pos       [2]float64
	Tokens    chroma.Iterator
	Processed []*codeLine
	w         float64
	Style     *chroma.Style
}

func (f *FencedCode) Bytes() []byte {
	buf := bytes.Buffer{}

	col := f.Style.Get(chroma.Text).Background
	r := spec.GraphicRect{
		Pos:     [2]float64{f.Pos[0], f.Pos[1] - globals.MmToPt(2)},
		W:       f.w,
		H:       f.Height() - globals.Cfg.Spaces.Code,
		Color:   [3]float64{math.Min(float64(col.Red())/255, 1), math.Min(float64(col.Green())/255, 1), math.Min(float64(col.Blue())/255, 1)},
		Filled:  true,
		Rounded: true,
	}
	buf.Write(r.Bytes())

	buf.WriteString("BT\n")
	buf.WriteString("q\n")
	buf.WriteString(fmt.Sprintf("%f %f Td\n", f.Pos[0]+globals.MmToPt(1), f.Pos[1]-float64(globals.Cfg.Code.FontSize)-globals.MmToPt(2+1)))
	buf.WriteString(fmt.Sprintf("%f TL\n", float64(globals.Cfg.Code.FontSize)))
	buf.WriteString(fmt.Sprintf("%f Tc\n", globals.Cfg.Code.CharacterSpacing))
	buf.WriteString(fmt.Sprintf("/%s %d Tf\n", spec.Monospace.Name, globals.Cfg.Code.FontSize))

	lineNumberFmt := ""
	if globals.Cfg.Code.LineNumbers {
		lineNumberFmt = fmt.Sprintf("(%%%dd ) Tj\n", len(fmt.Sprintf("%d", len(f.Processed))))
	}
	for k, line := range f.Processed {
		buf.WriteString(".5 .5 .5 rg\n")
		buf.WriteString(fmt.Sprintf(lineNumberFmt, k+1))
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
	return float64(len(f.Processed))*float64(globals.Cfg.Code.FontSize) + globals.MmToPt(4) + globals.Cfg.Spaces.Code
}

type token struct {
	Value string
	Color [3]float64
}

func (f *FencedCode) Process(width float64) {
	f.w = width
	tokens := make([]token, 0)
	for t := f.Tokens(); t != chroma.EOF; t = f.Tokens() {
		if strings.HasPrefix(t.Value, "\n") {
			tokens = append(tokens, token{Value: "\n", Color: [3]float64{0, 0, 0}})
			if strings.TrimPrefix(t.Value, "\n") != "" {
				tokens = append(tokens, token{Value: strings.TrimPrefix(t.Value, "\n"), Color: [3]float64{0, 0, 0}})
			}
			continue
		}
		if strings.HasSuffix(t.Value, "\n") {
			if strings.TrimSuffix(t.Value, "\n") != "" {
				tokens = append(tokens, token{Value: strings.TrimSuffix(t.Value, "\n"), Color: [3]float64{0, 0, 0}})
			}
			tokens = append(tokens, token{Value: "\n", Color: [3]float64{0, 0, 0}})
			continue
		}
		if t.Value == "" {
			continue
		}
		col := f.Style.Get(t.Type).Colour + 1
		tokens = append(tokens, token{Value: t.Value, Color: [3]float64{math.Min(float64(col.Red())/255, 1), math.Min(float64(col.Green())/255, 1), math.Min(float64(col.Blue())/255, 1)}})
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
