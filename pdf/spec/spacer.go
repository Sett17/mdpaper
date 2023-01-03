package spec

type Spacer struct {
	Pos [2]float64
	H   float64
}

func NewSpacer(h float64) *Spacer {
	return &Spacer{
		H: h,
	}
}

func (s *Spacer) Bytes() []byte {
	return []byte{}
}

func (s *Spacer) SetPos(x, y float64) {
	s.Pos = [2]float64{x, y}
}

func (s *Spacer) Height() float64 {
	return s.H
}

func (s *Spacer) Process(width float64) {
	// nothing to do
}
