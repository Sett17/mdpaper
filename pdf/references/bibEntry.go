package references

import (
	"github.com/sett17/mdpaper/v2/pdf/spec"
)

type BibEntry struct {
	spec.Text
}

func (b *BibEntry) Process(maxWidth float64) {
	b.Text.Process(maxWidth)
	for i, _ := range b.Processed {
		b.Processed[i].WordSpacing = 0.0
	}
}
