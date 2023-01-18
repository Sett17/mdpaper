package conversions

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/sett17/mdpaper/cli"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf/conversions/options"
	"github.com/sett17/mdpaper/pdf/elements"
	"github.com/sett17/mdpaper/pdf/spec"
	"github.com/yuin/goldmark/ast"
	"os"
	"os/exec"
	"strings"
)

func Mermaid(fcb *ast.FencedCodeBlock) (retO *spec.XObject, retA *spec.Addable) {
	optionString := strings.TrimPrefix(string(fcb.Info.Text(globals.File)), "mermaid")
	opts, err := options.Parse(optionString)
	if err != nil {
		cli.Error(fmt.Errorf("error parsing mermaid options: %w", err), false)
		cli.Warning(optionString)
	}

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
	mul := 1.0
	if f, ok := opts.GetFloat("width"); ok {
		mul = f
	}
	io, ia := spec.NewImageObjectFromFile(inputFile.Name()+".png", mul)
	retO = &io
	retA = &ia
	defer os.Remove(inputFile.Name() + ".png")
	return
}

func Code(fcb *ast.FencedCodeBlock) *spec.Addable {
	optionString := ""
	if fcb.Info != nil {
		optionString = options.Extract(string(fcb.Info.Text(globals.File)))
	}

	opts, err := options.Parse(optionString)
	if err != nil {
		cli.Error(fmt.Errorf("error parsing code options: %w", err), false)
	}

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

	ln := globals.Cfg.Code.LineNumbers
	if b, ok := opts.GetBool("lineNumbers"); ok {
		ln = b
	} else if b, ok := opts.GetBool("linenumbers"); ok {
		ln = b
	} else if b, ok := opts.GetBool("ln"); ok {
		ln = b
	}
	fs := globals.Cfg.Code.FontSize
	if f, ok := opts.GetInt("fontSize"); ok {
		fs = f
	} else if f, ok := opts.GetInt("fontsize"); ok {
		fs = f
	} else if f, ok := opts.GetInt("fs"); ok {
		fs = f
	}
	sn := 0
	if i, ok := opts.GetInt("startNumber"); ok {
		sn = i - 1
	} else if i, ok := opts.GetInt("startnumber"); ok {
		sn = i - 1
	} else if i, ok := opts.GetInt("sn"); ok {
		sn = i - 1
	}
	fc := elements.FencedCode{
		Tokens:      toks,
		Style:       styles.Get(globals.Cfg.Code.Style),
		LineNumbers: ln,
		FontSize:    fs,
		StartNumber: sn,
	}

	var a spec.Addable = &fc
	return &a
}
