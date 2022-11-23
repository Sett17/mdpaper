/*
Copyright 2018 The pdfcpu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/*
This file contains parts of multiple files, which are licensed under above license.
*/

package spec

import (
	"bytes"
	"compress/zlib"
	"encoding/ascii85"
	"errors"
	"fmt"
	"io"
)

type Filter interface {
	Encode(r io.Reader) (io.Reader, error)
	Decode(r io.Reader) (io.Reader, error)
}
type baseFilter struct {
	parms map[string]int
}

type ascii85Decode struct {
	baseFilter
}

const eodASCII85 = "~>"

// Encode implements encoding for an ASCII85Decode filter.
func (f ascii85Decode) Encode(r io.Reader) (io.Reader, error) {

	p, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	encoder := ascii85.NewEncoder(buf)
	encoder.Write(p)
	encoder.Close()

	// Add eod sequence
	buf.WriteString(eodASCII85)

	return buf, nil
}

// Decode implements decoding for an ASCII85Decode filter.
func (f ascii85Decode) Decode(r io.Reader) (io.Reader, error) {

	p, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("dump:\n%s", hex.Dump(p))

	l := len(p)
	if p[l-1] == 0x0A || p[l-1] == 0x0D {
		p = p[:l-1]
	}

	if !bytes.HasSuffix(p, []byte(eodASCII85)) {
		return nil, errors.New("pdfcpu: Decode: missing eod marker")
	}

	// Strip eod sequence: "~>"
	p = p[:len(p)-2]

	decoder := ascii85.NewDecoder(bytes.NewReader(p))

	buf, err := io.ReadAll(decoder)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(buf), nil
}

// flate starts here

// Portions of this code are based on ideas of image/png: reader.go:readImagePass
// PNG is documented here: www.w3.org/TR/PNG-Filters.html

// PDF allows a prediction step prior to compression applying TIFF or PNG prediction.
// Predictor algorithm.
const (
	PredictorNo      = 1  // No prediction.
	PredictorTIFF    = 2  // Use TIFF prediction for all rows.
	PredictorNone    = 10 // Use PNGNone for all rows.
	PredictorSub     = 11 // Use PNGSub for all rows.
	PredictorUp      = 12 // Use PNGUp for all rows.
	PredictorAverage = 13 // Use PNGAverage for all rows.
	PredictorPaeth   = 14 // Use PNGPaeth for all rows.
	PredictorOptimum = 15 // Use the optimum PNG prediction for each row.
)

// For predictor > 2 PNG filters (see RFC 2083) get applied and the first byte of each pixelrow defines
// the prediction algorithm used for all pixels of this row.
const (
	PNGNone    = 0x00
	PNGSub     = 0x01
	PNGUp      = 0x02
	PNGAverage = 0x03
	PNGPaeth   = 0x04
)

type flate struct {
	baseFilter
}

