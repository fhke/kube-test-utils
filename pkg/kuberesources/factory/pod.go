package factory

import (
	corev1 "k8s.io/api/core/v1"
)

func NewBasicPod(name, namespace, image string, labels map[string]string) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: NewObjectMeta(name, namespace, labels, nil),
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "main",
					Image: image,
				},
			},
		},
	}
}
