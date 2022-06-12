package testing

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/fhke/kube-test-utils/pkg/kind"
)

/*	Execute tests with an ephemeral kind cluster.

	Returns a value to pass to os.Exit() */
func TestMain(m *testing.M, optSetters ...testMainOptionSetter) int {
	// create options
	opts := newTestMainOptions(optSetters...)

	// run kind cluster
	cluster, err := kind.NewKindCluster(
		&kind.NewKindClusterOpts{
			KubeconfigPath: opts.kubeconfigPath,
			NodeImage:      opts.nodeImage,
		},
	)
	panicErr(err, "Failed to create kind cluster")

	// ensure cluster is deleted
	defer deleteCluster(cluster, opts)

	// set clientset and rest config structs if requested
	setClientsetAndConfig(cluster, opts)

	// Wait for cluster to be healthy - 5 minute timeout
	waitForClusterHealthy(cluster, time.Minute*5)

	// run cluster init function
	runInitFunc(cluster, opts)

	// run tests and return result
	return m.Run()
}

func deleteCluster(cluster *kind.KindCluster, opts *testMainOptions) {
	if opts.deleteKubeconfig {
		defer os.Remove(opts.kubeconfigPath)
	}
	cluster.MustDelete()
}

func runInitFunc(cluster *kind.KindCluster, opts *testMainOptions) {
	if opts.clusterInitFunc != nil {
		err := opts.clusterInitFunc(cluster.RestConfig(), cluster.Clientset())
		panicErr(err, "Error running cluster init function")
	}
}

func waitForClusterHealthy(cluster *kind.KindCluster, timeout time.Duration) {
	ctx, can := context.WithTimeout(context.Background(), timeout)
	defer can()
	err := cluster.WaitForClusterHealthy(ctx)
	panicErr(err, "Error while waiting for cluster to become healthy")
}

func setClientsetAndConfig(cluster *kind.KindCluster, opts *testMainOptions) {
	if opts.clientset != nil {
		// set clientset for cluster
		*opts.clientset = *cluster.Clientset()
	}
	if opts.restConfig != nil {
		// set rest config for cluster
		*opts.restConfig = *cluster.RestConfig()
	}
}

func panicErr(err error, msg string) {
	if err != nil {
		panic(
			fmt.Sprintf("%s: %s", msg, err),
		)
	}
}
