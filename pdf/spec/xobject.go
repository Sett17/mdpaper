package spec

import (
	"bytes"
	"strings"
)

type XObject struct {
	GenericObject
	Dictionary
	Stream
	Name string
}

func (x *XObject) Bytes() []byte {
	buf := bytes.Buffer{}
	beg, end := x.ByteParts()
	buf.Write(beg)
	buf.Write(x.Dictionary.Bytes())
	buf.Write(x.Stream.Bytes())
	buf.Write(end)
	return buf.Bytes()
}

func (x *XObject) Pointer() *Object {
	var z Object = x
	return &z
}

func NewXObject(name string) XObject {
	LastId++
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, ".", "_")
	name = strings.ReplaceAll(name, ":", "_")
	name = strings.ToLower(name)
	return XObject{GenericObject: GenericObject{id: LastId}, Name: name}
}
