package spec

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Image struct {
	Width, Height int
	Dictionary
	Stream
}

func NewImage(path string) Image {
	iFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer iFile.Close()
	iData, _, err := image.Decode(iFile)
	if err != nil {
		panic(err)
	}
	i := Image{}
	i.Set("Type", "/XObject")
	i.Set("Subtype", "/Image")
	i.Set("Width", iData.Bounds().Dx())
	i.Set("Height", iData.Bounds().Dy())
	i.Width = iData.Bounds().Dx()
	i.Height = iData.Bounds().Dy()
	i.Set("ColorSpace", "/DeviceRGB")
	i.Set("BitsPerComponent", 16)
	for j := 0; j < iData.Bounds().Dy(); j++ {
		for k := 0; k < iData.Bounds().Dx(); k++ {
			r, g, b, _ := iData.At(j, k).RGBA()
			i.Write([]byte{byte(r), byte(g), byte(b)})
		}
	}
	//i.Set("Filter", "/FlateDecode")
	//i.Stream.Deflate()
	//i.Set("Length", i.Len())
	i.Set("Size", i.Len())
	return i
}

type ImageObject struct {
	GenericObject
	Name string
	Image
}

func NewImageObject() ImageObject {
	LastId++
	return ImageObject{GenericObject: GenericObject{id: LastId}}
}

func (i *ImageObject) Bytes() []byte {
	buf := bytes.Buffer{}
	beg, end := i.ByteParts()
	buf.Write(beg)
	buf.Write(i.Dictionary.Bytes())
	buf.Write(i.Stream.Bytes())
	buf.Write(end)
	return buf.Bytes()
}

func (i *ImageObject) Pointer() *Object {
	var z Object = i
	return &z
}
