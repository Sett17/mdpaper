package elements

import (
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/spec"
)

type ListItem struct {
	spec.Text
	Prefix string
}

func (p *ListItem) Bytes() []byte {
	if len(p.Processed) == 0 {
		p.Processed = make([]*spec.TextLine, 1)
		p.Processed[0] = &spec.TextLine{}
	}
	firstLine := p.Processed[0]
	words := firstLine.Words
	fonts := firstLine.Fonts
	newWords := []string{p.Prefix}
	newFonts := []*spec.Font{spec.SerifRegular}
	newWords = append(newWords, words...)
	newFonts = append(newFonts, fonts...)
	newLine := spec.TextLine{
		Words:       newWords,
		Fonts:       newFonts,
		WordSpacing: firstLine.WordSpacing,
		Width:       firstLine.Width + spec.SerifRegular.WordWidth(p.Prefix, globals.Cfg.Text.FontSize),
		Offset:      0,
	}
	p.Processed[0] = &newLine
	return p.Text.Bytes()
}
