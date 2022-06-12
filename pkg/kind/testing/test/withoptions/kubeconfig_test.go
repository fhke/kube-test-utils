package kind_test

import (
	"context"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Create a kubernetes clientset using our kubeconfig and validate
// that there is one node.
func TestKubeconfigWorks(t *testing.T) {
	t.Parallel()

	// create the config object from kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	assert.NoError(t, err, "Building REST config should not produce an error")
	assert.NotNil(t, config, "REST config should not be nil")

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err, "Building clientset should not produce an error")
	assert.NotNil(t, clientset, "Clientset not be nil")

	// check number of nodes
	ctx, can := context.WithTimeout(context.Background(), time.Minute)
	defer can()
	nodeList, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	assert.NoError(t, err, "Expect list nodes to not return error")
	assert.Len(t, nodeList.Items, 1, "Expect there to be one node")
}
