package spec

import (
	"bytes"
	"io"
	"mdpaper/globals"
)

type Stream struct {
	RawData bytes.Buffer
}

func (s *Stream) Len() int {
	return s.RawData.Len()
}

func (s *Stream) isZero() bool {
	return s.RawData.Len() == 0
}

func (s *Stream) WriteString(str string) {
	s.RawData.WriteString(str)
}

func (s *Stream) Write(b []byte) {
	s.RawData.Write(b)
}

func (s *Stream) Deflate() {
	r, err := flate.Encode(flate{}, &s.RawData)
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(r)
	s.RawData = *bytes.NewBuffer(b)
	s.RawData.WriteByte(0x0a)
}

func (s *Stream) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("\nstream\n")
	buf.Write(s.RawData.Bytes())
	buf.WriteString("endstream\n")
	return buf.Bytes()
}

type StreamObject struct {
	GenericObject
	Content []*Bytable
	Dictionary
}

func (s *StreamObject) Add(add *Addable) {
	var b Bytable = *add
	s.Content = append(s.Content, &b)
}
func (s *StreamObject) AddBytable(add *Bytable) {
	s.Content = append(s.Content, add)
}

func (s *StreamObject) AddBytes(buf []byte) {
	var b Bytable = &GenericBytable{dat: &buf}
	s.Content = append(s.Content, &b)
}

func (s *StreamObject) Pointer() *Object {
	var z Object = s
	return &z
}

func (s *StreamObject) Bytes() []byte {
	buf := bytes.Buffer{}
	beg, end := s.ByteParts()
	buf.Write(beg)
	stream := Stream{}
	for _, c := range s.Content {
		stream.Write((*c).Bytes())
	}
	if globals.Cfg.Debug {
		s.Dictionary.Set("Length", stream.Len()-1)
	} else {
		stream.Deflate()
		s.Dictionary.Set("Filter", "[/FlateDecode]")
		s.Dictionary.Set("Length", stream.Len())
	}
	buf.Write(s.Dictionary.Bytes())
	buf.Write(stream.Bytes())
	buf.Write(end)
	return buf.Bytes()
}

func NewStreamObject() StreamObject {
	LastId++
	return StreamObject{GenericObject: GenericObject{id: LastId}}
}
