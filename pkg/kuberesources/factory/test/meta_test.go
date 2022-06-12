package factory_test

import (
	"testing"

	"github.com/fhke/kube-test-utils/pkg/kuberesources/factory"
	"github.com/stretchr/testify/assert"
)

func TestNewObjectMeta(t *testing.T) {
	t.Parallel()

	om := factory.NewObjectMeta(
		"test-pod",
		"default",
		map[string]string{
			"mylabel": "value1",
		},
		map[string]string{
			"myannotation": "value2",
		},
	)

	assert.Equal(t, "test-pod", om.Name)
	assert.Equal(t, "default", om.Namespace)
	assert.Len(t, om.Labels, 1)
	assert.Contains(t, om.Labels, "mylabel")
	assert.Equal(t, "value1", om.Labels["mylabel"])
	assert.Len(t, om.Annotations, 1)
	assert.Contains(t, om.Annotations, "myannotation")
	assert.Equal(t, "value2", om.Annotations["myannotation"])

}
