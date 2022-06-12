package kind

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/fhke/kube-test-utils/pkg/kuberesources/health"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Wait for ClusterHealthy() to return true
func (k *KindCluster) WaitForClusterHealthy(ctx context.Context) error {
	bo := backoff.NewExponentialBackOff()
	bo.InitialInterval = time.Second
	bo.RandomizationFactor = 2
	bo.Multiplier = 2
	bo.MaxElapsedTime = 0 // rely on the context to end the backoff

	return backoff.Retry(
		func() error {
			// create a child context with a timeout of a minute
			ctx2, can := context.WithTimeout(ctx, time.Minute)
			defer can()

			// check the cluster is healthy
			ready, err := k.ClusterHealthy(ctx2)
			if err != nil {
				// ClusterHealthy returned an error. Don't return as
				// PermanentError as apiserver may still be initialising
				return err
			} else if !ready {
				// no error, but pods not ready.
				return errors.New("cluster not healthy")
			} else {
				// all pods ready
				return nil
			}
		},
		bo,
	)
}

/* Check if all pods in the cluster are healthy, and the default/default
   service account been created indicating a successful kube-controller-manager
   reconcile */
func (k *KindCluster) ClusterHealthy(ctx context.Context) (bool, error) {
	// check if default service account in default namespace exists
	if found, err := k.defaultServiceAccountExists(ctx); err != nil {
		// unexpected error while checking for service account
		return false, fmt.Errorf("error while checking for service account: %s", err)
	} else if !found {
		// service account does not exist
		return false, nil
	}

	// check all pods in the cluster are ready
	if allReady, err := k.AllPodsReady(ctx, ""); err != nil {
		// unexpected error
		return false, fmt.Errorf("error while checking ready status of all pods: %s", err)
	} else if !allReady {
		// one or more pods not ready
		return false, nil
	}

	// if we got here, the cluster is healthy
	return true, nil

}

// Checks if all pods in the given namespace are ready. If namespace is blank, all pods in cluster will be checked
func (k *KindCluster) AllPodsReady(ctx context.Context, namespace string) (bool, error) {
	podList, err := k.Clientset().CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return false, fmt.Errorf("error listing pods: %s", err)
	}

	for _, pod := range podList.Items {
		if !health.PodReady(pod) {
			// pod is not ready, return now as there's no point continuing
			return false, nil
		}
	}

	// if we got here, all pods are ready and healthy
	return true, nil

}

func (k *KindCluster) defaultServiceAccountExists(ctx context.Context) (bool, error) {
	// try to get default/default service account
	_, err := k.clientset.CoreV1().ServiceAccounts("default").Get(ctx, "default", metav1.GetOptions{})

	if err == nil {
		// no error, service account exists
		return true, nil
	} else if kerrors.IsNotFound(err) {
		// not found, service account does not exist
		return false, nil
	} else {
		// unkown error
		return false, err
	}
}
