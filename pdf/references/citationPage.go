package references

import (
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spec"
)

var CitationHeading *elements.Heading

func GenerateCitationHeading() {
	h := elements.Heading{
		Text: spec.Text{
			FontSize:   20.0,
			LineHeight: 1.0,
			Offset:     0.0,
		}, Level: 1}
	seg := spec.Segment{
		Content: globals.Cfg.Citation.Heading,
		Font:    spec.SansBold,
	}
	h.Add(&seg)
	CitationHeading = &h
}

func Citations() (ret []*spec.Addable) {
	bib, err := globals.Citeproc.MakeBibliography()
	if err != nil {
		cli.Error(err, false)
		return
	}
	for _, entry := range bib {
		block := spec.Text{
			FontSize:   12,
			LineHeight: 1.4,
		}
		seg := spec.Segment{
			Content: entry,
			Font:    spec.SerifRegular,
		}
		block.Add(&seg)
		var a spec.Addable = &block
		ret = append(ret, &a)
	}
	return
}
