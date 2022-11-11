package spec

import (
	"bytes"
	"fmt"
	"sort"
)

type Dictionary struct {
	M map[string]interface{}
}

func (d Dictionary) String() string {
	return string(d.Bytes())
}

func NewDict() Dictionary {
	return Dictionary{M: map[string]interface{}{}}
}

func (d *Dictionary) Set(key string, val interface{}) {
	if d.M == nil {
		d.M = make(map[string]interface{})
	}
	d.M[key] = val
}

func (d *Dictionary) Get(key string) interface{} {
	return d.M[key]
}

func (d *Dictionary) Bytes() []byte {
	buf := bytes.Buffer{}
	if len(d.M) == 0 {
		buf.WriteString("<< >>")
		return buf.Bytes()
	}
	buf.WriteString("<<")
	resKeys := make([]string, 0)
	for k := range d.M {
		resKeys = append(resKeys, k)
	}
	sort.Strings(resKeys)
	for _, key := range resKeys {
		buf.WriteString(fmt.Sprintf("\n/%s %v", key, d.M[key]))
	}
	buf.WriteString(">>\n")
	return buf.Bytes()
}

type DictionaryObject struct {
	id int
	Dictionary
}

func (d *DictionaryObject) ID() int {
	return d.id
}

func (d *DictionaryObject) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%d 0 obj\n", d.id))
	buf.Write(d.Dictionary.Bytes())
	buf.WriteString("endobj\n")
	return buf.Bytes()
}

func (d *DictionaryObject) Reference() string {
	return fmt.Sprintf("%d 0 R", d.id)
}

func NewDictObject() DictionaryObject {
	LastId++
	return DictionaryObject{id: LastId}
}

func (d *DictionaryObject) Pointer() *Object {
	var z Object = d
	return &z
}
