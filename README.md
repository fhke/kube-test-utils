# Test utilities for Kubernetes

This library contains an opinionated set of utilities for running go tests on an ephemeral Kubernetes cluster using [kind](https://github.com/kubernetes-sigs/kind/).

## Quickstart

To add a kind cluster to your go tests, create a file called `main_test.go` in your test directory containing the following code:

```golang
import kindTesting "github.com/fhke/kube-test-utils/pkg/kind/testing"

var clientset = &kubernetes.Clientset{}

func TestMain(m *testing.M) {
	os.Exit(
		kindTesting.TestMain(
			m,
			kindTesting.WithClientset(clientset),
		),
	)
}
```

The `clientset` struct will be set to an admin clientset for the newly created cluster. Once the tests have been run, the kind cluster will be deleted.

## Options

_n.b. none of these options are required_

- `WithClientset` : Specify a pointer to an empty Kubernetes clientset that will be set for the cluster
- `WithRestConfig` : Specify a pointer to an empty Kubernetes rest config that will be set for the cluster
- `WithKubeconfig` : Specify a path to a file that will contain a kubeconfig for the cluster
- `WithClusterInitFunc` : Specify a function to be run before the tests. This allows you to configure the environment before tests run
- `WithNodeImage` : Specify an alternative kind node image
- `WithDeleteKubeConfig` : Delete the kubeconfig file after tests have run