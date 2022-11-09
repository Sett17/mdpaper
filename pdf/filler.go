package pdf

type Filler struct {
}

func (f Filler) Bytes() []byte {
	return []byte{}
}

func (f Filler) SetPos(_, _ float64) {
}

func (f Filler) Height() float64 {
	return 0
}

func (f Filler) Process(_ float64) {
}
