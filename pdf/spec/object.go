package spec

import (
	"bytes"
	"fmt"
)

var LastId = 0

type Object interface {
	ID() int
	Bytes() []byte
	Reference() string
	Pointer() *Object
}

type Bytable interface {
	Bytes() []byte
}

type GenericObject struct {
	id int
}

func (g *GenericObject) ID() int {
	return g.id
}

func (g *GenericObject) Bytes(b Bytable) []byte {
	buf := bytes.Buffer{}
	beg, end := g.ByteParts()
	buf.Write(beg)
	byt := b.Bytes()
	buf.Write(byt)
	if byt[len(byt)-1] != '\n' {
		buf.WriteString("\n")
	}
	buf.Write(end)
	return buf.Bytes()
}

func (g *GenericObject) Reference() string {
	return fmt.Sprintf("%d 0 R", g.id)
}

func (g *GenericObject) ByteParts() ([]byte, []byte) {
	return []byte(fmt.Sprintf("%d 0 obj\n", g.id)), []byte("endobj\n\n")
}
