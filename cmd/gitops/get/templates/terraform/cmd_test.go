package terraform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
	"github.com/weaveworks/weave-gitops/cmd/gitops/get/templates/terraform"
)

func TestEndpointNotSet(t *testing.T) {
	cmd := terraform.TerraformCommand(&config.Options{})
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.EqualError(t, err, "the Weave GitOps Enterprise HTTP API endpoint flag (--endpoint) has not been set")
}

func TestTemplateNameIsRequired(t *testing.T) {
	cmd := terraform.TerraformCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	})
	cmd.SetArgs([]string{
		"--list-parameters",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "terraform template name is required")
}
