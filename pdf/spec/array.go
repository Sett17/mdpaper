package spec

import (
	"bytes"
	"fmt"
)

type Array struct {
	Items []interface{}
}

func (a Array) String() string {
	return string(a.Bytes())
}

func NewArray() Array {
	return Array{make([]interface{}, 0)}
}

func (a *Array) Add(item ...interface{}) {
	a.Items = append(a.Items, item...)
}

func (a *Array) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	for i, item := range a.Items {
		buf.WriteString(fmt.Sprintf("%v", item))
		if i != len(a.Items)-1 {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("]")
	return buf.Bytes()
}

type ArrayObject struct {
	id int
	Array
}

func NewArrayObject() ArrayObject {
	LastId++
	return ArrayObject{id: LastId}
}

func (a *ArrayObject) Pointer() *Object {
	var z Object = a
	return &z
}

func (a *ArrayObject) ID() int {
	return a.id
}

func (a *ArrayObject) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%d 0 obj\n", a.id))
	buf.Write(a.Array.Bytes())
	buf.WriteString("\nendobj\n")
	return buf.Bytes()
}

func (a *ArrayObject) Reference() string {
	return fmt.Sprintf("%d 0 R", a.id)
}
