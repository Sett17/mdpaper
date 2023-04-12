package conversions

import (
	"fmt"
	"github.com/sett17/mdpaper/v2/cli"
	"github.com/sett17/mdpaper/v2/globals"
	"github.com/sett17/mdpaper/v2/goldmark-math"
	"github.com/sett17/mdpaper/v2/pdf/conversions/options"
	"github.com/sett17/mdpaper/v2/pdf/spec"
	"os"
	"os/exec"
)

func Math(m *goldmark_math.MathBlock) (retO *spec.XObject, retA *spec.Addable) {
	optionString := options.Extract(m.Options)
	opts, err := options.Parse(optionString)
	if err != nil {
		cli.Error(fmt.Errorf("error parsing mermaid options: %w", err), false)
		cli.Warning(optionString)
	}

	inputFile, err := os.CreateTemp("", "mdpapermath")
	if err != nil {
		cli.Error(err, false)
		return nil, nil
	}

	_, err = inputFile.WriteString(`\documentclass{article}
\usepackage{amsmath}
\pagenumbering{gobble}
\thispagestyle{empty}
\begin{document}
\begin{align*}
`)
	if err != nil {
		cli.Error(err, false)
		return nil, nil
	}
	_, err = inputFile.Write(m.Text(globals.File))
	if err != nil {
		cli.Error(err, false)
		return nil, nil
	}
	if m.Text(globals.File)[len(m.Text(globals.File))-1] != '\n' {
		_, err = inputFile.WriteString(`\n`)
	}
	_, err = inputFile.WriteString(`\end{align*}
\end{document}`)
	if err != nil {
		cli.Error(err, false)
		return nil, nil
	}

	inputFile.Close()

	latexCommand := exec.Command("latex", "-output-format=dvi", "-interaction=nonstopmode", fmt.Sprintf("%s", inputFile.Name()))
	latexCommand.Dir = os.TempDir()
	latexOutput, err := latexCommand.CombinedOutput()
	if err != nil {
		cli.Warning("There was an error running latex, output may not be as expected\n")
		cli.Info(string(latexOutput))
	}

	_, err = os.Stat(inputFile.Name() + ".dvi")
	if err != nil {
		cli.Error(fmt.Errorf("latex did not produce a dvi file"), false)
		return nil, nil
	}

	dvipngCommand := exec.Command("dvipng", "-D", "1000", "-T", "tight", "-o", fmt.Sprintf("%s.png", inputFile.Name()), fmt.Sprintf("%s.dvi", inputFile.Name()))
	dvipngCommand.Dir = os.TempDir()
	dvipngOutput, err := dvipngCommand.CombinedOutput()
	if err != nil {
		cli.Error(err, false)
		cli.Info(string(dvipngOutput))
		return nil, nil
	}

	w := .65
	if i, ok := opts.GetFloat("width"); ok {
		w = i
	}
	io, ia := spec.NewImageObjectFromFile(inputFile.Name()+".png", w)

	retO = &io
	retA = &ia
	//defer os.Remove(inputFile.Name())
	defer os.Remove(inputFile.Name() + ".png")
	defer os.Remove(inputFile.Name() + ".dvi")
	defer os.Remove(inputFile.Name() + ".aux")
	defer os.Remove(inputFile.Name() + ".log")

	return
}
