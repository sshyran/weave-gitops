package types

import (
	"bytes"

	helmv2 "github.com/fluxcd/helm-controller/api/v2beta1"
	kustomizev1 "github.com/fluxcd/kustomize-controller/api/v1beta2"
	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"
	pb "github.com/weaveworks/weave-gitops/pkg/api/core"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetGVK(kind string) *schema.GroupVersionKind {
	var result *schema.GroupVersionKind = nil
	switch kind {
	case kustomizev1.KustomizationKind:
		gvk := kustomizev1.GroupVersion.WithKind(kustomizev1.KustomizationKind)
		result = &gvk
	case helmv2.HelmReleaseKind:
		gvk := helmv2.GroupVersion.WithKind(helmv2.HelmReleaseKind)
		result = &gvk
	case sourcev1.GitRepositoryKind:
		gvk := sourcev1.GroupVersion.WithKind(sourcev1.GitRepositoryKind)
		result = &gvk
	case sourcev1.HelmChartKind:
		gvk := sourcev1.GroupVersion.WithKind(sourcev1.HelmChartKind)
		result = &gvk
	case sourcev1.HelmRepositoryKind:
		gvk := sourcev1.GroupVersion.WithKind(sourcev1.HelmRepositoryKind)
		result = &gvk
	case sourcev1.BucketKind:
		gvk := sourcev1.GroupVersion.WithKind(sourcev1.BucketKind)
		result = &gvk
	}
	return result
}

func K8sObjectToProto(object client.Object, clusterName string) (*pb.Object, error) {
	var buf bytes.Buffer
	serializer := json.NewSerializer(json.DefaultMetaFactory, nil, nil, false)
	if err := serializer.Encode(object, &buf); err != nil {
		return nil, err
	}

	obj := &pb.Object{
		Object:      buf.String(),
		ClusterName: clusterName,
	}

	return obj, nil
}
