package globals

import "embed"

func InToPt(in float64) float64 {
	return in * 72.0
}

func MmToPt(Mm float64) float64 {
	return InToPt(Mm / 25.4)
}

//go:embed fonts/*
var Fonts embed.FS
