package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/pdf"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
	"time"
)

//TODO custom errors

func main() {
	file := "paper_simple.md"
	configFile := "config.yaml"
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
		//goldmark.WithExtensions(meta.New(
		//	meta.WithStoresInDocument(),
		//)),
		goldmark.WithParserOptions()).Parser()
	ast := p.Parse(text.NewReader(inp))
	//ast.Dump(inp, 0)
	fmt.Printf("Parsed in %v\n", time.Since(start))
	cfgFile, err := os.ReadFile(configFile)
	if err == nil {
		err = yaml.Unmarshal(cfgFile, &globals.Cfg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Loaded config from %s\n", configFile)
	} else {
		if os.IsNotExist(err) {
			//create cfg file
			out, err := yaml.Marshal(globals.Cfg)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(configFile, out, 0644)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Created config file %s\n", configFile)
		}
	}
	pp := pdf.FromAst(ast)
	outName := strings.ReplaceAll(globals.Cfg.Paper.Title, " ", "_") + ".pdf"
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
	globals.Cfg.Paper.Debug = true
	pp.WriteDebug(dbgOut)
	fi, err = dbgOut.Stat()
	if err == nil {
		fmt.Printf("Debug written without stream size: %s\n", humanize.Bytes(uint64(fi.Size())))
	}

}
