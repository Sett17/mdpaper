package spec

import (
	"bytes"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"io"
)

type Stream struct {
	data    bytes.Buffer
	Content []*Bytable
}

func (s *Stream) Add(add *Addable) {
	var b Bytable = *add
	s.Content = append(s.Content, &b)
}
func (s *Stream) AddBytable(add *Bytable) {
	s.Content = append(s.Content, add)
}

func (s *Stream) Len() int {
	return s.data.Len()
}

func (s *Stream) isZero() bool {
	return s.data.Len() == 0
}

func (s *Stream) WriteString(str string) {
	var b Bytable = &GenericBytable{dat: []byte(str)}
	s.Content = append(s.Content, &b)
}

func (s *Stream) Write(buf []byte) {
	var b Bytable = &GenericBytable{dat: buf}
	s.Content = append(s.Content, &b)
}

func (s *Stream) Commit() {
	s.data.Reset()
	for _, b := range s.Content {
		s.data.Write((*b).Bytes())
	}
}

func (s *Stream) Deflate() {
	r, err := flate.Encode(flate{}, &s.data)
	if err != nil {
		cli.Error(err, true)
	}
	b, _ := io.ReadAll(r)
	s.data = *bytes.NewBuffer(b)
	s.data.WriteByte(0x0a)
}

func (s *Stream) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("\nstream\n")
	buf.Write(s.data.Bytes())
	buf.WriteString("endstream\n")
	return buf.Bytes()
}

type StreamObject struct {
	GenericObject
	Stream
	Dictionary
	AlwaysDeflate bool
}

func (s *StreamObject) Pointer() *Object {
	var z Object = s
	return &z
}

func (s *StreamObject) Bytes() []byte {
	buf := bytes.Buffer{}
	beg, end := s.ByteParts()
	buf.Write(beg)
	s.Stream.Commit()
	if globals.Cfg.Paper.Debug && !s.AlwaysDeflate {
		s.Dictionary.Set("Length", s.Stream.Len()-1)
	} else {
		s.Stream.Deflate()
		s.Dictionary.Set("Filter", "/FlateDecode")
		s.Dictionary.Set("Length", s.Stream.Len())
	}
	buf.Write(s.Dictionary.Bytes())
	buf.Write(s.Stream.Bytes())
	buf.Write(end)
	return buf.Bytes()
}

func NewStreamObject() StreamObject {
	LastId++
	return StreamObject{GenericObject: GenericObject{id: LastId}}
}
