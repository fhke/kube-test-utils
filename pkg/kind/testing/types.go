package testing

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type (

	// Options for TestMain()
	testMainOptionSetter func(*testMainOptions)
	testMainOptions      struct {
		clientset        *kubernetes.Clientset // an empty Clientset struct to populate
		restConfig       *rest.Config          // an empty Config to populate
		kubeconfigPath   string                // path to a kubeconfig to either create or merge cluster into
		deleteKubeconfig bool                  // delete the kubeconfig when the tests are finished?
		nodeImage        string                // image to run for kind node
		clusterInitFunc  ClusterInitFunc
	}
	ClusterInitFunc func(*rest.Config, kubernetes.Interface) error // clusterInitFunc is run in TestMain before tests
)
