package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/text"
	"mdpaper/globals"
	"mdpaper/pdf"
	"os"
	"strings"
	"time"
)

//TODO custom errors

func main() {
	file := "paper_simple.md"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	fmt.Println(file)
	inp, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	globals.File = inp
	start := time.Now()
	p := goldmark.New(
		goldmark.WithExtensions(meta.New(
			meta.WithStoresInDocument(),
		)),
		goldmark.WithParserOptions()).Parser()
	ast := p.Parse(text.NewReader(inp))
	//ast.Dump(inp, 0)
	fmt.Printf("Parsed in %v\n", time.Since(start))
	frontmatter := ast.OwnerDocument().Meta()
	globals.Cfg = globals.FromMap(frontmatter)
	pp := pdf.FromAst(ast)
	outName := strings.ReplaceAll(globals.Cfg.Title, " ", "_") + ".pdf"
	outp, err := os.Create(outName)
	if err != nil {
		panic(err)
	}
	pp.WriteFile(outp)
	fmt.Printf("\nDone in %v\n", time.Since(start))
	fi, err := outp.Stat()
	if err == nil {
		fmt.Printf("File '%s' size: %s\n", outName, humanize.Bytes(uint64(fi.Size())))
	}

	dbgOut, err := os.Create("debug.txt")
	if err != nil {
		panic(err)
	}
	globals.Cfg.Debug = true
	pp.WriteDebug(dbgOut)
	fi, err = dbgOut.Stat()
	if err == nil {
		fmt.Printf("Debug written without stream size: %s\n", humanize.Bytes(uint64(fi.Size())))
	}

}
