package kind_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test the rest config is set
func TestRestConfig(t *testing.T) {
	assert.NotEmpty(t, restConfig.Host)
}
