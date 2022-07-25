package credentials_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weaveworks/weave-gitops/cmd/gitops/root"
)

func TestEndpointNotSet(t *testing.T) {
	cmd := root.RootCmd()
	cmd.SetArgs([]string{
		"get", "credentials",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "the Weave GitOps Enterprise HTTP API endpoint flag (--endpoint) has not been set")
}
