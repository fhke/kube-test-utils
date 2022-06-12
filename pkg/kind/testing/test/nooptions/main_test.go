package kind_test

import (
	"os"
	"testing"

	kindTesting "github.com/fhke/kube-test-utils/pkg/kind/testing"
)

func TestMain(m *testing.M) {
	// check that TestMain runs correctly with no options set
	os.Exit(
		kindTesting.TestMain(
			m,
		),
	)
}
