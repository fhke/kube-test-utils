package factory_test

import (
	"testing"

	"github.com/fhke/kube-test-utils/pkg/kuberesources/factory"
	"github.com/stretchr/testify/assert"
)

func TestNewBasicPodDisruptionBudget(t *testing.T) {
	t.Parallel()

	pdb := factory.NewBasicPodDisruptionBudget("my-pdb", "my-namespace", 0, 1, map[string]string{"testlabel": "testvalue"})

	// check metadata
	assert.Equal(t, "my-pdb", pdb.Name)
	assert.Equal(t, "my-namespace", pdb.Namespace)
	assert.Len(t, pdb.Labels, 0)
	assert.Len(t, pdb.Annotations, 0)

	// check spec
	assert.Equal(t, pdb.Spec.MaxUnavailable.IntVal, int32(0))
	assert.Equal(t, pdb.Spec.MinAvailable.IntVal, int32(1))
	assert.Len(t, pdb.Spec.Selector.MatchLabels, 1)
	assert.Equal(t, "testvalue", pdb.Spec.Selector.MatchLabels["testlabel"])
}
