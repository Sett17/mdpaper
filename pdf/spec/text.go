package spec

import (
	"bytes"
	"fmt"
	"github.com/sett17/mdpaper/v2/cli"
	"math"
	"sort"
	"strings"
)

type Segment struct {
	Content string
	Font    *Font
}

type Text struct {
	Segments   []*Segment
	Pos        [2]float64
	FontSize   int
	LineHeight float64
	Processed  JustifiedText
	Offset     float64
	Width      float64
	Margin     float64
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
	cutoffText = deEscape(cutoffText)
	var leftoverSegs []*Segment

	cutOffSeg, cutoffLocation := findCutoffSegment(p.Segments, cutoffText)
	if cutOffSeg == -1 {
		cli.Error(fmt.Errorf("could not find cutoff segment for '%s'", cutoffText), true)
	}

	splitSegment := p.Segments[cutOffSeg]
	splitSegment.Content = splitSegment.Content[cutoffLocation:]

	leftoverSegs = append(leftoverSegs, splitSegment)
	leftoverSegs = append(leftoverSegs, p.Segments[cutOffSeg+1:]...)

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
	return (float64(len(p.Processed)))*p.LineHeight*float64(p.FontSize) + p.Margin
}

func (p *Text) Process(maxWidth float64) {
	p.Width = maxWidth - p.Offset/2
	p.Processed = ProcessSegments(p.Segments, maxWidth, p.FontSize, p.Offset)
}

func (p *Text) Add(a ...*Segment) {
	p.Segments = append(p.Segments, a...)
}

func (p *Text) Bytes() []byte {
	buf := bytes.Buffer{}

	buf.WriteString("BT\n")
	buf.WriteString(fmt.Sprintf("%f %f TD\n", p.Pos[0], p.Pos[1]))
	buf.WriteString(fmt.Sprintf("%f TL\n", float64(p.FontSize)*p.LineHeight))

	// we can assume that paragraph has been processed

	buf.WriteString("T*\n")

	buf.Write(p.Processed.Bytes(p.FontSize))

	buf.WriteString("ET\n")

	return buf.Bytes()
}

func SplitLongString(s string, splitWidth float64, fontSize int, font *Font) (pre string, rem string) {
	// if remaining space is smaller than 3 charcters just putit into the next line...
	//if splitWidth <= font.WordWidth("xxx", fontSize) {
	//	return "", s
	//}

	//first try to split at slashes for links ot filepaths
	pre = s
	for i := strings.LastIndexAny(pre, "/\\"); i != -1; i = strings.LastIndexAny(pre, "/\\") {
		pre = pre[:i]
		preWidth := font.WordWidth(pre, fontSize)
		if preWidth < splitWidth {
			if preWidth > splitWidth*.8 {
				return pre, s[len(pre):]
			} else {
				break
			}
		}
	}

	//if regular string use kinda of a binary search
	pre = s

	extra := ""
	if Entropy(s) < 4 { // not sure about the number 4, but if under 4 likely to be some regular text and not part of a url or something
		extra = "-"
	}

	for {
		if font.WordWidth(pre+extra, fontSize) < splitWidth {
			return pre + extra, s[len(pre):]
		}
		pre = pre[:len(pre)-1]
	}
}

func Entropy(s string) float64 {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	var entropy float64
	for _, c := range counts {
		p := float64(c) / float64(len(s))
		entropy -= p * math.Log2(p)
	}
	return entropy
}
