package conversions

import (
	"bytes"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/sett17/mdpaper/cli"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"os"
	"os/exec"
)

func Mermaid(fcb *ast.FencedCodeBlock) (retO *spec.XObject, retA *spec.Addable) {
	inputFile, err := os.CreateTemp("", "mdpapermmd")
	if err != nil {
		cli.Error(err, true)
	}
	defer os.Remove(inputFile.Name())
	buf := bytes.Buffer{}
	for i := 0; i < fcb.Lines().Len(); i++ {
		at := fcb.Lines().At(i)
		buf.Write(at.Value(globals.File))
	}
	_, err = inputFile.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	inputFile.Close()
	err = exec.Command("mmdc", "-i", inputFile.Name(), "-o", inputFile.Name()+".png", "-w", "1000").Run()
	if err != nil {
		cli.Warning("mmdc failed")
	}
	io, ia := spec.NewImageObjectFromFile(inputFile.Name()+".png", 1.0)
	retO = &io
	retA = &ia
	defer os.Remove(inputFile.Name() + ".png")
	return
}

func Code(fcb *ast.FencedCodeBlock) *spec.Addable {
	lexer := lexers.Get(string(fcb.Language(globals.File)))
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	text := bytes.Buffer{}
	for i := 0; i < fcb.Lines().Len(); i++ {
		at := fcb.Lines().At(i)
		text.Write(at.Value(globals.File))
	}
	toks, err := lexer.Tokenise(nil, text.String())
	if err != nil {
		cli.Error(err, true)
	}

	fc := elements.FencedCode{
		Tokens: toks,
		Style:  styles.Get(globals.Cfg.Code.Style),
	}

	var a spec.Addable = &fc
	return &a
}
