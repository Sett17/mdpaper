package spec

import "reflect"

func InToPt(in float64) float64 {
	return in * 72.0
}

func MmToPt(Mm float64) float64 {
	return InToPt(Mm / 25.4)
}

func Cast(addable *Addable, p reflect.Type) reflect.Value {
	return reflect.ValueOf(addable).Convert(p)
}
