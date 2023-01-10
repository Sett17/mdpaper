package elements

import (
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/spec"
)

type ListItem struct {
	spec.Text
	Prefix string
}

func (p *ListItem) Bytes() []byte {
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
