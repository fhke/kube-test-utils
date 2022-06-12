package wait

import (
	"context"
	"errors"
	"fmt"

	"github.com/fhke/kube-test-utils/pkg/kuberesources/health"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

// wait for a pod to become ready
func PodReady(ctx context.Context, cs kubernetes.Interface, name, namespace string) error {
	// get watcher
	watcher, err := cs.CoreV1().Pods(namespace).Watch(ctx, metav1.SingleObject(metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
	}))
	if err != nil {
		return err
	}

	defer watcher.Stop()

	for {
		select {
		case event, ok := <-watcher.ResultChan():
			// received event from watcher
			if !ok {
				// channel closed
				return errors.New("watcher channel closed unexpectedly")
			}

			switch eType := event.Type; eType {
			case watch.Added:
				if podIsHealthy(event.Object) {
					return nil
				}
			case watch.Modified:
				if podIsHealthy(event.Object) {
					return nil
				}
			default:
				return fmt.Errorf("watcher returned unexpected event type %s", eType)
			}
		case <-ctx.Done():
			// context completed before the pod became ready
			return ctx.Err()
		}
	}

}

func podIsHealthy(podObject runtime.Object) bool {
	// cast object to corev1.Pod
	pod, ok := podObject.(*v1.Pod)
	if !ok {
		// can't cast object to pod
		return false
	}

	return health.PodReady(*pod)
}
