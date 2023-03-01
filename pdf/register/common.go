package register

import (
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/elements"
)

type common struct {
	heading *elements.Heading
	paper   *abstracts.Paper
	Pages   []*abstracts.Page
}

func (c *common) GenerateEntries() {
	c.paper = &abstracts.Paper{}
}

func (c *common) GeneratePages() {
	pages := make([]*abstracts.Page, 0)
	for !c.paper.Finished() {
		page := abstracts.NewPage(c.paper, 1)
		pages = append(pages, page)
	}
	c.Pages = pages
}

func (c *common) Heading() *elements.Heading {
	return c.heading
}
