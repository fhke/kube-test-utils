package kind

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/kind/pkg/cluster"
)

type (
	// Options for NewKindCluster()
	// all fields are optional
	NewKindClusterOpts struct {
		KubeconfigPath string
		NodeImage      string
	}

	KindCluster struct {
		name, kubeconfigPath string                // name of kind cluster & path to kubeconfig
		kindProvider         *cluster.Provider     // kind cluster provider to delete cluster
		clientset            *kubernetes.Clientset // clientset for cluster
		restConfig           *rest.Config          // rest config for cluster
	}
)
