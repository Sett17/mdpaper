package spec

// Font non-complete/opinionated Implementation
type Font struct {
	Widths     Array
	FontName   string
	Name       string
	FirstChar  int
	LastChar   int
	Descriptor Dictionary
}

func (f *Font) AddToPDF(pdf *PDF) (ref string, name string) {
	widthsObj := NewArrayObject()
	widthsObj.Array = f.Widths
	pdf.AddObject(widthsObj.Pointer())
	widthsRef := widthsObj.Reference()

	descriptorObj := NewDictObject()
	descriptorObj.Dictionary = f.Descriptor
	pdf.AddObject(descriptorObj.Pointer())
	descriptorRef := descriptorObj.Reference()

	fontObj := NewDictObject()
	fontObj.Set("Type", "/Font")
	fontObj.Set("Subtype", "/Type1")
	fontObj.Set("Name", f.Name)
	fontObj.Set("BaseFont", f.FontName)
	fontObj.Set("FirstChar", f.FirstChar)
	fontObj.Set("LastChar", f.LastChar)
	fontObj.Set("Widths", widthsRef)
	fontObj.Set("FontDescriptor", descriptorRef)
	pdf.AddObject(fontObj.Pointer())
	return fontObj.Reference(), f.Name
}

func (f *Font) WordWidth(w string, fs int) float64 {
	var width float64
	for _, r := range w {
		width += f.CharWidth(r)
	}
	return width * float64(fs)
}

func (f *Font) CharWidth(r rune) float64 {
	return float64(f.Widths.Items[int(r)-f.FirstChar].(int)) / 1000
}

var TimesRegular = Font{
	FontName:  "/Times-Roman",
	Name:      "times",
	FirstChar: 32,
	LastChar:  251,
	Widths: Array{
		Items: []interface{}{
			250, 333, 408, 500, 500, 833, 778, 333, 333, 333, 500, 564, 250, 333, 250, 278, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 278, 278, 564, 564, 564, 444, 921, 722, 667, 667, 722, 611, 556, 722, 722, 333, 389, 722, 611, 889, 722, 722, 556, 722, 667, 556, 611, 722, 722, 944, 722, 722, 611, 333, 278, 333, 469, 500, 333, 444, 500, 444, 500, 444, 333, 500, 500, 278, 278, 500, 278, 778, 500, 500, 500, 500, 333, 389, 278, 500, 500, 722, 500, 500, 444, 480, 200, 480, 541, 333, 500, 500, 167, 500, 500, 500, 500, 180, 444, 500, 333, 333, 556, 556, 500, 500, 500, 250, 453, 350, 333, 444, 444, 500, 1000, 1000, 444, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 1000, 889, 276, 611, 722, 889, 310, 667, 278, 278, 500, 722, 500,
		},
	},
	Descriptor: Dictionary{
		M: map[string]interface{}{
			"Type":     "/FontDescriptor",
			"FontName": "/Times-Roman",
			"Flags":    0b00000000000000000000000000000010,
			"FontBBox": Array{
				Items: []interface{}{
					-168, -218, 1000, 898,
				},
			},
			"ItalicAngle": 0,
			"Ascent":      683,
			"Descent":     -217,
			"CapHeight":   669,
		},
	},
}

var TimesBold = Font{
	FontName:  "/Times-Bold",
	Name:      "timesbold",
	FirstChar: 32,
	LastChar:  251,
	Widths: Array{
		Items: []interface{}{
			250, 333, 408, 500, 500, 833, 778, 333, 333, 333, 500, 564, 250, 333, 250, 278, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 278, 278, 564, 564, 564, 444, 921, 722, 667, 667, 722, 611, 556, 722, 722, 333, 389, 722, 611, 889, 722, 722, 556, 722, 667, 556, 611, 722, 722, 944, 722, 722, 611, 333, 278, 333, 469, 500, 333, 444, 500, 444, 500, 444, 333, 500, 500, 278, 278, 500, 278, 778, 500, 500, 500, 500, 333, 389, 278, 500, 500, 722, 500, 500, 444, 480, 200, 480, 541, 333, 500, 500, 167, 500, 500, 500, 500, 180, 444, 500, 333, 333, 556, 556, 500, 500, 500, 250, 453, 350, 333, 444, 444, 500, 1000, 1000, 444, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 1000, 889, 276, 611, 722, 889, 310, 667, 278, 278, 500, 722, 500,
		},
	},
	Descriptor: Dictionary{
		M: map[string]interface{}{
			"Type":     "/FontDescriptor",
			"FontName": "/Times-Bold",
			"Flags":    0b00000000000000000000000000000010,
			"FontBBox": Array{
				Items: []interface{}{
					-168, -218, 1000, 935,
				},
			},
			"ItalicAngle": 0,
			"Ascent":      683,
			"Descent":     -217,
			"CapHeight":   676,
		},
	},
}

