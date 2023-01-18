package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/nickng/bibtex"
	"github.com/sett17/mdpaper/cli"
	"github.com/sett17/mdpaper/globals"
	"github.com/sett17/mdpaper/goldmark-cite"
	goldmark_math "github.com/sett17/mdpaper/goldmark-math"
	"github.com/sett17/mdpaper/pdf"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/text"
	"os"
	"strings"
	"time"
)

func main() {
	cfmt.RegisterStyle("purple", func(s string) string {
		return cfmt.Sprintf("{{%s}}::#BA8EF7", s)
	})

	cli.ParseForHelp(os.Args[1:])
	cli.ParseForVersion(os.Args[1:])

	mdFiles := cli.Parse(os.Args[1:])
	if len(mdFiles) == 0 {
		cli.Error(fmt.Errorf("no input file"), false)
		cli.HelpProgArg.Func("")
	}

	if !globals.DidConfig {
		cli.CfgFunc("config.yaml")
	}

	cli.Output("Input files: %q\n", mdFiles)

	for _, file := range mdFiles {
		inp, err := os.ReadFile(file)
		if err != nil {
			cli.Error(err, false)
			continue
		}
		globals.File = append(globals.File, inp...)
		globals.File = append(globals.File, 0x0a)
	}

	start := time.Now()

	p := goldmark.New(
		goldmark.WithExtensions(
			&goldmark_cite.Extender{Indices: &globals.BibIndices},
			&goldmark_math.Extender{},
			meta.Meta, // just to ignore frontmatter
		),
		goldmark.WithParserOptions()).Parser()
	ast := p.Parse(text.NewReader(globals.File))

	//ast.Dump(inp, 0)
	//return

	parsed := time.Now()
	cli.Output("Parsed in %v\n", parsed.Sub(start))

	if globals.Cfg.Citation.Enabled {
		bibFile, err := os.OpenFile(globals.Cfg.Citation.File, os.O_RDONLY, 0644)
		if err == nil {
			bt, err := bibtex.Parse(bibFile)
			if err != nil {
				cli.Error(err, true)
			}
			for _, v := range bt.Entries {
				globals.Bibs[v.CiteName] = v
			}
		}
		citT := time.Now()
		cli.Output("Bibtex loaded in %v\n\n", citT.Sub(parsed))
	}

	pp := pdf.FromAst(ast)

	ppT := time.Now()
	cli.Output("PDF generated in %v\n", ppT.Sub(parsed))

	replaceSet := []byte{'#', '<', '>', '$', '+', '%', '!', '`', '&', '*', '\'', '|', '{', '}', '?', '"', '=', '\\', '/', ':', '@', ' '}
	outBytes := []byte(strings.ToLower(globals.Cfg.Paper.Title))
	for i, v := range outBytes {
		for _, r := range replaceSet {
			if v == r {
				outBytes[i] = '_'
			}
		}
	}
	outName := string(outBytes) + ".pdf"
	outp, err := os.Create(outName)
	if err != nil {
		cli.Error(err, true)
	}
	beforeWrite := time.Now()
	pp.WriteFile(outp)
	doneWrite := time.Now()

	fi, err := outp.Stat()
	cli.Output("%s of PDF put into %s, in %v\n", humanize.Bytes(uint64(fi.Size())), outName, doneWrite.Sub(beforeWrite))

	cli.Output("Total time: %v\n", doneWrite.Sub(start))

	if globals.Cfg.Paper.Debug {
		dbgOut, err := os.Create("debug.txt")
		if err != nil {
			cli.Error(err, true)
		}
		globals.Cfg.Paper.Debug = true
		pp.WriteDebug(dbgOut)
		fi, err = dbgOut.Stat()
		cli.Output("%s of debug info put into debug.txt, in %v\n", humanize.Bytes(uint64(fi.Size())), doneWrite.Sub(beforeWrite))
	}
}
