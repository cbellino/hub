package commands

import (
	"fmt"
	"testing"

	"github.com/github/hub/Godeps/_workspace/src/github.com/bmizerany/assert"
)

func TestReplaceCheckoutParam(t *testing.T) {
	checkoutURL := "https://github.com/github/hub/pull/12"
	args := NewArgs([]string{"checkout", checkoutURL})
	replaceCheckoutParam(args, checkoutURL, "jingweno", "origin/master")

	cmd := args.ToCmd()
	assert.Equal(t, "git checkout --track -B jingweno origin/master", cmd.String())
}

func TestTransformCheckoutArgs(t *testing.T) {
	args := NewArgs([]string{"checkout", "-b", "https://github.com/github/hub/pull/12"})
	err := transformCheckoutArgs(args)

	assert.Equal(t, "Unsupported flag -b when checking out pull request", fmt.Sprintf("%s", err))

	args = NewArgs([]string{"checkout", "--orphan", "https://github.com/github/hub/pull/12"})
	err = transformCheckoutArgs(args)

	assert.Equal(t, "Unsupported flag --orphan when checking out pull request", fmt.Sprintf("%s", err))
}
