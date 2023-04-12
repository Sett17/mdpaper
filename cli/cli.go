package cli

import (
	"encoding/json"
	"fmt"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sett17/mdpaper/v2/globals"
	"io/ioutil"
	"net/http"
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

type Release struct {
	TagName string `json:"tag_name"`
}

func CheckVersion() {
	url := "https://api.github.com/repos/Sett17/mdpaper/releases/latest"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[CheckVersion] Error sending HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[CheckVersion] Error reading HTTP response body: %v\n", err)
		return
	}

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		fmt.Printf("[CheckVersion] Error parsing JSON response: %v\n", err)
		return
	}

	// Remove leading 'v' character from GitHub API response tag name
	onlineVersion := strings.TrimPrefix(release.TagName, "v")

	if onlineVersion != globals.Version {
		fmt.Printf("A new version (%s) of mdpaper is available! Run \ngo install github.com/Sett17/mdpaper/v2@{%s}\n to update.\n", release.TagName, release.TagName)
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
