package kind

import (
	"fmt"

	"sigs.k8s.io/kind/pkg/cluster"
)

// delete cluster
func (k *KindCluster) Delete() error {
	// delete cluster
	return deleteKindCluster(k.kindProvider, k.name, k.kubeconfigPath)
}

// delete cluster, panic if it fails
func (k *KindCluster) MustDelete() {
	if err := k.Delete(); err != nil {
		panic(
			fmt.Sprintf("Failed to delete kind cluster %q: %s", k.name, err),
		)
	}
}

func deleteKindCluster(provider *cluster.Provider, name, kubeconfigPath string) error {
	return provider.Delete(
		name,
		kubeconfigPath,
	)

}
