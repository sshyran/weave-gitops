package clusters_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
	"github.com/weaveworks/weave-gitops/cmd/gitops/get/clusters"
	"github.com/weaveworks/weave-gitops/pkg/adapters"
)

func TestGetCluster(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		status    int
		response  interface{}
		options   *config.Options
		args      []string
		result    string
		errString string
	}{
		{
			name:     "cluster kubeconfig",
			url:      "http://localhost:8000/v1/clusters/dev-cluster/kubeconfig",
			status:   http.StatusOK,
			response: httpmock.File("../../../../pkg/adapters/testdata/cluster_kubeconfig.json"),
			options: &config.Options{
				Endpoint: "http://localhost:8000",
			},
			args: []string{
				"dev-cluster",
				"--print-kubeconfig",
			},
		},
		{
			name: "http error",
			options: &config.Options{
				Endpoint: "not_a_valid_url",
			},
			args: []string{
				"dev-cluster",
				"--print-kubeconfig",
			},
			errString: "failed to parse endpoint: parse \"not_a_valid_url\": invalid URI for request",
		},
		{
			name:    "no endpoint",
			options: &config.Options{},
			args: []string{
				"dev-cluster",
				"--print-kubeconfig",
			},
			errString: "the Weave GitOps Enterprise HTTP API endpoint flag (--endpoint) has not been set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := adapters.NewHTTPClient()
			httpmock.ActivateNonDefault(client.GetBaseClient())
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder(
				http.MethodGet,
				tt.url,
				func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(tt.status, tt.response)
				},
			)

			cmd := clusters.ClusterCommand(tt.options, client)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			if tt.errString == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.errString)
			}
		})
	}
}
