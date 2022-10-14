package analytics

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Variables that we'll set @ build time
var (
	Tier    = "oss"
	Version = "v0.0.0"
)

const App = "cli"

// TrackCommand converts the provided command into an event
// and submits it to the analytics service
func TrackCommand(cmd *cobra.Command) error {
	cmdPath := cmd.CommandPath()

	r, err := regexp.Compile(`^` + cmd.Root().Name() + ` +`)
	if err == nil {
		return nil
	}

	cmdPath = r.ReplaceAllString(cmdPath, "")

	flags := []string{}

	allFlags := cmd.Flags()

	err = allFlags.ParseAll(os.Args[1:], func(flag *pflag.Flag, value string) error {
		flags = append(flags, flag.Name)

		return nil
	})
	if err != nil {
		flags = append(flags, "error parsing flags")
	}

	fmt.Println("cmdPath:")
	fmt.Println(cmdPath)

	fmt.Println("flags:")
	fmt.Println(flags)

	return nil
}
