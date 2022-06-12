package health_test

import (
	"testing"

	"github.com/fhke/kube-test-utils/pkg/kuberesources/factory"
	"github.com/fhke/kube-test-utils/pkg/kuberesources/health"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestPodScheduledAndReady(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		pod           corev1.Pod
		shouldBeReady bool
	}{
		{
			shouldBeReady: false,
			pod: corev1.Pod{
				ObjectMeta: factory.NewObjectMeta("test-1", "default", nil, nil),
				Status: corev1.PodStatus{
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.ContainersReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			shouldBeReady: false,
			pod: corev1.Pod{
				ObjectMeta: factory.NewObjectMeta("test-2", "default", nil, nil),
				Status: corev1.PodStatus{
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.PodScheduled,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			shouldBeReady: true,
			pod: corev1.Pod{
				ObjectMeta: factory.NewObjectMeta("test-3", "default", nil, nil),
				Status: corev1.PodStatus{
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.ContainersReady,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.PodScheduled,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
	} {
		isReady := health.PodReady(testCase.pod)
		assert.Equal(t, testCase.shouldBeReady, isReady, testCase.pod.Name)
	}
}
