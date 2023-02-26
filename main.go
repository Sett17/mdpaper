package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/goccy/go-json"
	"github.com/i582/cfmt/cmd/cfmt"
	citeproc "github.com/sett17/citeproc-js-go"
	"github.com/sett17/citeproc-js-go/csljson"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/goldmark-cite"
	goldmark_figref "github.com/sett17/mdpaper/v2/goldmark-figref"
	goldmark_math "github.com/sett17/mdpaper/v2/goldmark-math"
	"github.com/sett17/mdpaper/v2/pdf"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
	"os"
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
			&goldmark_cite.CitationExtension{},
			&goldmark_math.Extender{},
			&goldmark_figref.FigRefExtension{},
			meta.Meta, // just to ignore frontmatter
			extension.NewTable(),
		),
		goldmark.WithParserOptions()).Parser()
	ast := p.Parse(text.NewReader(globals.File))

	//ast.Dump(globals.File, 0)
	//return

	parsed := time.Now()
	cli.Output("Parsed in %v\n", parsed.Sub(start))

	if globals.Cfg.Citation.Enabled {
		globals.Citeproc = citeproc.NewSession()

		if globals.Cfg.Citation.CSLFile != "" {
			err := globals.Citeproc.SetCslFile(globals.Cfg.Citation.CSLFile)
			if err != nil {
				cli.Error(err, true)
			}
		}

		if globals.Cfg.Citation.LocaleFile != "" {
			err := globals.Citeproc.SetLocaleFile(globals.Cfg.Citation.LocaleFile)
			if err != nil {
				cli.Error(err, true)
			}
		}

		err := globals.Citeproc.Init()
		if err != nil {
			cli.Error(err, false)
			cli.Info("Citations will be turned off\n")
			globals.Cfg.Citation.Enabled = false
		}

		bibFile, err := os.ReadFile(globals.Cfg.Citation.File)
		if err == nil {
			citsList := make([]csljson.Item, 0)
			err := json.Unmarshal(bibFile, &citsList)
			if err != nil {
				cli.Error(err, true)
			}
			for _, cit := range citsList {
				//globals.Citations[cit.CitationKey] = cit
				globals.Citations[cit.ID] = cit
			}
			err = globals.Citeproc.AddCitation(citsList...)
			if err != nil {
				cli.Error(err, false)
				cli.Info("Citations will be turned off\n")
				globals.Cfg.Citation.Enabled = false
			}
		} else {
			cli.Error(fmt.Errorf("could not read bibliography file"), false)
			cli.Info("Citations will be turned off\n")
			globals.Cfg.Citation.Enabled = false
		}
		citT := time.Now()
		cli.Output("CSLJSON loaded in %v\n\n", citT.Sub(parsed))
	}

	pp := pdf.FromAst(ast)

	outName := globals.NameEncode(globals.Cfg.Paper.Title) + ".pdf"
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

	cli.Warning("Please look over the PDF yourself and make sure that it is correct.\nmdpaper is not responsible for the correctness of the output.\n")
}