// Encode implements encoding for a Flate filter.
func (f flate) Encode(r io.Reader) (io.Reader, error) {

	// TODO Optional decode parameters may need predictor preprocessing.

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	defer w.Close()

	_, err := io.Copy(w, r)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// Decode implements decoding for a Flate filter.
func (f flate) Decode(r io.Reader) (io.Reader, error) {

	rc, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	// Optional decode parameters need postprocessing.
	return f.decodePostProcess(rc)
}

func passThru(rin io.Reader) (*bytes.Buffer, error) {
	var b bytes.Buffer
	_, err := io.Copy(&b, rin)
	return &b, err
}

func intMemberOf(i int, list []int) bool {
	for _, v := range list {
		if i == v {
			return true
		}
	}
	return false
}

func applyHorDiff(row []byte, colors int) ([]byte, error) {
	// This works for 8 bits per color only.
	for i := 1; i < len(row)/colors; i++ {
		for j := 0; j < colors; j++ {
			row[i*colors+j] += row[(i-1)*colors+j]
		}
	}
	return row, nil
}

func processRow(pr, cr []byte, p, colors, bytesPerPixel int) ([]byte, error) {

	//fmt.Printf("pr(%v) =\n%s\n", &pr, hex.Dump(pr))
	//fmt.Printf("cr(%v) =\n%s\n", &cr, hex.Dump(cr))

	if p == PredictorTIFF {
		return applyHorDiff(cr, colors)
	}

	// Apply the filter.
	cdat := cr[1:]
	pdat := pr[1:]

	// Get row filter from 1st byte
	f := int(cr[0])

	// The value of Predictor supplied by the decoding filter need not match the value
	// used when the data was encoded if they are both greater than or equal to 10.

	switch f {

	case PNGNone:
		// No operation.

	case PNGSub:
		for i := bytesPerPixel; i < len(cdat); i++ {
			cdat[i] += cdat[i-bytesPerPixel]
		}

	case PNGUp:
		for i, p := range pdat {
			cdat[i] += p
		}

	case PNGAverage:
		// The average of the two neighboring pixels (left and above).
		// Raw(x) - floor((Raw(x-bpp)+Prior(x))/2)
		for i := 0; i < bytesPerPixel; i++ {
			cdat[i] += pdat[i] / 2
		}
		for i := bytesPerPixel; i < len(cdat); i++ {
			cdat[i] += uint8((int(cdat[i-bytesPerPixel]) + int(pdat[i])) / 2)
		}

	}

	return cdat, nil
}

func (f flate) parameters() (colors, bpc, columns int, err error) {

	// Colors, int
	// The number of interleaved colour components per sample.
	// Valid values are 1 to 4 (PDF 1.0) and 1 or greater (PDF 1.3). Default value: 1.
	// Used by PredictorTIFF only.
	colors, found := f.parms["Colors"]
	if !found {
		colors = 1
	} else if colors == 0 {
		return 0, 0, 0, fmt.Errorf("pdfcpu: filter FlateDecode: \"Colors\" must be > 0")
	}

	// BitsPerComponent, int
	// The number of bits used to represent each colour component in a sample.
	// Valid values are 1, 2, 4, 8, and (PDF 1.5) 16. Default value: 8.
	// Used by PredictorTIFF only.
	bpc, found = f.parms["BitsPerComponent"]
	if !found {
		bpc = 8
	} else if !intMemberOf(bpc, []int{1, 2, 4, 8, 16}) {
		return 0, 0, 0, fmt.Errorf("pdfcpu: filter FlateDecode: Unexpected \"BitsPerComponent\": %d", bpc)
	}

	// Columns, int
	// The number of samples in each row. Default value: 1.
	columns, found = f.parms["Columns"]
	if !found {
		columns = 1
	}

	return colors, bpc, columns, nil
}

// decodePostProcess
func (f flate) decodePostProcess(r io.Reader) (io.Reader, error) {

	predictor, found := f.parms["Predictor"]
	if !found || predictor == PredictorNo {
		return passThru(r)
	}

	if !intMemberOf(
		predictor,
		[]int{PredictorTIFF,
			PredictorNone,
			PredictorSub,
			PredictorUp,
			PredictorAverage,
			PredictorPaeth,
			PredictorOptimum,
		}) {
		return nil, fmt.Errorf("pdfcpu: filter FlateDecode: undefined \"Predictor\" %d", predictor)
	}

	colors, bpc, columns, err := f.parameters()
	if err != nil {
		return nil, err
	}

	bytesPerPixel := (bpc*colors + 7) / 8
	rowSize := (bpc*colors*columns + 7) / 8

	m := rowSize
	if predictor != PredictorTIFF {
		// PNG prediction uses a row filter byte prefixing the pixelbytes of a row.
		m++
	}

	// cr and pr are the bytes for the current and previous row.
	cr := make([]byte, m)
	pr := make([]byte, m)

	// Output buffer
	var b bytes.Buffer

	for {

		// Read decompressed bytes for one pixel row.
		n, err := io.ReadFull(r, cr)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			// eof
			if n == 0 {
				break
			}
		}

		if n != m {
			return nil, fmt.Errorf("pdfcpu: filter FlateDecode: read error, expected %d bytes, got: %d", m, n)
		}

		d, err1 := processRow(pr, cr, predictor, colors, bytesPerPixel)
		if err1 != nil {
			return nil, err1
		}

		_, err1 = b.Write(d)
		if err1 != nil {
			return nil, err1
		}

		if err == io.EOF {
			break
		}

		// Swap byte slices.
		pr, cr = cr, pr
	}

	if b.Len()%rowSize > 0 {
		return nil, fmt.Errorf("pdfcpu: filter FlateDecode: postprocessing failed")
	}

	return &b, nil
}