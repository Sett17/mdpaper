package cli

import (
	"github.com/i582/cfmt/cmd/cfmt"
	"strings"
)

func Parse(args []string) (mdFiles []string) {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
			mdFiles = append(mdFiles, arg)
		} else {
			arg = strings.TrimLeft(arg, "-")
			for _, pArg := range ProgArgs {
				if pArg.Short == arg || pArg.Long == arg {
					defer pArg.Func(args[i+1])
					i++
				}
			}
		}
	}
	return
}

func ParseForHelp(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			arg = strings.TrimLeft(arg, "-")
			if arg == HelpProgArg.Short || arg == HelpProgArg.Long {
				HelpProgArg.Func("")
				return
			}
		}
	}
}

func ParseForVersion(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			arg = strings.TrimLeft(arg, "-")
			if arg == VersionProgArg.Short || arg == VersionProgArg.Long {
				VersionProgArg.Func("")
				return
			}
		}
	}
}

var Logo = cfmt.Sprint(`

    ███╗   ███╗██████╗
    ████╗ ████║██╔══██╗
    ██╔████╔██║██║  ██║
    ██║╚██╔╝██║██║  ██║
    ██║ ╚═╝ ██║██████╔╝
    ╚═╝     ╚═╝╚═════╝
  ██████╗  █████╗ ██████╗ ███████╗██████╗ 
  ██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗
  ██████╔╝███████║██████╔╝█████╗  ██████╔╝
  ██╔═══╝ ██╔══██║██╔═══╝ ██╔══╝  ██╔══██╗
  ██║     ██║  ██║██║     ███████╗██║  ██║
  ╚═╝     ╚═╝  ╚═╝╚═╝     ╚══════╝╚═╝  ╚═╝
by Sett|{{https://github.com/Sett17/mdpaper}}::#BA8EF7
`)