var TimesItalic = Font{
	FontName:  "/Times-Italic",
	Name:      "timesitalic",
	FirstChar: 32,
	LastChar:  251,
	Widths: Array{
		Items: []interface{}{
			250, 333, 408, 500, 500, 833, 778, 333, 333, 333, 500, 564, 250, 333, 250, 278, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 278, 278, 564, 564, 564, 444, 921, 722, 667, 667, 722, 611, 556, 722, 722, 333, 389, 722, 611, 889, 722, 722, 556, 722, 667, 556, 611, 722, 722, 944, 722, 722, 611, 333, 278, 333, 469, 500, 333, 444, 500, 444, 500, 444, 333, 500, 500, 278, 278, 500, 278, 778, 500, 500, 500, 500, 333, 389, 278, 500, 500, 722, 500, 500, 444, 480, 200, 480, 541, 333, 500, 500, 167, 500, 500, 500, 500, 180, 444, 500, 333, 333, 556, 556, 500, 500, 500, 250, 453, 350, 333, 444, 444, 500, 1000, 1000, 444, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 1000, 889, 276, 611, 722, 889, 310, 667, 278, 278, 500, 722, 500,
		},
	},
	Descriptor: Dictionary{
		M: map[string]interface{}{
			"Type":     "/FontDescriptor",
			"FontName": "/Times-Italic",
			"Flags":    0b00000000000000000000000000000010,
			"FontBBox": Array{
				Items: []interface{}{
					-169, -217, 1010, 883,
				},
			},
			"ItalicAngle": -15,
			"Ascent":      683,
			"Descent":     -217,
			"CapHeight":   669,
		},
	},
}

var TimesBoldItalic = Font{
	FontName:  "/Times-BoldItalic",
	Name:      "timesbolditalic",
	FirstChar: 32,
	LastChar:  251,
	Widths: Array{
		Items: []interface{}{
			250, 389, 555, 500, 500, 833, 778, 333, 333, 333, 500, 570, 250, 333, 250, 278, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 333, 333, 570, 570, 570, 500, 832, 667, 667, 667, 722, 667, 667, 722, 778, 389, 500, 667, 611, 889, 722, 722, 611, 722, 667, 556, 611, 722, 667, 889, 667, 611, 611, 333, 278, 333, 570, 500, 333, 500, 500, 444, 500, 444, 333, 500, 556, 278, 278, 500, 278, 778, 556, 500, 500, 500, 389, 389, 278, 556, 444, 667, 500, 444, 389, 348, 220, 348, 570, 389, 500, 500, 167, 500, 500, 500, 500, 278, 500, 500, 333, 333, 556, 556, 500, 500, 500, 250, 500, 350, 333, 500, 500, 500, 1000, 1000, 500, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 1000, 944, 266, 611, 722, 944, 300, 722, 278, 278, 500, 722, 500,
		},
	},
	Descriptor: Dictionary{
		M: map[string]interface{}{
			"Type":     "/FontDescriptor",
			"FontName": "/Times-BoldItalic",
			"Flags":    0b00000000000000000000000000000010,
			"FontBBox": Array{
				Items: []interface{}{
					-200, -218, 996, 921,
				},
			},
			"ItalicAngle": -15,
			"Ascent":      683,
			"Descent":     -217,
			"CapHeight":   676,
		},
	},
}

var HelveticaRegular = Font{
	FontName:  "/Helvetica",
	Name:      "helvetica",
	FirstChar: 32,
	LastChar:  251,
	Widths: Array{
		Items: []interface{}{
			278, 278, 355, 556, 556, 889, 667, 222, 333, 333, 389, 584, 278, 333, 278, 278, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 278, 278, 584, 584, 584, 556, 1015, 667, 667, 722, 722, 667, 611, 778, 722, 278, 500, 667, 556, 833, 722, 778, 667, 778, 722, 667, 611, 722, 667, 944, 667, 667, 611, 278, 278, 278, 469, 556, 222, 556, 556, 500, 556, 556, 278, 556, 556, 222, 222, 500, 222, 833, 556, 556, 556, 556, 333, 500, 278, 556, 500, 722, 500, 500, 500, 334, 260, 334, 584, 333, 556, 556, 167, 556, 556, 556, 556, 191, 333, 556, 333, 333, 500, 500, 556, 556, 556, 278, 537, 350, 222, 333, 333, 556, 1000, 1000, 611, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 333, 1000, 1000, 370, 556, 778, 1000, 365, 889, 278, 222, 611, 944, 611,
		},
	},
	Descriptor: Dictionary{
		M: map[string]interface{}{
			"Type":     "/FontDescriptor",
			"FontName": "/Helvetica",
			"Flags":    0b00000000000000000000000000000010,
			"FontBBox": Array{
				Items: []interface{}{
					-166, -225, 1000, 931,
				},
			},
			"ItalicAngle": 0,
			"Ascent":      718,
			"Descent":     -207,
			"CapHeight":   718,
		},
	},
}

var CourierMono = Font{
	FontName:  "/Courier",
	Name:      "courier",
	FirstChar: 32,
	LastChar:  251,
	Widths: Array{
		Items: []interface{}{
			500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500,
		},
	},
	Descriptor: Dictionary{
		M: map[string]interface{}{
			"Type":     "/FontDescriptor",
			"FontName": "/Courier",
			"Flags":    0b00000000000000000000000000000010,
			"FontBBox": Array{
				Items: []interface{}{
					-23, -250, 715, 805,
				},
			},
			"ItalicAngle": 0,
			"Ascent":      629,
			"Descent":     -157,
			"CapHeight":   562,
		},
	},
}
