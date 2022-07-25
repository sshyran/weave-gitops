package credentials

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/weaveworks/weave-gitops/cmd/gitops/cmderrors"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
	"github.com/weaveworks/weave-gitops/pkg/adapters"
	"github.com/weaveworks/weave-gitops/pkg/templates"
	"k8s.io/cli-runtime/pkg/printers"
)

func CredentialCommand(opts *config.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "credential",
		Aliases: []string{"credentials"},
		Short:   "Get CAPI credentials",
		Example: `
# Get all CAPI credentials
gitops get credentials
		`,
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE:       getCredentialCmdPreRunE(&opts.Endpoint),
		RunE:          getCredentialCmdRunE(opts),
	}

	return cmd
}

func getCredentialCmdPreRunE(endpoint *string) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, s []string) error {
		if *endpoint == "" {
			return cmderrors.ErrNoWGEEndpoint
		}

		return nil
	}
}

func getCredentialCmdRunE(opts *config.Options) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		client := adapters.NewHTTPClient().EnableCLIAuth()
		err := client.ConfigureClientWithOptions(opts, os.Stdout)
		if err != nil {
			return err
		}

		w := printers.GetNewTabWriter(os.Stdout)

		defer w.Flush()

		return templates.GetCredentials(client, w)
	}
}
