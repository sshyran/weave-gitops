package bootstrap

import (
	"fmt"

	"github.com/weaveworks/weave-gitops/cmd/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	zeta bool
)

var Cmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "install gitops and flux",
	Example: `
# Install gitops and flux
gitops bootstrap
`,
}

func init() {
	Cmd.Flags().BoolVarP(&zeta, "zeta", "z", true, "perform only the pre-installation checks")

	internal.WrapCommand(
		Cmd,
		internal.BackingCLI{CmdName: "flux", Alias: "gitops", VersionPath: "version", Version: "v0.27.3"},
		internal.Wrapper{
			Before: func(fs *pflag.FlagSet, args []string) error {
				fmt.Printf("BEFORE...\n")
				return nil
			},
			After: func(fs *pflag.FlagSet, args []string, wrappedOutput []byte) error {
				fmt.Printf("AFTER...\n%s\n", wrappedOutput)
				return nil
			},
		},
		[]string{"endpoint", "git-host-types", "insecure-skip-tls-verify", "override-in-cluster", "zeta"})
}
