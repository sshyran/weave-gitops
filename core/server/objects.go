package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/fluxcd/helm-controller/api/v2beta1"
	helmv2 "github.com/fluxcd/helm-controller/api/v2beta1"
	"github.com/fluxcd/pkg/ssa"
	"github.com/weaveworks/weave-gitops/core/clustersmngr"
	"github.com/weaveworks/weave-gitops/core/server/types"
	pb "github.com/weaveworks/weave-gitops/pkg/api/core"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getRawHelmReleaseInventory(ctx context.Context, helmRelease v2beta1.HelmRelease, c clustersmngr.Client, cluster string) ([]schema.GroupVersionKind, error) {
	storageNamespace := helmRelease.GetStorageNamespace()

	storageName := helmRelease.GetReleaseName()

	storageVersion := helmRelease.Status.LastReleaseRevision
	if storageVersion < 1 {
		// skip release if it failed to install
		return nil, nil
	}

	storageSecret := &v1.Secret{}
	secretName := fmt.Sprintf("sh.helm.release.v1.%s.v%v", storageName, storageVersion)
	key := client.ObjectKey{
		Name:      secretName,
		Namespace: storageNamespace,
	}

	if err := c.Get(ctx, cluster, key, storageSecret); err != nil {
		return nil, err
	}

	releaseData, releaseFound := storageSecret.Data["release"]
	if !releaseFound {
		return nil, fmt.Errorf("failed to decode the Helm storage object for HelmRelease '%s'", helmRelease.Name)
	}

	byteData, err := base64.StdEncoding.DecodeString(string(releaseData))
	if err != nil {
		return nil, err
	}

	var magicGzip = []byte{0x1f, 0x8b, 0x08}
	if bytes.Equal(byteData[0:3], magicGzip) {
		r, err := gzip.NewReader(bytes.NewReader(byteData))
		if err != nil {
			return nil, err
		}

		defer r.Close()

		uncompressedByteData, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}

		byteData = uncompressedByteData
	}

	storage := types.HelmReleaseStorage{}
	if err := json.Unmarshal(byteData, &storage); err != nil {
		return nil, fmt.Errorf("failed to decode the Helm storage object for HelmRelease '%s': %w", helmRelease.Name, err)
	}

	objects, err := ssa.ReadObjects(strings.NewReader(storage.Manifest))
	if err != nil {
		return nil, fmt.Errorf("failed to read the Helm storage object for HelmRelease '%s': %w", helmRelease.Name, err)
	}

	result := []schema.GroupVersionKind{}

	for _, obj := range objects {
		result = append(result, obj.GroupVersionKind())
	}
	return result, nil
}

func (cs *coreServer) ListObjects(ctx context.Context, msg *pb.ListObjectsRequest) (*pb.ListObjectsResponse, error) {
	clustersClient := clustersmngr.ClientFromCtx(ctx)

	gvk := types.GetGVK(msg.Kind)
	if gvk == nil {
		return nil, fmt.Errorf("Looking up objects of type %v not supported", msg.Kind)
	}

	clist := clustersmngr.NewClusteredList(func() client.ObjectList {
		l := unstructured.UnstructuredList{}
		l.SetGroupVersionKind(*gvk)
		return &l
	})

	opts := []client.ListOption{}
	if msg.Pagination != nil {
		opts = append(opts, client.Limit(msg.Pagination.PageSize))
		opts = append(opts, client.Continue(msg.Pagination.PageToken))
	}

	respErrors := []*pb.ListError{}

	if err := clustersClient.ClusteredList(ctx, clist, true, opts...); err != nil {
		var errs clustersmngr.ClusteredListError
		if !errors.As(err, &errs) {
			return nil, err
		}

		for _, e := range errs.Errors {
			respErrors = append(respErrors, &pb.ListError{ClusterName: e.Cluster, Namespace: e.Namespace, Message: e.Err.Error()})
		}
	}

	var results []*pb.Object

	for n, lists := range clist.Lists() {
		for _, l := range lists {
			list, ok := l.(*unstructured.UnstructuredList)
			if !ok {
				continue
			}
			for _, object := range list.Items {
				var inv []schema.GroupVersionKind
				if msg.Kind == helmv2.HelmReleaseKind {
					var helmRelease helmv2.HelmRelease
					err := runtime.DefaultUnstructuredConverter.FromUnstructured(object.Object, &helmRelease)
					if err != nil {
						return nil, err
					}
					inv, err = getRawHelmReleaseInventory(ctx, helmRelease, clustersClient, n)
					if err != nil {
						return nil, err
					}
				}

				k, err := types.K8sObjectToProto(&object, n, inv)

				if err != nil {
					return nil, fmt.Errorf("converting items: %w", err)
				}

				results = append(results, k)
			}
		}
	}

	return &pb.ListObjectsResponse{
		Objects:       results,
		NextPageToken: clist.GetContinue(),
		Errors:        respErrors,
	}, nil
}

func (cs *coreServer) GetObject(ctx context.Context, msg *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	clustersClient := clustersmngr.ClientFromCtx(ctx)

	gvk := types.GetGVK(msg.Kind)
	if gvk == nil {
		return nil, fmt.Errorf("Looking up objects of type %v not supported", msg.Kind)
	}

	obj := unstructured.Unstructured{}
	obj.SetGroupVersionKind(*gvk)

	key := client.ObjectKey{
		Name:      msg.Name,
		Namespace: msg.Namespace,
	}

	if err := clustersClient.Get(ctx, msg.ClusterName, key, &obj); err != nil {
		return nil, err
	}

	var inv []schema.GroupVersionKind = nil
	if msg.Kind == helmv2.HelmReleaseKind {
		var helmRelease helmv2.HelmRelease
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, &helmRelease)
		if err != nil {
			return nil, err
		}
		inv, err = getRawHelmReleaseInventory(ctx, helmRelease, clustersClient, msg.ClusterName)
		if err != nil {
			return nil, err
		}
	}

	res, err := types.K8sObjectToProto(&obj, msg.ClusterName, inv)

	if err != nil {
		return nil, fmt.Errorf("converting kustomization to proto: %w", err)
	}

	return &pb.GetObjectResponse{Object: res}, nil
}
