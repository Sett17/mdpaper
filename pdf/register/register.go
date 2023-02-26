package register

import (
	"github.com/sett17/mdpaper/v2/pdf/abstracts"
	"github.com/sett17/mdpaper/v2/pdf/elements"
)

type Register interface {
	GenerateEntries()
	GeneratePages() []*abstracts.Page
	Heading() *elements.Heading
}
