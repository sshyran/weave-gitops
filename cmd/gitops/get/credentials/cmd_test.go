package credentials_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
	"github.com/weaveworks/weave-gitops/cmd/gitops/get/credentials"
)

func TestEndpointNotSet(t *testing.T) {
	cmd := credentials.CredentialCommand(&config.Options{})
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.EqualError(t, err, "the Weave GitOps Enterprise HTTP API endpoint flag (--endpoint) has not been set")
}
