package kind

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/kind/pkg/cluster"
)

// return saved clientset
func (k *KindCluster) Clientset() *kubernetes.Clientset {
	return k.clientset
}

// return saved rest config
func (k *KindCluster) RestConfig() *rest.Config {
	return k.restConfig
}

// get a clientset and rest config for a kind cluster
func newClientsetAndConfig(provider *cluster.Provider, name string) (*rest.Config, *kubernetes.Clientset, error) {
	// get kubeconfig as a string
	kubeconfig, err := provider.KubeConfig(name, false)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get kubeconfig for kind cluster %q: %s", name, err)
	}

	// get rest config
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert kubeconfig to REST config for kind cluster %q: %s", name, err)
	}

	// create clientset
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create clientset from rest config: %s", err)
	}

	return config, cs, err
}
