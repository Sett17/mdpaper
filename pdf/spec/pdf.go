package spec

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

type PDF struct {
	objects []*Object
	root    *Object
}

func (p *PDF) AddObject(obj *Object) {
	p.objects = append(p.objects, obj)
}

func (p *PDF) SetRoot(obj *Object) {
	p.root = obj
}

func (p *PDF) Bytes() []byte {
	d := bytes.Buffer{}
	// Header
	d.Write([]byte("%PDF-1.5\n"))
	d.Write([]byte{0x25, 0xEE, 0xEE, 0xEE, 0xEE, 0x0A}) // comment line with non ascii characters so other tools know this file contains binary data

	// Content
	offsets := make(map[int]int)

	offsets[(*p.root).ID()] = d.Len()
	d.Write((*p.root).Bytes())

	sort.SliceStable(p.objects, func(i, j int) bool {
		if so, ok := (*p.objects[i]).(*StreamObject); ok {
			if so.M["Type"] == "XObject" {
				return false
			}
		}
		return (*p.objects[i]).ID() < (*p.objects[j]).ID()
	})
	for _, s := range p.objects {
		offsets[(*s).ID()] = d.Len()
		d.Write((*s).Bytes())
	}

	// Cross-reference table
	xrefOffset := d.Len()
	d.WriteString("xref\n")
	d.WriteString(fmt.Sprintf("0 %d\n", LastId+1))
	d.Write(objZeroXRef().Bytes())
	d.Write(xRefEntry{uint32(offsets[(*p.root).ID()]), 0, 'n'}.Bytes())

	for _, s := range p.objects {
		d.Write(xRefEntry{uint32(offsets[(*s).ID()]), 0, 'n'}.Bytes())
	}

	//Trailer
	d.WriteString("trailer\n")
	trailerDict := NewDict()
	trailerDict.Set("Size", LastId+1)
	trailerDict.Set("Root", (*p.root).Reference())
	d.Write(trailerDict.Bytes())
	d.WriteString("startxref\n")
	d.WriteString(fmt.Sprintf("%d\n", xrefOffset))
	d.WriteString("%%EOF\n")
	return d.Bytes()
}

func (p *PDF) WriteTo(f *os.File) {
	_, err := f.Write(p.Bytes())
	if err != nil {
		panic(err)
	}
}
