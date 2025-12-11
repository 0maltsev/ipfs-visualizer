package services

import (
	"context"

	"project/internal/kube"
)

type ClusterService struct {
	kube kube.KubeClient
}

func NewClusterService(k kube.KubeClient) *ClusterService {
	return &ClusterService{kube: k}
}

func (s *ClusterService) List(ctx context.Context) ([]Cluster, error) {
	// TODO: list StatefulSets with label "app=ipfs-cluster"
	return nil, nil
}

func (s *ClusterService) Create(ctx context.Context, req ClusterCreateRequest) error {
	// TODO: create StatefulSet, Service, headless service, CM, Secret
	return nil
}

func (s *ClusterService) Delete(ctx context.Context, clusterId string) error {
	// TODO: remove all kube objects
	return nil
}
