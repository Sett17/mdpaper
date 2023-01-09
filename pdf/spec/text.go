package spec

import (
	"bytes"
	"fmt"
	"github.com/sett17/mdpaper/cli"
	"github.com/sett17/mdpaper/globals"
	"math"
	"sort"
	"strings"
)

type Segment struct {
	Content string
	Font    *Font
}

type TextLine struct {
	Words       []string
	Fonts       []*Font
	WordSpacing float64
	Width       float64
	Offset      float64
}

func (l *TextLine) Add(str string, font *Font) {
	//s := strings.NewReplacer("(", "\\(", ")", "\\)", "\\", "\\\\").Replace(str)
	s := strings.ReplaceAll(str, "\\", "\\\\")
	s = strings.ReplaceAll(s, "(", "\\(")
	s = strings.ReplaceAll(s, ")", "\\)")

	l.Words = append(l.Words, s)
	l.Fonts = append(l.Fonts, font)
}

func (l *TextLine) CalculateSpacing(maxWidth float64) {
	buf := strings.Builder{}
	for _, w := range l.Words {
		buf.WriteString(w)
	}
	spaces := float64(strings.Count(buf.String(), " "))
	if spaces == 0 {
		l.WordSpacing = 1.0
		return
	}
	diff := maxWidth - l.Width + globals.MmToPt(.7)
	l.WordSpacing = diff / spaces
}

func (l *TextLine) Center(fullWidth float64) {
	l.WordSpacing = 0
	diff := fullWidth - l.Width
	l.Offset = diff / 2
}

func (l *TextLine) String() string {
	return strings.Join(l.Words, " ")
}

type Text struct {
	Segments   []*Segment
	Pos        [2]float64
	FontSize   int
	LineHeight float64
	Processed  []*TextLine
	Offset     float64
	Width      float64
}

func (p *Text) String() string {
	str := strings.Builder{}
	for _, segment := range p.Segments {
		str.WriteString(segment.Content)
	}
	return str.String()
}

type segmentFit struct {
	segmentIdx int
	wordCount  int
}

func findCutoffSegment(segments []*Segment, cutoffText string) (int, int) {
	if segments == nil {
		return -1, -1
	}
	fit := make([]segmentFit, len(segments))
	cutoffSplitDirty := strings.Split(cutoffText, " ")
	cutoffSplit := make([]string, 0)
	for _, s := range cutoffSplitDirty {
		if s != "" {
			cutoffSplit = append(cutoffSplit, s)
		}
	}

	for i, segment := range segments {
		for j := len(cutoffSplit); j > 0; j-- {
			search := strings.Join(cutoffSplit[:j], " ")
			if strings.Contains(segment.Content, search) {
				fit[i] = segmentFit{i, j}
				break
			}
		}
	}

	sort.Slice(fit, func(i, j int) bool {
		return fit[i].wordCount > fit[j].wordCount
	})

	for _, f := range fit {
		if checkCorrectCutoff(segments, f.segmentIdx, f, cutoffSplit) {
			cutoffFromThisSegment := strings.Join(cutoffSplit[:f.wordCount], " ")
			cutoffLocation := strings.Index(segments[f.segmentIdx].Content, cutoffFromThisSegment)
			return f.segmentIdx, cutoffLocation
		}
	}

	return -1, -1
}

func checkCorrectCutoff(segments []*Segment, segmentIdx int, fit segmentFit, cutoffSplit []string) bool {
	wordCount := fit.wordCount
	cutoffPartInSegment := strings.Join(cutoffSplit[:wordCount], " ")
	if cutoffPartInSegment == strings.Join(cutoffSplit, " ") {
		return true
	}

	remainingCutoff := cutoffSplit[wordCount:]
	nextSegContent := strings.TrimSpace(segments[segmentIdx+1].Content)
	nextSegContentWordCount := len(strings.Split(nextSegContent, " "))

	nextCutoff := remainingCutoff
	if nextSegContentWordCount < len(remainingCutoff) {
		nextCutoff = remainingCutoff[:nextSegContentWordCount]
	}

	return strings.HasPrefix(nextSegContent, strings.Join(nextCutoff, " "))
}

func (p *Text) SplitDelegate(percent float64) (Addable, Addable) {
	procCutoff := int(math.Floor(float64(len(p.Processed)) * percent))
	if procCutoff == 0 {
		return nil, p
	}
	cutoffText := p.Processed[procCutoff].String()
	var leftoverSegs []*Segment

	cutOffSeg, cutoffLocation := findCutoffSegment(p.Segments, cutoffText)
	if cutOffSeg == -1 {
		cli.Error(fmt.Errorf("could not find cutoff segment"), true)
	}

	splitSegment := p.Segments[cutOffSeg]
	splitSegment.Content = splitSegment.Content[cutoffLocation:]

	leftoverSegs = append(leftoverSegs, splitSegment)
	leftoverSegs = append(leftoverSegs, p.Segments[cutOffSeg+1:]...)

	//for i, segment := range p.Segments {
	//	if strings.Contains(segment.Content, cutoffText) {
	//		// Split the segment
	//		split := strings.Split(segment.Content, cutoffText)
	//		split[1] = cutoffText + split[1]
	//		p.Segments[i].Content = split[0]
	//		segsAfter = p.Segments[i+1:]
	//		p.Segments = p.Segments[:i+1]
	//		leftoverSegs = append(leftoverSegs, &Segment{Content: split[1], Font: segment.Font})
	//		leftoverSegs = append(leftoverSegs, segsAfter...)
	//		break
	//	}
	//}

	a1 := &Text{
		Segments:   p.Segments,
		Pos:        p.Pos,
		FontSize:   p.FontSize,
		LineHeight: p.LineHeight,
		Processed:  p.Processed[:procCutoff],
		Offset:     p.Offset,
	}
	a2 := &Text{
		Segments:   leftoverSegs,
		FontSize:   p.FontSize,
		LineHeight: p.LineHeight,
		Processed:  make([]*TextLine, 0),
		Offset:     p.Offset,
	}
	return a1, a2
}

