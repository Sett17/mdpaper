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
	GenericObject
	Array
}

func NewArrayObject() ArrayObject {
	LastId++
	return ArrayObject{GenericObject{id: LastId}, NewArray()}
}

func (a *ArrayObject) Bytes() []byte {
	return a.GenericObject.Bytes(&a.Array)
}

func (a *ArrayObject) Pointer() *Object {
	var o Object = a
	return &o
}
