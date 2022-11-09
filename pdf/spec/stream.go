package spec

import (
	"bytes"
	"compress/flate"
	"fmt"
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
	var b bytes.Buffer
	fl, _ := flate.NewWriter(&b, flate.NoCompression)
	//fl := zlib.NewWriter(&b)
	fl.Write(s.RawData.Bytes())
	fl.Close()
	s.RawData = b
}

func (s *Stream) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("stream\n")
	buf.Write(s.RawData.Bytes())
	buf.WriteString("\nendstream\n")
	return buf.Bytes()
}

type StreamObject struct {
	id      int
	Deflate bool
	Content []*Addable
	Dictionary
}

func (s *StreamObject) Add(add *Addable) {
	s.Content = append(s.Content, add)
}

func (s *StreamObject) Pointer() *Object {
	var z Object = s
	return &z
}

func (s *StreamObject) ID() int {
	return s.id
}

func (s *StreamObject) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%d 0 obj\n", s.id))
	stream := Stream{}
	for _, c := range s.Content {
		stream.Write((*c).Bytes())
	}
	if s.Deflate {
		stream.Deflate()
		s.Dictionary.Set("Filter", "FlateDecode")
		s.Dictionary.Set("Length", stream.Len())
	} else {
		s.Dictionary.Set("Size", stream.Len())
	}
	buf.Write(s.Dictionary.Bytes())
	buf.Write(stream.Bytes())
	buf.WriteString("endobj\n")
	return buf.Bytes()
}

func (s *StreamObject) Reference() string {
	return fmt.Sprintf("%d 0 R", s.id)
}

func NewStreamObject() StreamObject {
	LastId++
	return StreamObject{id: LastId}
}
