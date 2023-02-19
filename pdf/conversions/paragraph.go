package conversions

import (
	"github.com/sett17/mdpaper/v2/pdf/elements"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"github.com/yuin/goldmark/ast"
)

func Paragraph(p *ast.Paragraph, centered bool) *spec.Addable {
	textBlock := ast.TextBlock{BaseBlock: p.BaseBlock}
	textBlockAddable := TextBlock(&textBlock)
	textBlockText := (*textBlockAddable).(*spec.Text)

	para := elements.Paragraph{
		Text:     *textBlockText,
		Centered: centered,
	}

	var a spec.Addable = &para
	return &a
}
