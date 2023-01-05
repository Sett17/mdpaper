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
	_, _ = cfmt.Printf("{{%s}}::lightBlue", fmt.Sprintf(format, a...))
}

func Output(format string, a ...any) {
	printPrefix()
	_, _ = cfmt.Printf("{{%s}}::white", fmt.Sprintf(format, a...))
}

func Other(format string, a ...any) {
	printPrefix()
	_, _ = cfmt.Printf("{{%s}}::gray", fmt.Sprintf(format, a...))
}

func Warning(format string, a ...any) {
	printPrefix()
	_, _ = cfmt.Printf("{{%s}}::bgYellow|black", fmt.Sprintf(format, a...))
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
