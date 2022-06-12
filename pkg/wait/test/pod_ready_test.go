package wait_test

import (
	"context"
	"os"
	"testing"
	"time"

	kindTesting "github.com/fhke/kube-test-utils/pkg/kind/testing"
	"github.com/fhke/kube-test-utils/pkg/kuberesources/factory"
	"github.com/fhke/kube-test-utils/pkg/wait"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var clientset = &kubernetes.Clientset{}

// Test wait.PodReady() with a pod that will never go ready
func TestPodReadyCtxTimeout(t *testing.T) {
	t.Parallel()

	// create a context with a cancel func
	ctx, canCtx := context.WithTimeout(context.Background(), time.Minute)
	defer canCtx()

	// create a channel to receive error from goroutine
	errChan := make(chan error, 1)

	// Run PodReady() in a background goroutine
	go waitPodReadyBackground(ctx, clientset, "test-pod-1", "default", errChan)

	// create a pod
	clientset.CoreV1().Pods("default").Create(
		ctx,
		factory.NewBasicPod("test-pod-1", "default", "image-that-doesnt-exist", nil),
		metav1.CreateOptions{},
	)

	// use returned error from PodReady() for assertions
	assert.Error(t, <-errChan)
}

// Test wait.PodReady() with a pod that should become ready
func TestPodReady(t *testing.T) {
	t.Parallel()

	// create a context with a cancel func
	ctx, canCtx := context.WithTimeout(context.Background(), time.Minute)
	defer canCtx()

	// create a channel to receive error from goroutine
	errChan := make(chan error, 1)

	// Run PodReady() in a background goroutine
	go waitPodReadyBackground(ctx, clientset, "test-pod-2", "default", errChan)

	// create a pod
	clientset.CoreV1().Pods("default").Create(
		ctx,
		factory.NewBasicPod("test-pod-2", "default", "nginx:mainline", nil),
		metav1.CreateOptions{},
	)

	// use returned error from PodReady() for assertions
	assert.NoError(t, <-errChan)
}

func waitPodReadyBackground(ctx context.Context, cs kubernetes.Interface, name, namespace string, errChan chan<- error) {
	errChan <- wait.PodReady(ctx, cs, name, namespace)
}

func TestMain(m *testing.M) {
	os.Exit(kindTesting.TestMain(
		m,
		kindTesting.WithClientset(clientset),
	))
}
