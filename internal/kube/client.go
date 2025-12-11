package kube

import (
	"path/filepath"

	"ipfs-visualizer/config"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
)

type KubeClient struct {
	Clientset *kubernetes.Clientset
}

func NewClient(cfg *config.Config) (KubeClient, error) {
	var (
		restCfg *rest.Config
		err     error
	)

	if cfg.Kube.InCluster {
		restCfg, err = rest.InClusterConfig()
	} else {
		kubeconfig := filepath.Clean(cfg.Kube.Kubeconfig)
		restCfg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		return KubeClient{}, err
	}

	cs, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return KubeClient{}, err
	}

	return KubeClient{Clientset: cs}, nil
}
