package globals

import (
	citeproc "github.com/sett17/citeproc-js-go"
	"github.com/sett17/citeproc-js-go/csljson"
)

var Citations = make(map[string]csljson.Item)

var Citeproc *citeproc.Session
