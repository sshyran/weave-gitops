package check

import (
	"context"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"

	"github.com/weaveworks/weave-gitops/pkg/kube/kubefakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/weaveworks/weave-gitops/pkg/flux/fluxfakes"
)

var _ = Describe("Check", func() {

	var fakeFluxClient *fluxfakes.FakeFlux
	var fakeKubeClient *kubefakes.FakeKube
	someError := errors.New("some error")
	var context context.Context
	const successfulFluxPreCheckOutput = `► checking prerequisites
✔ Kubernetes 1.21.1 >=1.19.0-0
✔ prerequisites checks passed`
	const expectedFluxVersion = "1.21.1"
	BeforeEach(func() {
		fakeFluxClient = &fluxfakes.FakeFlux{}
		fakeKubeClient = &kubefakes.FakeKube{}
	})

	It("should fail running flux.PreCheck", func() {
		fakeFluxClient.PreCheckReturns("", someError)

		out, err := Pre(context, fakeKubeClient, fakeFluxClient, expectedFluxVersion)
		Expect(err.Error()).To(ContainSubstring(someError.Error()))
		Expect(out).Should(BeEmpty())
	})

	It("should show flux is not installed", func() {

		fakeFluxClient.PreCheckReturns(successfulFluxPreCheckOutput, nil)

		fakeKubeClient.GetNamespacesReturns(&corev1.NamespaceList{
			Items: []corev1.Namespace{},
		}, nil)

		output, err := Pre(context, fakeKubeClient, fakeFluxClient, expectedFluxVersion)
		Expect(err).ShouldNot(HaveOccurred())

		Expect(output).To(Equal(`✔ Kubernetes 1.21.1 >=1.19.0-0
✔ Flux is not installed`))
	})

	It("should fail while getting namespaces", func() {

		fakeFluxClient.PreCheckReturns(successfulFluxPreCheckOutput, nil)

		fakeKubeClient.GetNamespacesReturns(nil, someError)

		out, err := Pre(context, fakeKubeClient, fakeFluxClient, expectedFluxVersion)
		Expect(err.Error()).To(ContainSubstring(someError.Error()))
		Expect(out).Should(BeEmpty())
	})

	It("should fail parsing actual version", func() {

		output, err := validateFluxVersion("", "")
		Expect(err).Should(HaveOccurred())
		Expect(output).To(BeEmpty())

	})

	It("should fail parsing expected version", func() {

		output, err := validateFluxVersion("v0.0.0", "")
		Expect(err).Should(HaveOccurred())
		Expect(output).To(BeEmpty())

	})

	It("should fail when there is no Kubernetes when running flux pre check", func() {

		fakeFluxClient.PreCheckReturns("", nil)

		output, err := Pre(context, fakeKubeClient, fakeFluxClient, expectedFluxVersion)
		Expect(err).Should(HaveOccurred())
		Expect(output).To(BeEmpty())

	})

	It("should fail when expected flux version is no valid", func() {

		fakeFluxClient.PreCheckReturns(successfulFluxPreCheckOutput, nil)

		fakeKubeClient.GetNamespacesReturns(&corev1.NamespaceList{
			Items: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app.kubernetes.io/part-of": "flux",
							"app.kubernetes.io/version": expectedFluxVersion,
						},
					},
				},
			},
		}, nil)

		output, err := Pre(context, fakeKubeClient, fakeFluxClient, "")
		Expect(err).Should(HaveOccurred())
		Expect(output).To(BeEmpty())

	})

	It("should show flux version is compatible", func() {

		fakeFluxClient.PreCheckReturns(successfulFluxPreCheckOutput, nil)

		fakeKubeClient.GetNamespacesReturns(&corev1.NamespaceList{
			Items: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app.kubernetes.io/part-of": "flux",
							"app.kubernetes.io/version": expectedFluxVersion,
						},
					},
				},
			},
		}, nil)

		output, err := Pre(context, fakeKubeClient, fakeFluxClient, expectedFluxVersion)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(output).To(Equal(`✔ Kubernetes 1.21.1 >=1.19.0-0
✗ Flux 1.21.1 =1.21.1
Current flux version is compatible`))
	})

	It("should show that the current flux version is not compatible", func() {

		fakeFluxClient.PreCheckReturns(successfulFluxPreCheckOutput, nil)

		differentFluxVersion := "v0.0.0"

		fakeKubeClient.GetNamespacesReturns(&corev1.NamespaceList{
			Items: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app.kubernetes.io/part-of": "flux",
							"app.kubernetes.io/version": differentFluxVersion,
						},
					},
				},
			},
		}, nil)

		output, err := Pre(context, fakeKubeClient, fakeFluxClient, expectedFluxVersion)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(output).To(Equal(`✔ Kubernetes 1.21.1 >=1.19.0-0
✗ Flux 0.0.0 !=1.21.1
Current flux version is not compatible`))
	})

})
