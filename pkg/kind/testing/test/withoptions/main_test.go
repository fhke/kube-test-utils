package kind_test

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	kindTesting "github.com/fhke/kube-test-utils/pkg/kind/testing"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/kind/pkg/apis/config/defaults"
)

var (
	clientset  = &kubernetes.Clientset{}
	restConfig = &rest.Config{}
	kubeconfig = fmt.Sprintf("/tmp/kube-integration-config-%d", time.Now().Unix())
	nodeImage  = "github.com/fhke/kube-test-utils/kind-integration:temp"
)

func TestMain(m *testing.M) {
	// pull and tag a new node image from the kind default
	mustDockerPullAndTag(defaults.Image, nodeImage)

	exitC := kindTesting.TestMain(
		m,
		kindTesting.WithClientset(clientset),
		kindTesting.WithRestConfig(restConfig),
		kindTesting.WithKubeconfig(kubeconfig),
		kindTesting.WithClusterInitFunc(initCluster),
		kindTesting.WithNodeImage(nodeImage),
		kindTesting.WithDeleteKubeconfig(),
	)

	// our kubeconfig should not exist as we specified WithDeleteKubeconfig()
	fileMustNotExist(kubeconfig)

	os.Exit(exitC)
}

func mustDockerPullAndTag(old, new string) {
	dockerCommand("rmi", new)
	mustDockerCommand("pull", old)
	mustDockerCommand("tag", old, new)
}

func mustDockerCommand(args ...string) {
	err := dockerCommand(args...)
	if err != nil {
		panic(err)
	}
}

func dockerCommand(args ...string) error {
	return exec.Command("docker", args...).Run()
}

func fileMustNotExist(path string) {
	_, err := os.Stat(path)
	if err == nil {
		// file exists
		panic(
			fmt.Sprintf("file %q should not exist", path),
		)
	} else if !errors.Is(err, os.ErrNotExist) {
		panic(
			fmt.Sprintf("unexpected error while checking for file %q: %s", path, err),
		)
	}
}
