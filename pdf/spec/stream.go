package spec

import (
	"bytes"
	"io"
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
	r, err := ascii85Decode.Encode(ascii85Decode{}, &s.RawData)
	//r, err := flate.Encode(flate{}, &s.RawData)
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(r)
	s.RawData = *bytes.NewBuffer(b)
}

func (s *Stream) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("stream\n")
	buf.Write(s.RawData.Bytes())
	buf.WriteString("endstream\n")
	return buf.Bytes()
}

type StreamObject struct {
	GenericObject
	Deflate bool
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
	if s.Deflate {
		stream.Deflate()
		//s.Dictionary.Set("Filter", "FlateDecode")
		s.Dictionary.Set("Filter", "ASCII85Decode")
		s.Dictionary.Set("Length", stream.Len())
	} else {
		//s.Dictionary.Set("Length", stream.Len())
		s.Dictionary.Set("Length", stream.Len()-1)
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
