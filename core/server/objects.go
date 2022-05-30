package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/weaveworks/weave-gitops/core/clustersmngr"
	"github.com/weaveworks/weave-gitops/core/server/types"
	pb "github.com/weaveworks/weave-gitops/pkg/api/core"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

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
				k, err := types.K8sObjectToProto(&object, n)

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

	res, err := types.K8sObjectToProto(&obj, msg.ClusterName)

	if err != nil {
		return nil, fmt.Errorf("converting kustomization to proto: %w", err)
	}

	return &pb.GetObjectResponse{Object: res}, nil
}