func (p *Text) SetPos(x, y float64) {
	p.Pos = [2]float64{x, y}
}

func (p *Text) Height() float64 {
	return (float64(len(p.Processed)))*p.LineHeight*float64(p.FontSize) + globals.MmToPt(globals.Cfg.Spaces.Paragraph)
}

func (p *Text) Process(maxWidth float64) {
	p.Processed = make([]*TextLine, 0)

	maxWidth -= p.Offset / 2
	p.Width = maxWidth
	l := &TextLine{WordSpacing: 1.0}
	for i := 0; i < len(p.Segments); {
		s := p.Segments[i]
		if len(s.Content) == 0 {
			i++
			continue
		}
		//split := strings.Split(s.Content, " ")
		splitSmall := strings.Split(s.Content, " ")
		split := make([]string, 0)
		for _, s := range splitSmall {
			//if len(s) > 0 {
			split = append(split, strings.SplitAfter(s, "/")...)
			//}
		}
		for j := 0; j < len(split); {
			w := split[j]
			if j != 0 && strings.TrimSpace(w) == "" {
				j++
				continue
			}
			if s.Font.WordWidth(w, p.FontSize)+l.Width <= maxWidth {
				suffix := ""
				if j != len(split)-1 {
					suffix = " "
				}
				l.Width += s.Font.WordWidth(w+suffix, p.FontSize)
				l.Add(w+suffix, s.Font)
				j++
			} else {
				if l.Width != 0 {
					l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
				}
				l.CalculateSpacing(maxWidth)
				p.Processed = append(p.Processed, l)
				l = &TextLine{WordSpacing: 1.0}
			}
		}
		i++
	}
	if len(l.Words) == 0 {
		return
	}
	l.Words[len(l.Words)-1] = strings.TrimRight(l.Words[len(l.Words)-1], " ")
	//l.Center(maxWidth)
	//l.CalculateSpacing(maxWidth)
	p.Processed = append(p.Processed, l)
	//}
}

func (p *Text) Add(a *Segment) {
	p.Segments = append(p.Segments, a)
}

func (p *Text) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("BT\n")

	buf.WriteString(fmt.Sprintf("%f %f TD\n", p.Pos[0]+p.Offset, p.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n", p.LineHeight*float64(p.FontSize)))

	// we can assume that paragraph has been processed

	buf.WriteString("T*\n")

	//TODO refactor all the text byte stuff to some common place
	var currFont *Font = nil
	for i, l := range p.Processed {
		lineBuffer := strings.Builder{}
		if l.Offset != 0 {
			buf.WriteString(fmt.Sprintf("%f 0 Td\n", l.Offset))
		}
		buf.WriteString(fmt.Sprintf("%f Tw\n", l.WordSpacing))
		for j := 0; j < len(l.Words); j++ {
			if l.Fonts[j] != currFont {
				if lineBuffer.Len() > 0 {
					buf.WriteString(fmt.Sprintf("("))
					buf.Write(globals.PDFEncode(lineBuffer.String()))
					buf.WriteString(fmt.Sprintf(") Tj\n"))
					lineBuffer.Reset()
				}
				buf.WriteString(fmt.Sprintf("/%s %d Tf\n", l.Fonts[j].Name, p.FontSize))
				//buf.WriteString(fmt.Sprintf("/%s %d Tf\n", "Times-Roman", p.FontSize))
				currFont = l.Fonts[j]
			}
			lineBuffer.WriteString(l.Words[j])
		}
		if lineBuffer.Len() > 0 {
			//buf.WriteString(fmt.Sprintf("(%s) Tj\n", u))
			buf.WriteString(fmt.Sprintf("("))
			buf.Write(globals.PDFEncode(lineBuffer.String()))
			buf.WriteString(fmt.Sprintf(") Tj\n"))
		}
		if i != len(p.Processed)-1 {
			buf.WriteString("T* ")
		}
		if l.Offset != 0 {
			buf.WriteString(fmt.Sprintf("%f 0 Td\n", -l.Offset))
		}
	}

	buf.WriteString("ET\n")

	//rect := GraphicRect{
	//	Pos:   [2]float64{p.Pos[0] + p.Offset, p.Pos[1]},
	//	W:     p.Width,
	//	H:     p.Height(),
	//	Color: [3]float64{0.5, 0.5, 0.0},
	//}
	//buf.WriteString("\n")
	//buf.Write(rect.Bytes())

	return buf.Bytes()
}
