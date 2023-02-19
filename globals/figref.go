package globals

type FigureInformation struct {
	Number int
	Title  string
	Key    string
	Used   []int
}

var figureIndex = 0

func NewFigureInformation(title string, key string) *FigureInformation {
	figureIndex++
	return &FigureInformation{
		Number: figureIndex,
		Title:  title,
		Key:    key,
		Used:   []int{},
	}
}

var Figures = make(map[string]*FigureInformation)
