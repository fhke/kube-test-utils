package testing

import (
	"fmt"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func newTestMainOptions(opts ...testMainOptionSetter) *testMainOptions {
	// create options struct with defaults
	ret := &testMainOptions{
		kubeconfigPath: fmt.Sprintf("/tmp/kind-ci-kubeconfig-%d", time.Now().UnixNano()),
	}

	// set user-provided options
	for _, opt := range opts {
		opt(ret)
	}

	return ret
}

func WithClientset(clientset *kubernetes.Clientset) testMainOptionSetter {
	return func(tmo *testMainOptions) {
		tmo.clientset = clientset
	}
}

func WithRestConfig(config *rest.Config) testMainOptionSetter {
	return func(tmo *testMainOptions) {
		tmo.restConfig = config
	}
}

func WithKubeconfig(path string) testMainOptionSetter {
	return func(tmo *testMainOptions) {
		tmo.kubeconfigPath = path
	}
}

func WithDeleteKubeconfig() testMainOptionSetter {
	return func(tmo *testMainOptions) {
		tmo.deleteKubeconfig = true
	}
}

func WithNodeImage(image string) testMainOptionSetter {
	return func(tmo *testMainOptions) {
		tmo.nodeImage = image
	}
}

func WithClusterInitFunc(f ClusterInitFunc) testMainOptionSetter {
	return func(tmo *testMainOptions) {
		tmo.clusterInitFunc = f
	}
}
