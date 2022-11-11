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
	"time"
)

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
	outp, err := os.Create(globals.Cfg.Title + ".pdf")
	if err != nil {
		panic(err)
	}
	pp.WriteTo(outp)
	fmt.Printf("Done in %v\n", time.Since(start))
	fi, err := outp.Stat()
	if err == nil {
		fmt.Printf("File size: %s\n", humanize.Bytes(uint64(fi.Size())))
	}
	dbgOut, err := os.Create("debug.txt")
	if err != nil {
		panic(err)
	}
	pp.WriteTo(dbgOut)

	//doc := parser.FromFile(inp)
	//global.Log("Parsed markdown file")
	//
	//p := pdf.FromMd(&doc)
	//
	//f1, err := os.Create(parser.FrontMatter.Title + ".pdf")
	//if err != nil {
	//	panic(err)
	//}
	//p.WriteTo(f1)
	//global.Log("Done")
	//f2, err := os.Create(parser.FrontMatter.Title + ".txt")
	//if err != nil {
	//	panic(err)
	//}
	//p.WriteTo(f2)
	//fmt.Printf("%#v\n", doc)
}
