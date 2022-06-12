package kind

import (
	"fmt"
	"time"

	"sigs.k8s.io/kind/pkg/cluster"
)

// Create a kind cluster
func NewKindCluster(opts *NewKindClusterOpts) (*KindCluster, error) {
	kindProvider := cluster.NewProvider()

	// generate cluster name
	// TODO - add Name/NamePrefix options to NewKindClusterOpts
	name := fmt.Sprintf("kind-ci-%d", time.Now().UnixNano())

	// create cluster
	if err := kindProvider.Create(
		name,
		opts.newKindCreateOpts()...,
	); err != nil {
		deleteKindCluster(kindProvider, name, opts.KubeconfigPath)
		return nil, fmt.Errorf("could not create kind cluster %q: %s", name, err)
	}

	// get clientset and rest config for cluster
	config, cs, err := newClientsetAndConfig(kindProvider, name)
	if err != nil {
		deleteKindCluster(kindProvider, name, opts.KubeconfigPath)
		return nil, fmt.Errorf("could not generate clientset for cluster: %s", err)
	}

	// create KindCluster
	return &KindCluster{
		name:           name,
		kubeconfigPath: opts.KubeconfigPath,
		kindProvider:   kindProvider,
		clientset:      cs,
		restConfig:     config,
	}, nil

}
