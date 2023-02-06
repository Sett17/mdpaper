package spec

import (
	"crypto/md5"
	"fmt"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"os"
	"sort"
	"time"
)

type PDF struct {
	objects []*Object
	Root    string
	Info    string
}

func (p *PDF) AddObject(obj ...*Object) {
	p.objects = append(p.objects, obj...)
}

type writeTracker struct {
	*os.File
	L int
}

func (w *writeTracker) Write(b []byte) {
	l, err := w.File.Write(b)
	if err != nil {
		cli.Error(err, true)
	}
	w.L += l
}

func (w *writeTracker) WriteString(s string) {
	l, err := w.File.WriteString(s)
	if err != nil {
		cli.Error(err, true)
	}
	w.L += l
}

func (p *PDF) WriteFile(f *os.File) {
	d := &writeTracker{File: f}
	//region HEADER
	d.Write([]byte("%PDF-1.5\n"))

	d.Write([]byte{0x25, 0xe2, 0xe3, 0xcf, 0xd3, 0x0a}) // comment line with non ascii characters so other tools know this file contains binary data
	//endregion

	//region CONTENT
	offsets := make(map[int]int)

	sort.SliceStable(p.objects, func(i, j int) bool {
		if so, ok := (*p.objects[i]).(*StreamObject); ok {
			if so.M["Type"] == "XObject" {
				return false
			}
		}
		return (*p.objects[i]).ID() < (*p.objects[j]).ID()
	})
	for _, s := range p.objects {
		offsets[(*s).ID()] = d.L
		d.Write((*s).Bytes())
	}
	//endregion

	//region XREF
	xrefOffset := d.L
	d.WriteString("xref\n")
	d.WriteString(fmt.Sprintf("0 %d\n", LastId+1))
	d.Write(objZeroXRef().Bytes())

	for _, s := range p.objects {
		d.Write(xRefEntry{uint32(offsets[(*s).ID()]), 0, 'n'}.Bytes())
	}
	//endregion

	//region TRAILER
	d.WriteString("trailer\n")
	trailerDict := NewDict()
	trailerDict.Set("Size", LastId+1)
	trailerDict.Set("Root", p.Root)
	//trailerDict.Set("Info", p.Info)
	//calculate ID
	ID := fmt.Sprintf("<%X>", md5.Sum([]byte(fmt.Sprintf("%s%s%s%d", time.Now(), globals.Cfg.Paper.Author, globals.Cfg.Paper.Title, LastId))))
	trailerDict.Set("ID", Array{Items: []interface{}{ID, ID}})
	d.Write(trailerDict.Bytes())
	d.WriteString("startxref\n")
	d.WriteString(fmt.Sprintf("%d\n", xrefOffset))
	d.WriteString("%%EOF\n")
	//endregion
}

func (p *PDF) WriteDebug(f *os.File) {
	d := &writeTracker{File: f}
	//region HEADER DBG
	d.Write([]byte("%PDF-1.5\n"))

	d.Write([]byte{0x25, 0xe2, 0xe3, 0xcf, 0xd3, 0x0a}) // comment line with non ascii characters so other tools know this file contains binary data
	//endregion

	//region CONTENT DBG
	offsets := make(map[int]int)

	filtered := make([]*Object, 0, len(p.objects))
	for _, s := range p.objects {
		switch v := (*s).(type) {
		case *StreamObject:
			if v.M["Length1"] != nil { //exclude fonts
				continue
			}
		case *XObject:
			continue
		}
		filtered = append(filtered, s)
	}

	sort.SliceStable(filtered, func(i, j int) bool {
		return (*filtered[i]).ID() < (*filtered[j]).ID()
	})
	for _, s := range filtered {
		offsets[(*s).ID()] = d.L
		d.Write((*s).Bytes())
	}
	//endregion

	//region XREF DBG
	xrefOffset := d.L
	d.WriteString("xref\n")
	d.WriteString(fmt.Sprintf("0 %d\n", LastId+1))
	d.Write(objZeroXRef().Bytes())

	for _, s := range filtered {
		d.Write(xRefEntry{uint32(offsets[(*s).ID()]), 0, 'n'}.Bytes())
	}
	//endregion

	//region TRAILER DBG
	d.WriteString("trailer\n")
	trailerDict := NewDict()
	trailerDict.Set("Size", LastId+1)
	trailerDict.Set("Root", p.Root)
	trailerDict.Set("Info", p.Info)
	//calculate ID
	ID := fmt.Sprintf("<%X>", md5.Sum([]byte(fmt.Sprintf("%s%s%s%d", time.Now(), globals.Cfg.Paper.Author, globals.Cfg.Paper.Title, LastId))))
	trailerDict.Set("ID", Array{Items: []interface{}{ID, ID}})
	d.Write(trailerDict.Bytes())
	d.WriteString("\nstartxref\n")
	d.WriteString(fmt.Sprintf("%d\n", xrefOffset))
	d.WriteString("%%EOF\n")
	//endregion
}
