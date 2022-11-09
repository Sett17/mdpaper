package spec

import (
	"bytes"
	"fmt"
)

type xRefEntry struct {
	Offset     uint32
	Generation uint16
	Kind       uint8
}

func objZeroXRef() xRefEntry {
	return xRefEntry{
		Offset:     0,
		Generation: 65535,
		Kind:       'f',
	}
}

func (x xRefEntry) Bytes() []byte {
	buf := bytes.Buffer{}

	buf.WriteString(fmt.Sprintf("%010d %05d %c\n", x.Offset, x.Generation, x.Kind))

	return buf.Bytes()
}
