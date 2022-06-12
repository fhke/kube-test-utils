package kind_test

import (
	"context"
	"testing"
	"time"

	"github.com/fhke/kube-test-utils/pkg/kuberesources/factory"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// create a namespace to test clusterInitFunc
func initCluster(_ *rest.Config, cs kubernetes.Interface) error {
	ctx, can := context.WithTimeout(context.Background(), time.Minute)
	defer can()

	_, err := cs.CoreV1().Namespaces().Create(ctx, &corev1.Namespace{
		ObjectMeta: factory.NewObjectMeta("test-ns", "", nil, nil),
	}, metav1.CreateOptions{})

	return err
}

// check the namespace in the clusterInitFunc was created
func TestInitFuncOutput(t *testing.T) {
	t.Parallel()

	ctx, can := context.WithTimeout(context.Background(), time.Minute)
	defer can()

	_, err := clientset.CoreV1().Namespaces().Get(ctx, "test-ns", metav1.GetOptions{})
	assert.NoError(t, err)
}
