package globals

const (
	A4Width  = 595.2755905511812
	A4Height = 841.8897637795276
	//DBG      = true
	//DBG = false
)

var ColumnHeight = A4Height

func InToPt(in float64) float64 {
	return in * 72.0
}

func MmToPt(Mm float64) float64 {
	return InToPt(Mm / 25.4)
}
