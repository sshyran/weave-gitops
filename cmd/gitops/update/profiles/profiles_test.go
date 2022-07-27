package profiles_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
	"github.com/weaveworks/weave-gitops/cmd/gitops/update/profiles"
)

var _ = Describe("Update Profile(s)", func() {
	var cmd *cobra.Command

	BeforeEach(func() {
		cmd = profiles.UpdateCommand(&config.Options{
			Endpoint: "localhost:8080",
		})
	})

	When("the flags are valid", func() {
		It("accepts all known flags for updating a profile", func() {
			cmd.SetArgs([]string{
				"--name", "podinfo",
				"--version", "0.0.1",
				"--cluster", "prod",
				"--config-repo", "https://ssh@github:test/test.git",
				"--auto-merge", "true",
			})

			err := cmd.Execute()
			Expect(err.Error()).NotTo(ContainSubstring("unknown flag"))
		})
	})

	When("flags are not valid", func() {
		It("fails if --name, --cluster, --version or --config-repo are not provided", func() {
			cmd.SetArgs([]string{})
			err := cmd.Execute()
			Expect(err).To(MatchError("required flag(s) \"cluster\", \"config-repo\", \"name\", \"version\" not set"))
		})

		It("fails if given version is not valid semver", func() {
			cmd.SetArgs([]string{
				"--name", "podinfo",
				"--config-repo", "ssh://git@github.com/owner/config-repo.git",
				"--cluster", "prod",
				"--version", "&%*/v",
			})

			err := cmd.Execute()
			Expect(err).To(MatchError(ContainSubstring("error parsing --version=&%*/v")))
		})
	})

	When("a flag is unknown", func() {
		It("fails", func() {
			cmd.SetArgs([]string{
				"--unknown", "param",
			})

			err := cmd.Execute()
			Expect(err).To(MatchError("unknown flag: --unknown"))
		})
	})
})
