package automation

import (
	"context"
	"crypto/md5"
	"fmt"
	"strings"

	//	"github.com/go-git/go-billy/v5/memfs"
	//	gogit "github.com/go-git/go-git/v5"
	//	"github.com/go-git/go-git/v5/config"
	//	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	wego "github.com/weaveworks/weave-gitops/api/v1alpha1"
	//	"github.com/weaveworks/weave-gitops/pkg/git"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
	"github.com/weaveworks/weave-gitops/pkg/models"
	//	"github.com/weaveworks/weave-gitops/pkg/testutils"
	"sigs.k8s.io/yaml"
)

var (
	ctx context.Context

	emptyRepoURL = gitproviders.RepoURL{}
)

func createRepoURL(url string) gitproviders.RepoURL {
	repoURL, err := gitproviders.NewRepoURL(url)
	Expect(err).NotTo(HaveOccurred())
	return repoURL
}

var _ = Describe("Generate manifests", func() {
	var app models.Application

	var _ = BeforeEach(func() {
		app = models.Application{
			Name:                "bar",
			Namespace:           wego.DefaultNamespace,
			HelmSourceURL:       "",
			GitSourceURL:        createRepoURL("ssh://git@github.com/foo/bar.git"),
			Branch:              "main",
			Path:                "./kustomize",
			AutomationType:      models.AutomationTypeKustomize,
			SourceType:          models.SourceTypeGit,
			HelmTargetNamespace: "",
		}

		gitProviders.GetDefaultBranchReturns("main", nil)

		ctx = context.Background()
	})

	Context("add app with config in app repo", func() {
		BeforeEach(func() {
			app.GitSourceURL = createRepoURL("ssh://git@github.com/foo/bar.git")
			app.ConfigURL = emptyRepoURL
		})

		Describe("generates source manifest", func() {
			It("creates GitRepository when source type is git", func() {
				app.SourceType = models.SourceTypeGit
				_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
				Expect(err).ShouldNot(HaveOccurred())

				Expect(fluxClient.CreateSourceGitCallCount()).To(Equal(1))

				name, url, branch, secretRef, namespace := fluxClient.CreateSourceGitArgsForCall(0)
				Expect(name).To(Equal("bar"))
				Expect(url.String()).To(Equal("ssh://git@github.com/foo/bar.git"))
				Expect(branch).To(Equal("main"))
				Expect(secretRef).To(Equal("wego-test-cluster-bar"))
				Expect(namespace).To(Equal(wego.DefaultNamespace))
			})

			It("creates HelmRepository when source type is helm", func() {
				app.HelmSourceURL = "https://charts.kube-ops.io"
				app.Path = "loki"
				app.Name = "loki"
				app.SourceType = models.SourceTypeHelm
				app.ConfigURL = createRepoURL("ssh://git@github.com/owner/config-repo.git")

				_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
				Expect(err).ShouldNot(HaveOccurred())

				Expect(fluxClient.CreateSourceHelmCallCount()).To(Equal(1))

				name, url, namespace := fluxClient.CreateSourceHelmArgsForCall(0)
				Expect(name).To(Equal("loki"))
				Expect(url).To(Equal("https://charts.kube-ops.io"))
				Expect(namespace).To(Equal(wego.DefaultNamespace))
			})
		})

		Describe("generates application goat", func() {
			It("creates a kustomization if deployment type kustomize", func() {
				_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
				Expect(err).ShouldNot(HaveOccurred())

				Expect(fluxClient.CreateKustomizationCallCount()).To(Equal(1))

				name, source, path, namespace := fluxClient.CreateKustomizationArgsForCall(0)
				Expect(name).To(Equal("bar"))
				Expect(source).To(Equal("bar"))
				Expect(path).To(Equal("./kustomize"))
				Expect(namespace).To(Equal(wego.DefaultNamespace))
			})

			It("creates a helm release using a git source if source type is git", func() {
				app.AutomationType = models.AutomationTypeHelm
				app.Path = "loki"
				app.Name = "bar"
				app.ConfigURL = createRepoURL("ssh://github.com/owner/repo")

				app.Path = "./charts/my-chart"
				app.AutomationType = models.AutomationTypeHelm

				_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
				Expect(err).ShouldNot(HaveOccurred())

				Expect(fluxClient.CreateHelmReleaseGitRepositoryCallCount()).To(Equal(1))

				name, source, path, namespace, targetNamespace := fluxClient.CreateHelmReleaseGitRepositoryArgsForCall(0)
				Expect(name).To(Equal("bar"))
				Expect(source).To(Equal("bar"))
				Expect(path).To(Equal("./charts/my-chart"))
				Expect(namespace).To(Equal(wego.DefaultNamespace))
				Expect(targetNamespace).To(Equal(""))
			})

			It("creates a helm release for git repository with target namespace if source type is git", func() {
				app.Path = "./charts/my-chart"
				app.AutomationType = models.AutomationTypeHelm
				app.HelmTargetNamespace = "sock-shop"

				_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
				Expect(err).ShouldNot(HaveOccurred())

				Expect(fluxClient.CreateHelmReleaseGitRepositoryCallCount()).To(Equal(1))

				name, source, path, namespace, targetNamespace := fluxClient.CreateHelmReleaseGitRepositoryArgsForCall(0)
				Expect(name).To(Equal("bar"))
				Expect(source).To(Equal("bar"))
				Expect(path).To(Equal("./charts/my-chart"))
				Expect(namespace).To(Equal(wego.DefaultNamespace))
				Expect(targetNamespace).To(Equal("sock-shop"))
			})
		})

		Context("add app with external config repo", func() {
			BeforeEach(func() {
				app.GitSourceURL = createRepoURL("https://github.com/user/repo")
				app.ConfigURL = createRepoURL("https://github.com/foo/bar")
			})

			Describe("generates source manifest", func() {
				It("creates GitRepository when source type is git", func() {
					_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
					Expect(err).ShouldNot(HaveOccurred())

					Expect(fluxClient.CreateSourceGitCallCount()).To(Equal(1))

					name, url, branch, secretRef, namespace := fluxClient.CreateSourceGitArgsForCall(0)
					Expect(name).To(Equal("bar"))
					Expect(url.String()).To(Equal("ssh://git@github.com/user/repo.git"))
					Expect(branch).To(Equal("main"))
					Expect(secretRef).To(Equal("wego-test-cluster-repo"))
					Expect(namespace).To(Equal(wego.DefaultNamespace))
				})

				It("creates HelmRepository when source type is helm", func() {
					app.HelmSourceURL = "https://charts.kube-ops.io"
					app.Path = "loki"
					app.Name = "loki"
					app.SourceType = models.SourceTypeHelm

					_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
					Expect(err).ShouldNot(HaveOccurred())

					Expect(fluxClient.CreateSourceHelmCallCount()).To(Equal(1))

					name, url, namespace := fluxClient.CreateSourceHelmArgsForCall(0)
					Expect(name).To(Equal("loki"))
					Expect(url).To(Equal("https://charts.kube-ops.io"))
					Expect(namespace).To(Equal(wego.DefaultNamespace))
				})
			})

			Describe("generateAppYaml", func() {
				FIt("generates the app yaml", func() {
					myAppModel := models.Application{
						Name:                "my-app",
						Namespace:           wego.DefaultNamespace,
						HelmSourceURL:       "",
						GitSourceURL:        createRepoURL("ssh://git@github.com/example/my-source"),
						Branch:              "main",
						Path:                "manifests",
						AutomationType:      models.AutomationTypeKustomize,
						SourceType:          models.SourceTypeGit,
						HelmTargetNamespace: "",
					}

					hash := "wego-" + getHash(myAppModel.GitSourceURL.String(), myAppModel.Path, myAppModel.Branch)
					myApp := AppToWegoApp(myAppModel)

					myApp.ObjectMeta.Labels = map[string]string{WeGOAppIdentifierLabelKey: hash}

					out, err := generateAppYaml(myAppModel)
					Expect(err).To(BeNil())

					result := wego.Application{}
					// Convert back to a struct to make the comparison more forgiving.
					// A straight string/byte comparison doesn't account for un-ordered keys in yaml.
					Expect(yaml.Unmarshal(out.Content, &result))

					diff := cmp.Diff(result, myApp)
					Expect(diff).To(Equal(""))

					// Not entirely sure how to get gomega to pretty-print the output of `diff`,
					// so we assert it should be empty above, and conditionally print the diff to make a nice assertion message.
					// `diff` is a formatted string
					if diff != "" {
						automationSvc.(*AutomationSvc).Logger.Println(diff)
					}
				})
			})

			Describe("generates application goat", func() {
				It("creates a kustomization if deployment type kustomize", func() {
					_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
					Expect(err).ShouldNot(HaveOccurred())

					Expect(fluxClient.CreateKustomizationCallCount()).To(Equal(1))

					name, source, path, namespace := fluxClient.CreateKustomizationArgsForCall(0)
					Expect(name).To(Equal("bar"))
					Expect(source).To(Equal("bar"))
					Expect(path).To(Equal("./kustomize"))
					Expect(namespace).To(Equal(wego.DefaultNamespace))
				})

				It("creates helm release using a helm repository if source type is helm", func() {
					app.HelmSourceURL = "https://charts.kube-ops.io"
					app.SourceType = models.SourceTypeHelm
					app.AutomationType = models.AutomationTypeHelm
					app.Path = "loki"
					app.Name = "loki"
					app.ConfigURL = createRepoURL("ssh://github.com/owner/repo")

					_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
					Expect(err).ShouldNot(HaveOccurred())

					Expect(fluxClient.CreateHelmReleaseHelmRepositoryCallCount()).To(Equal(1))

					name, chart, namespace, targetNamespace := fluxClient.CreateHelmReleaseHelmRepositoryArgsForCall(0)
					Expect(name).To(Equal("loki"))
					Expect(chart).To(Equal("loki"))
					Expect(namespace).To(Equal(wego.DefaultNamespace))
					Expect(targetNamespace).To(Equal(""))
				})

				It("creates a helm release using a git source if source type is git", func() {
					app.AutomationType = models.AutomationTypeHelm
					app.Path = "loki"
					app.Name = "bar"
					app.ConfigURL = createRepoURL("ssh://github.com/owner/repo")

					app.Path = "./charts/my-chart"
					app.AutomationType = models.AutomationTypeHelm

					_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
					Expect(err).ShouldNot(HaveOccurred())

					Expect(fluxClient.CreateHelmReleaseGitRepositoryCallCount()).To(Equal(1))

					name, source, path, namespace, targetNamespace := fluxClient.CreateHelmReleaseGitRepositoryArgsForCall(0)
					Expect(name).To(Equal("bar"))
					Expect(source).To(Equal("bar"))
					Expect(path).To(Equal("./charts/my-chart"))
					Expect(namespace).To(Equal(wego.DefaultNamespace))
					Expect(targetNamespace).To(Equal(""))
				})

				It("creates helm release for helm repository with target namespace if source type is helm", func() {
					app.HelmSourceURL = "https://charts.kube-ops.io"
					app.SourceType = models.SourceTypeHelm
					app.Path = "loki"
					app.Name = "loki"
					app.HelmTargetNamespace = "sock-shop"
					app.AutomationType = models.AutomationTypeHelm
					app.ConfigURL = createRepoURL("ssh://git@github.com/owner/config-repo.git")

					_, err := automationSvc.GenerateAutomation(ctx, app, "test-cluster")
					Expect(err).ShouldNot(HaveOccurred())

					Expect(fluxClient.CreateHelmReleaseHelmRepositoryCallCount()).To(Equal(1))

					name, chart, namespace, targetNamespace := fluxClient.CreateHelmReleaseHelmRepositoryArgsForCall(0)
					Expect(name).To(Equal("loki"))
					Expect(chart).To(Equal("loki"))
					Expect(namespace).To(Equal(wego.DefaultNamespace))
					Expect(targetNamespace).To(Equal("sock-shop"))
				})
			})
		})
	})
})

