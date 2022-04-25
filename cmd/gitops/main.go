package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/weaveworks/weave-gitops/cmd/gitops/root"
	"github.com/weaveworks/weave-gitops/cmd/internal"
)

func main() {
	rootCmd := root.RootCmd(resty.New())
	cmds := rootCmd.Commands()
	cmdFound := false
	backing := internal.BackingCLI{CmdName: "flux", Alias: "gitops", VersionPath: "version", Version: "v0.27.3"}

	for _, a := range cmds {
		for _, b := range os.Args[1:] {
			if a.Name() == b {
				cmdFound = true
				break
			}
		}
	}

	if cmdFound {
		if err := rootCmd.Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		if _, err := backing.Invoke(strings.Join(os.Args[1:], " "), false); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}
