package kind_test

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// validate that at least one container is using the nodeImage
func TestNodeImageSet(t *testing.T) {
	out, err := dockerCommandWithStdout("ps", "--format", fmt.Sprintf("{{if eq .Image %q}}{{.ID}}{{ end }}", nodeImage))
	assert.NoError(t, err)

	dockerImageRe, err := regexp.Compile(`^\n*[a-z0-9]{10,}.*`)
	assert.NoError(t, err)
	assert.Regexp(t, dockerImageRe, out)
}

func dockerCommandWithStdout(args ...string) (string, error) {
	// create buffer for stdout
	buf := new(bytes.Buffer)
	// create command
	cmd := exec.Command("docker", args...)
	// write stdout to buffer
	cmd.Stdout = buf
	// run command
	if err := cmd.Run(); err != nil {
		return "", err
	}
	// read stdout
	stringBuf := new(strings.Builder)
	_, err := io.Copy(stringBuf, buf)
	if err != nil {
		return "", err
	}

	return stringBuf.String(), nil
}
