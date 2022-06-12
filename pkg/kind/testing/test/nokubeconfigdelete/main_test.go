package kind_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	kindTesting "github.com/fhke/kube-test-utils/pkg/kind/testing"
)

var kubeconfig = fmt.Sprintf("/tmp/kube-test-integration-%d", time.Now().UnixNano())

func TestMain(m *testing.M) {
	exitC := kindTesting.TestMain(
		m,
		kindTesting.WithKubeconfig(kubeconfig),
	)

	// kubeconfig should still exist after tests complete
	if _, err := os.Stat(kubeconfig); err != nil {
		panic(
			fmt.Sprintf("kubeconfig should still exist at %q: %s", kubeconfig, err),
		)
	}

	os.Exit(exitC)
}
