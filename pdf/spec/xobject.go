package spec

import (
	"strings"
)

type XObject struct {
	StreamObject
	Name string
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
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "\\", "_")
	name = strings.ToLower(name)
	return XObject{StreamObject{GenericObject: GenericObject{id: LastId}}, name}
}
