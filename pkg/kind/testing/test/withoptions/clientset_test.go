package kind_test

import (
	"context"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/stretchr/testify/assert"
)

// Check that the clientset we generated is valid by listing nodes
func TestClientsetWorks(t *testing.T) {
	t.Parallel()

	// check number of nodes
	ctx, can := context.WithTimeout(context.Background(), time.Minute)
	defer can()
	nodeList, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	assert.NoError(t, err, "Expect list nodes to not return error")
	assert.Len(t, nodeList.Items, 1, "Expect there to be one node")
}