var _ = Describe("Test app hash", func() {

	It("should return right hash for a helm app", func() {

		wegoapp := wego.Application{
			Spec: wego.ApplicationSpec{
				Branch:         "main",
				URL:            "https://github.com/owner/repo1",
				DeploymentType: wego.DeploymentTypeHelm,
			},
		}
		wegoapp.Name = "nginx"

		app, err := WegoAppToApp(wegoapp)
		Expect(err).ToNot(HaveOccurred())

		appHash := GetAppHash(app)
		expectedHash := getHash(app.GitSourceURL.String(), app.Name, app.Branch)

		Expect(appHash).To(Equal("wego-" + expectedHash))

	})

	It("should return right hash for a kustomize app", func() {
		wegoapp := wego.Application{
			Spec: wego.ApplicationSpec{
				Branch:         "main",
				URL:            "https://github.com/owner/repo1",
				Path:           "custompath",
				DeploymentType: wego.DeploymentTypeKustomize,
			},
		}

		app, err := WegoAppToApp(wegoapp)
		Expect(err).ToNot(HaveOccurred())

		appHash := GetAppHash(app)
		expectedHash := getHash(app.GitSourceURL.String(), app.Path, app.Branch)

		Expect(appHash).To(Equal("wego-" + expectedHash))

	})
})

func getHash(inputs ...string) string {
	final := []byte(strings.Join(inputs, ""))
	return fmt.Sprintf("%x", md5.Sum(final))
}