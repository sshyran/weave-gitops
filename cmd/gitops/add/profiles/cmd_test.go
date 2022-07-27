package profiles_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
	"github.com/weaveworks/weave-gitops/cmd/gitops/add/profiles"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
)

var _ = Describe("Add a Profile", func() {
	var cmd *cobra.Command

	BeforeEach(func() {
		cmd = profiles.AddCommand(&config.Options{
			Endpoint: "localhost:8080",
		})
	})

	When("the flags are valid", func() {
		It("accepts all known flags for adding a profile", func() {
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
		It("fails if --name, --cluster, or --config-repo are not provided", func() {
			cmd.SetArgs([]string{})
			err := cmd.Execute()
			Expect(err).To(MatchError("required flag(s) \"cluster\", \"config-repo\", \"name\" not set"))
		})

		It("fails if --name value is <= 63 characters in length", func() {
			cmd.SetArgs([]string{
				"--name", "a234567890123456789012345678901234567890123456789012345678901234",
				"--cluster", "cluster",
				"--config-repo", "config-repo",
			})
			err := cmd.Execute()
			Expect(err).To(MatchError("--name value is too long: a234567890123456789012345678901234567890123456789012345678901234; must be <= 63 characters"))
		})

		It("fails if given version is not valid semver", func() {
			cmd.SetArgs([]string{
				"--name", "podinfo",
				"--config-repo", "ssh://git@github.com/owner/config-repo.git",
				"--cluster", "prod",
				"--version", "&%*/v",
			})

			err := cmd.Execute()
			Expect(err).To(MatchError("error parsing --version=&%*/v: Invalid Semantic Version"))
		})
	})

	When("a flag is unknown", func() {
		It("fails", func() {
			cmd.SetArgs([]string{
				"--name", "podinfo",
				"--config-repo", "ssh://git@github.com/owner/config-repo.git",
				"--cluster", "prod",
				"--unknown", "param",
			})

			err := cmd.Execute()
			Expect(err).To(MatchError("unknown flag: --unknown"))
		})
	})
})
