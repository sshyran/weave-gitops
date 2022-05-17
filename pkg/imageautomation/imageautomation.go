package imageautomation

import (
	"context"
	"fmt"

	imgautov1 "github.com/fluxcd/image-automation-controller/api/v1beta1"
	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/weaveworks/weave-gitops/core/clustersmngr"
	pb "github.com/weaveworks/weave-gitops/pkg/api/imageautomation"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type imageAutomationServer struct {
	pb.UnimplementedImageAutomationServer

	logger         logr.Logger
	clientsFactory clustersmngr.ClientsFactory
}

func Hydrate(ctx context.Context, mux *runtime.ServeMux, srv pb.ImageAutomationServer) error {
	return pb.RegisterImageAutomationHandlerServer(ctx, mux, srv)
}

func NewServer(logger logr.Logger, clientsFactory clustersmngr.ClientsFactory) pb.ImageAutomationServer {
	return &imageAutomationServer{logger: logger, clientsFactory: clientsFactory}
}

func (ias *imageAutomationServer) ListImageAutomations(ctx context.Context, msg *pb.ListImageAutomationRequest) (*pb.ListImageAutomationResponse, error) {
	c := clustersmngr.ClientFromCtx(ctx)

	clist := clustersmngr.NewClusteredList(func() client.ObjectList {
		return &imgautov1.ImageUpdateAutomationList{}
	})

	if err := c.ClusteredList(ctx, clist); err != nil {
		return nil, fmt.Errorf("listing image update automations: %w", err)
	}

	results := []*pb.ImageUpdateAutomation{}

	for _, lists := range clist.Lists() {
		for _, l := range lists {
			list, ok := l.(*imgautov1.ImageUpdateAutomationList)
			if !ok {
				continue
			}

			for _, iua := range list.Items {
				m := &pb.ImageUpdateAutomation{
					Kind:        iua.Kind,
					Apiversion:  iua.APIVersion,
					Name:        iua.Name,
					Namespace:   iua.Namespace,
					ClusterName: iua.ClusterName,
					Spec: &pb.ImageUpdateAutomationSpec{
						SourceRef: &pb.CrossNamespaceSourceReference{
							Kind:      iua.Spec.SourceRef.Kind,
							Name:      iua.Spec.SourceRef.Name,
							Namespace: iua.Spec.SourceRef.Namespace,
						},
						GitSpec: &pb.GitSpec{
							Commit: &pb.CommitSpec{
								Author: &pb.CommitUser{Name: iua.Spec.GitSpec.Commit.Author.Name},
							},
						},
						Update: &pb.UpdateStrategy{
							Strategy: string(iua.Spec.Update.Strategy),
							Path:     iua.Spec.Update.Path,
						},
					},
				}

				results = append(results, m)
			}
		}
	}

	return &pb.ListImageAutomationResponse{Automations: results}, nil
}
