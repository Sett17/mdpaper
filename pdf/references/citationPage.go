package references

import (
	"fmt"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spec"
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
	bibs := make([]string, len(globals.BibIndices)+1)
	for key, idx := range globals.BibIndices {
		bibs[idx] = fmt.Sprintf("[%d] %s", idx, globals.IEEE(globals.Bibs[key]))
	}
	for _, bib := range bibs {
		block := spec.Text{
			FontSize:   14,
			LineHeight: 1.4,
		}
		seg := spec.Segment{
			Content: bib,
			Font:    spec.SerifRegular,
		}
		block.Add(&seg)
		var a spec.Addable = &block
		ret = append(ret, &a)
	}
	return
}
