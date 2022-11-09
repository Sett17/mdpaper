package spec

var LastId = 0

type Object interface {
	ID() int
	Bytes() []byte
	Reference() string
	Pointer() *Object
}

//	type Object struct {
//		ID    int
//		Thing Addable
//		Data  Stream
//	}
//func (o *Object) header() []byte {
//	buf := bytes.Buffer{}
//	buf.WriteString(fmt.Sprintf("%d 0 obj\n", o.ID))
//	return buf.Bytes()
//}
//
//func (o *Object) footer() []byte {
//	buf := bytes.Buffer{}
//	buf.WriteString("endobj\n")
//	return buf.Bytes()
//}
//
//func (o *Object) Reference() string {
//	return fmt.Sprintf("%d 0 R", o.ID)
//}
//
//func (o *Object) Bytes() []byte {
//	buf := bytes.Buffer{}
//	buf.Write(o.header())
//	if !o.Data.isZero() {
//		d, ok := o.Thing.(Dictionary)
//		if !ok {
//			panic("Object.Data can only be used when Object.Thing is a Dictionary")
//		}
//		d.Set("Size", fmt.Sprintf("%d", o.Data.Len()))
//	}
//	buf.Write(o.Thing.Bytes())
//	if !o.Data.isZero() {
//		buf.Write(o.Data.Bytes())
//	}
//	buf.Write(o.footer())
//	return buf.Bytes()
//}
//
//func (o *Object) SplitBytes() ([]byte, []byte, []byte) {
//	return o.header(), o.Thing.Bytes(), o.footer()
//}
//
//func NewObject(thing Addable) Object {
//	LastId++
//	return Object{ID: LastId, Thing: thing}
//}
