package spec

type Addable interface {
	Bytes() []byte
	SetPos(x, y float64)
	Height() float64
	Process(width float64)
}

type Splittable interface {
	Split(percent float64) (Addable, Addable)
}
