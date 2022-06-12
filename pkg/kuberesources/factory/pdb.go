package factory

import (
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewBasicPodDisruptionBudget(name, namespace string, maxUnavailable, minAvailable int, selectorLabels map[string]string) *policyv1.PodDisruptionBudget {
	maxU := intstr.FromInt(maxUnavailable)
	minU := intstr.FromInt(minAvailable)

	return &policyv1.PodDisruptionBudget{
		ObjectMeta: NewObjectMeta(name, namespace, nil, nil),
		Spec: policyv1.PodDisruptionBudgetSpec{
			MaxUnavailable: &maxU,
			MinAvailable:   &minU,
			Selector: &metav1.LabelSelector{
				MatchLabels: selectorLabels,
			},
		},
	}
}
