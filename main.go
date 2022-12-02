package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/nickng/bibtex"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/goldmark-cite"
	"github.com/sett17/mdpaper/pdf"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/text"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"time"
)

//TODO custom errors

func main() {
	file := "paper.md"
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
		goldmark.WithExtensions(
			&goldmark_cite.Extender{Indices: &globals.BibIndices},
			meta.Meta, // just to ignore frontmatter
		),
		goldmark.WithParserOptions()).Parser()
	ast := p.Parse(text.NewReader(inp))

	//ast.Dump(inp, 0)
	//return

	parsed := time.Now()
	fmt.Printf("Parsed in %v\n", parsed.Sub(start))

	cfgFile, err := os.ReadFile(configFile)
	configT := time.Now()
	if err == nil {
		err = yaml.Unmarshal(cfgFile, &globals.Cfg)
		configT = time.Now()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Loaded config from %s in %v\n", configFile, configT.Sub(parsed))
	}
	//else {
	//	if os.IsNotExist(err) {
	out, err := yaml.Marshal(globals.Cfg)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(configFile, out, 0644)
	if err != nil {
		panic(err)
	}
	configT = time.Now()
	fmt.Printf("Created config file %s\n", configFile)
	//	}
	//}

	if globals.Cfg.Citation.Enabled {
		bibFile, err := os.OpenFile(globals.Cfg.Citation.File, os.O_RDONLY, 0644)
		if err == nil {
			bt, err := bibtex.Parse(bibFile)
			if err != nil {
				panic(err)
			}
			for _, v := range bt.Entries {
				globals.Bibs[v.CiteName] = v
			}
		}
		citT := time.Now()
		fmt.Printf("Bibtex loaded in %v\n\n", citT.Sub(configT))
	}

	pp := pdf.FromAst(ast)

	ppT := time.Now()
	fmt.Printf("PDF generated in %v\n", ppT.Sub(parsed))

	outName := strings.ReplaceAll(globals.Cfg.Paper.Title, " ", "_") + ".pdf"
	outp, err := os.Create(outName)
	if err != nil {
		panic(err)
	}
	beforeWrite := time.Now()
	pp.WriteFile(outp)
	doneWrite := time.Now()

	fi, err := outp.Stat()
	fmt.Printf("%s of PDF put into %s, in %v\n", humanize.Bytes(uint64(fi.Size())), outName, doneWrite.Sub(beforeWrite))

	fmt.Printf("Total time: %v\n", doneWrite.Sub(start))

	dbgOut, err := os.Create("debug.txt")
	if err != nil {
		panic(err)
	}
	globals.Cfg.Paper.Debug = true
	pp.WriteDebug(dbgOut)
	fi, err = dbgOut.Stat()
	//if err == nil {
	//fmt.Printf("Debug written without stream: %s\n", humanize.Bytes(uint64(fi.Size())))
	//}

}
