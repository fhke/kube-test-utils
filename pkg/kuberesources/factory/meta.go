package factory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func NewObjectMeta(name, namespace string, labels, annotations map[string]string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:        name,
		Namespace:   namespace,
		Labels:      labels,
		Annotations: annotations,
	}

}
