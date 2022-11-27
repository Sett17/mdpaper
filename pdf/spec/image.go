package spec

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mdpaper/globals"
	"os"
)

func NewImageObject(path string, mul float64) (XObject, Addable) {
	LastId++
	iName := ""
	var iData image.Image
	iFile, err := os.Open(path)
	//if err != nil {
	//	panic(err)
	//}
	if err != nil {
		r := FillingRect{
			GraphicRect: GraphicRect{
				Pos: [2]float64{},
			},
		}
		return XObject{}, &r
	} else {
		iData, _, err = image.Decode(iFile)
		iName = iFile.Name()
		if err != nil {
			panic(err)
		}
	}
	defer iFile.Close()
	x := NewXObject(iName)
	x.Dictionary.Set("Type", "/XObject")
	x.Dictionary.Set("Subtype", "/Image")
	pixelMul := 1
	if float64(iData.Bounds().Dx()) > 800 {
		//pixelMul = 2
	}
	x.Dictionary.Set("Width", iData.Bounds().Dx()/pixelMul)
	x.Dictionary.Set("Height", iData.Bounds().Dy()/pixelMul)
	//x.Dictionary.Width = iData.Bounds().Dx()
	//x.Dictionary.Height = iData.Bounds().Dy()
	x.Dictionary.Set("ColorSpace", "/DeviceRGB")
	x.Dictionary.Set("BitsPerComponent", 8)
	for j := 0; j < iData.Bounds().Dy()/pixelMul; j++ {
		for k := 0; k < iData.Bounds().Dx()/pixelMul; k++ {
			r, g, b, _ := iData.At(k*pixelMul, j*pixelMul).RGBA()
			x.Write([]byte{byte(r), byte(g), byte(b)})
		}
	}
	x.WriteString("\n")
	x.Set("Size", x.Len())
	ia := ImageAddable{
		ImageName: x.Name,
		W:         float64(iData.Bounds().Dx()),
		H:         float64(iData.Bounds().Dy()),
		Pos:       [2]float64{},
		Mul:       mul,
	}
	return x, &ia
}

type ImageAddable struct {
	ImageName string
	W         float64
	H         float64
	Pos       [2]float64
	Mul       float64
	Offset    float64
}

func (i *ImageAddable) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("q\n")
	//buf.WriteString(fmt.Sprintf("%f 0 0 %f %f %f cm\n", .25, .25, i.Pos[0], i.Pos[1]-i.H))
	buf.WriteString(fmt.Sprintf("%f 0 0 %f %f %f cm\n", i.W, i.H, i.Pos[0]+i.Offset, i.Pos[1]-i.H))
	buf.WriteString(fmt.Sprintf("/%s Do\n", i.ImageName))
	buf.WriteString("Q\n")
	return buf.Bytes()
}

func (i *ImageAddable) SetPos(x, y float64) {
	i.Pos = [2]float64{x, y - globals.MmToPt(3)}
}

func (i *ImageAddable) Height() float64 {
	return i.H + globals.MmToPt(3.5)
}

func (i *ImageAddable) Process(width float64) {
	adjWidth := width * i.Mul
	i.Offset = (width - adjWidth) / 2
	ratio := i.W / i.H
	i.W = adjWidth
	i.H = adjWidth / ratio
}
