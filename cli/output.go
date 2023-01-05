package cli

import (
	"fmt"
	"github.com/i582/cfmt/cmd/cfmt"
	"os"
	"strings"
)

func printPrefix() {
	//cfmt.Print("{{mdpaper:}}::purple ")
}

func Info(format string, a ...any) {
	printPrefix()
	linebreak := ""
	if endsWidthNewline(format) {
		format = format[:len(format)-1]
		linebreak = "\n"
	}
	_, _ = cfmt.Printf("{{%s}}::lightBlue%s", fmt.Sprintf(format, a...), linebreak)
}

func Output(format string, a ...any) {
	printPrefix()
	linebreak := ""
	if endsWidthNewline(format) {
		format = format[:len(format)-1]
		linebreak = "\n"
	}
	_, _ = cfmt.Printf("{{%s}}::white%s", fmt.Sprintf(format, a...), linebreak)
}

func Other(format string, a ...any) {
	printPrefix()
	linebreak := ""
	if endsWidthNewline(format) {
		format = format[:len(format)-1]
		linebreak = "\n"
	}
	_, _ = cfmt.Printf("{{%s}}::gray%s", fmt.Sprintf(format, a...), linebreak)
}

func Warning(format string, a ...any) {
	printPrefix()
	linebreak := ""
	if endsWidthNewline(format) {
		format = format[:len(format)-1]
		linebreak = "\n"
	}
	_, _ = cfmt.Printf("{{%s}}::bgYellow|black%s", fmt.Sprintf(format, a...), linebreak)
}

func Error(err error, exit bool) {
	_, _ = cfmt.Printf("{{ERROR:}}::bgRed\n{{%s}}::red\n", err.Error())
	if exit {
		os.Exit(1)
	}
}

func Separator() {
	_, _ = cfmt.Printf("{{%s}}::gray\n", strings.Repeat("─", 80))
}

func endsWidthNewline(s string) bool {
	return s[len(s)-1] == 0x0a
}
