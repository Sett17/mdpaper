package cli

import (
	"crypto/md5"
	"fmt"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sett17/mdpaper/globals"
	"gopkg.in/yaml.v3"
	"os"
)

type progArgFunction func(string)

type ProgArg struct {
	Name  string
	Short string
	Long  string
	Help  string

	Func progArgFunction
}

var HelpProgArg = ProgArg{
	Name:  "Help",
	Short: "h",
	Long:  "help",
	Func: func(_ string) {
		cfmt.Printf(`
{{Usage:}}::underline mdpaper [OPTIONS] MARKDOWN_FILE
{{If multiple files are specified, only the last one will be used}}::gray`)
		fmt.Println(Logo)
		for _, arg := range ProgArgs {
			cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", arg.Short, arg.Long, arg.Help)
		}
		cfmt.Printf("\n  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", "h", "help", "Show this help message")
		cfmt.Printf("  {{-%s}}::purple, {{--%s}}::purple\t{{%s}}::gray\n", VersionProgArg.Short, VersionProgArg.Long, VersionProgArg.Help)
		os.Exit(0)
	},
}

var VersionProgArg = ProgArg{
	Name:  "Version",
	Short: "v",
	Long:  "version",
	Help:  "Show the version of this program",
	Func: func(_ string) {
		cfmt.Printf("mdpaper version %s\n", globals.Version)
		os.Exit(0)
	},
}

var ProgArgs = []ProgArg{
	{
		Name:  "Config",
		Short: "c",
		Long:  "config",
		Help:  "Sets the path to the config file",
		Func: func(path string) {
			CfgFunc(path)
		},
	},
}

func CfgFunc(path string) {
	cfgFile, err := os.ReadFile(path)
	if err == nil {
		err = yaml.Unmarshal(cfgFile, &globals.Cfg)
		if err != nil {
			panic(err)
		}
		Output("Loaded config from %s\n", path)
	}
	out, err := yaml.Marshal(globals.Cfg)
	if err != nil {
		panic(err)
	}
	if md5.Sum(out) != md5.Sum(cfgFile) {
		err = os.WriteFile(path, out, 0644)
		if err != nil {
			panic(err)
		}
		Info("Updated config file %s\n", path)
	}
	globals.DidConfig = true

}
