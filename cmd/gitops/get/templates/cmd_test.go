package templates_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
	"github.com/weaveworks/weave-gitops/cmd/gitops/get/templates"
)

func TestEndpointNotSet(t *testing.T) {
	cmd := templates.TemplateCommand(&config.Options{})
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.EqualError(t, err, "the Weave GitOps Enterprise HTTP API endpoint flag (--endpoint) has not been set")
}

func TestProviderIsNotValid(t *testing.T) {
	cmd := templates.TemplateCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	})
	cmd.SetArgs([]string{
		"--provider",
		"--endpoint", "http://localhost:8000",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "provider \"--endpoint\" is not valid")
}

func TestTemplateNameIsRequired(t *testing.T) {
	cmd := templates.TemplateCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	})
	cmd.SetArgs([]string{
		"--list-parameters",
		"--provider", "aws",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "template name is required")
}
