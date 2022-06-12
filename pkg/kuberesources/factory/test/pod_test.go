package factory_test

import (
	"testing"

	"github.com/fhke/kube-test-utils/pkg/kuberesources/factory"
	"github.com/stretchr/testify/assert"
)

func TestNewBasicPod(t *testing.T) {
	t.Parallel()

	pod := factory.NewBasicPod("my-pod", "my-ns", "nginx:mainline", map[string]string{"testlabel": "testvalue"})

	// check metadata
	assert.Equal(t, "my-pod", pod.Name)
	assert.Equal(t, "my-ns", pod.Namespace)
	assert.Len(t, pod.Labels, 1)
	assert.Contains(t, pod.Labels, "testlabel")
	assert.Equal(t, "testvalue", pod.Labels["testlabel"])
	assert.Len(t, pod.Annotations, 0)

	// check spec
	assert.Len(t, pod.Spec.Containers, 1)
	assert.Equal(t, pod.Spec.Containers[0].Image, "nginx:mainline")

}
