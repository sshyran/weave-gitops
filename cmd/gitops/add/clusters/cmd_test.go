package clusters_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"github.com/weaveworks/weave-gitops/cmd/gitops/add/clusters"
	"github.com/weaveworks/weave-gitops/cmd/gitops/cmderrors"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
	"github.com/weaveworks/weave-gitops/cmd/gitops/root"
	"github.com/weaveworks/weave-gitops/pkg/adapters"
	"github.com/weaveworks/weave-gitops/pkg/testutils"
)

func TestSetSeparateValues(t *testing.T) {
	client := adapters.NewHTTPClient()

	httpmock.ActivateNonDefault(client.GetBaseClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		http.MethodPost,
		"http://localhost:8000/v1/templates/cluster-template-eks-fargate/render?template_kind=CAPITemplate",
		func(r *http.Request) (*http.Response, error) {
			var vs adapters.TemplateParameterValuesAndCredentials

			err := json.NewDecoder(r.Body).Decode(&vs)
			assert.NoError(t, err)

			assert.Equal(t, "dev", vs.Values["CLUSTER_NAME"])
			assert.Equal(t, "us-east-1", vs.Values["AWS_REGION"])
			assert.Equal(t, "ssh_key", vs.Values["AWS_SSH_KEY_NAME"])
			assert.Equal(t, "1.19", vs.Values["KUBERNETES_VERSION"])

			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("../../../../pkg/adapters/testdata/rendered_template_capi.json"))
		},
	)

	cmd := clusters.ClusterCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	}, client)
	cmd.SetArgs([]string{
		"--from-template=cluster-template-eks-fargate",
		"--set=CLUSTER_NAME=dev",
		"--set=AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
		"--dry-run",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestSetMultipleValues(t *testing.T) {
	client := adapters.NewHTTPClient()

	httpmock.ActivateNonDefault(client.GetBaseClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		http.MethodPost,
		"http://localhost:8000/v1/templates/cluster-template-eks-fargate/render?template_kind=CAPITemplate",
		func(r *http.Request) (*http.Response, error) {
			var vs adapters.TemplateParameterValuesAndCredentials

			err := json.NewDecoder(r.Body).Decode(&vs)
			assert.NoError(t, err)

			assert.Equal(t, "dev", vs.Values["CLUSTER_NAME"])
			assert.Equal(t, "us-east-1", vs.Values["AWS_REGION"])
			assert.Equal(t, "ssh_key", vs.Values["AWS_SSH_KEY_NAME"])
			assert.Equal(t, "1.19", vs.Values["KUBERNETES_VERSION"])

			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("../../../../pkg/adapters/testdata/rendered_template_capi.json"))
		},
	)

	cmd := clusters.ClusterCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	}, client)
	cmd.SetArgs([]string{
		"--from-template=cluster-template-eks-fargate",
		"--set=CLUSTER_NAME=dev,AWS_REGION=us-east-1,AWS_SSH_KEY_NAME=ssh_key,KUBERNETES_VERSION=1.19",
		"--dry-run",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestSetMultipleAndSeparateValues(t *testing.T) {
	client := adapters.NewHTTPClient()

	httpmock.ActivateNonDefault(client.GetBaseClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		http.MethodPost,
		"http://localhost:8000/v1/templates/cluster-template-eks-fargate/render?template_kind=CAPITemplate",
		func(r *http.Request) (*http.Response, error) {
			var vs adapters.TemplateParameterValuesAndCredentials

			err := json.NewDecoder(r.Body).Decode(&vs)
			assert.NoError(t, err)

			assert.Equal(t, "dev", vs.Values["CLUSTER_NAME"])
			assert.Equal(t, "us-east-1", vs.Values["AWS_REGION"])
			assert.Equal(t, "ssh_key", vs.Values["AWS_SSH_KEY_NAME"])
			assert.Equal(t, "1.19", vs.Values["KUBERNETES_VERSION"])

			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("../../../../pkg/adapters/testdata/rendered_template_capi.json"))
		},
	)

	cmd := clusters.ClusterCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	}, client)
	cmd.SetArgs([]string{
		"--from-template=cluster-template-eks-fargate",
		"--set=CLUSTER_NAME=dev,AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
		"--dry-run",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestEndpointNotSet(t *testing.T) {
	client := adapters.NewHTTPClient()
	cmd := clusters.ClusterCommand(&config.Options{}, client)
	cmd.SetArgs([]string{
		"--from-template=cluster-template-eks-fargate",
		"--set=CLUSTER_NAME=dev",
		"--set=AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
		"--dry-run",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "the Weave GitOps Enterprise HTTP API endpoint flag (--endpoint) has not been set")
}

func TestGitProviderToken(t *testing.T) {
	t.Cleanup(testutils.Setenv("GITHUB_TOKEN", "test-token"))

	client := adapters.NewHTTPClient()

	httpmock.ActivateNonDefault(client.GetBaseClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodPost,
		"http://localhost:8000/v1/clusters",
		func(r *http.Request) (*http.Response, error) {
			h, ok := r.Header["Git-Provider-Token"]
			assert.True(t, ok)
			assert.Contains(t, h, "test-token")

			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("../../../../pkg/adapters/testdata/pull_request_created.json"))
		},
	)

	cmd := clusters.ClusterCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	}, client)
	cmd.SetArgs([]string{
		"--from-template=cluster-template-eks-fargate",
		"--url=https://github.com/weaveworks/test-repo",
		"--set=CLUSTER_NAME=dev",
		"--set=AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestGitProviderToken_NoURL(t *testing.T) {
	cmd := root.RootCmd()
	cmd.SetArgs([]string{
		"add", "cluster",
		"--from-template=cluster-template-eks-fargate",
		"--set=CLUSTER_NAME=dev",
		"--set=AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
		"--endpoint", "http://localhost:8000",
	})

	err := cmd.Execute()
	assert.ErrorIs(t, err, cmderrors.ErrNoURL)
}

func TestGitProviderToken_InvalidURL(t *testing.T) {
	cmd := root.RootCmd()
	cmd.SetArgs([]string{
		"add", "cluster",
		"--from-template=cluster-template-eks-fargate",
		"--url=invalid_url",
		"--set=CLUSTER_NAME=dev",
		"--set=AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
		"--endpoint", "http://localhost:8000",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "cannot parse url: could not get provider name from URL invalid_url: no git providers found for \"invalid_url\"")
}

func TestParseProfiles_ValidRequest(t *testing.T) {
	t.Cleanup(testutils.Setenv("GITHUB_TOKEN", "test-token"))

	client := adapters.NewHTTPClient()

	httpmock.ActivateNonDefault(client.GetBaseClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodPost,
		"http://localhost:8000/v1/clusters",
		func(r *http.Request) (*http.Response, error) {
			h, ok := r.Header["Git-Provider-Token"]
			assert.True(t, ok)
			assert.Contains(t, h, "test-token")

			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("../../../../pkg/adapters/testdata/pull_request_created.json"))
		},
	)

	cmd := clusters.ClusterCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	}, client)
	cmd.SetArgs([]string{
		"--from-template=cluster-template-eks-fargate",
		"--url=https://github.com/weaveworks/test-repo",
		"--set=CLUSTER_NAME=dev",
		"--set=AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
		"--profile=name=foo-profile,version=0.0.1",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestParseProfiles_InvalidKey(t *testing.T) {
	t.Cleanup(testutils.Setenv("GITHUB_TOKEN", "test-token"))

	client := adapters.NewHTTPClient()

	httpmock.ActivateNonDefault(client.GetBaseClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodPost,
		"http://localhost:8000/v1/clusters",
		func(r *http.Request) (*http.Response, error) {
			h, ok := r.Header["Git-Provider-Token"]
			assert.True(t, ok)
			assert.Contains(t, h, "test-token")

			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("../../../../pkg/adapters/testdata/pull_request_created.json"))
		},
	)

	cmd := clusters.ClusterCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	}, client)
	cmd.SetArgs([]string{
		"--from-template=cluster-template-eks-fargate",
		"--url=https://github.com/weaveworks/test-repo",
		"--set=CLUSTER_NAME=dev",
		"--set=AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
		"--profile=test=foo-profile",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "error parsing profiles: invalid key: test")
}

func TestParseProfiles_InvalidValue(t *testing.T) {
	t.Cleanup(testutils.Setenv("GITHUB_TOKEN", "test-token"))

	client := adapters.NewHTTPClient()

	httpmock.ActivateNonDefault(client.GetBaseClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodPost,
		"http://localhost:8000/v1/clusters",
		func(r *http.Request) (*http.Response, error) {
			h, ok := r.Header["Git-Provider-Token"]
			assert.True(t, ok)
			assert.Contains(t, h, "test-token")

			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("../../../../pkg/adapters/testdata/pull_request_created.json"))
		},
	)

	cmd := clusters.ClusterCommand(&config.Options{
		Endpoint: "http://localhost:8000",
	}, client)
	cmd.SetArgs([]string{
		"add", "cluster",
		"--from-template=cluster-template-eks-fargate",
		"--url=https://github.com/weaveworks/test-repo",
		"--set=CLUSTER_NAME=dev",
		"--set=AWS_REGION=us-east-1",
		"--set=AWS_SSH_KEY_NAME=ssh_key",
		"--set=KUBERNETES_VERSION=1.19",
		"--profile=name=foo;profile",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "error parsing profiles: invalid value for name: foo;profile")
}

func TestInsecureSkipVerifyFalse(t *testing.T) {
	client := adapters.NewHTTPClient()

	cmd := clusters.ClusterCommand(&config.Options{
		InsecureSkipTLSVerify: true,
	}, client)

	// Command is incomplete and should raise an error, it helps us short circuit here to quickly
	// test that the client has been set
	err := cmd.Execute()
	assert.Error(t, err)

	transport, ok := client.GetBaseClient().Transport.(*http.Transport)
	assert.True(t, ok)
	// Its set to nil and uses whatever the golang defaults are (InsecureSkipVerify: false)
	assert.Nil(t, transport.TLSClientConfig)
}
