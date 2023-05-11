package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/containerd/nerdctl/pkg/testutil"
	"gotest.tools/v3/assert"
)

// added
func TestPocImagesList(t *testing.T) {
	base := testutil.NewBaseWithNamespace(t, "k8s.io")
	strs := base.Cmd("images").OutLines()
	for _, s := range strs {
		println(s)
	}
}

func TestPocContainerList(t *testing.T) {
	base := testutil.NewBaseWithNamespace(t, "k8s.io")
	strs := base.Cmd("ps").OutLines()
	for _, s := range strs {
		println(s)
	}
}

func TestPocContainerList2(t *testing.T) {
	base := testutil.NewBaseWithNamespace(t, "default")
	strs := base.Cmd("ps").OutLines()
	for _, s := range strs {
		println(s)
	}
}

func TestPocBuild(t *testing.T) {
	base := testutil.NewBaseWithNamespace(t, "k8s.io")
	imageName := "build-test"

	dockerfile := fmt.Sprintf(`FROM %s
CMD ["echo", "nerdctl-build-test-string"]
	`, testutil.CommonImage)

	buildCtx, err := createBuildContext(dockerfile)
	defer os.RemoveAll(buildCtx)
	assert.NilError(t, err)
	strs := base.Cmd("build", "-t", imageName, buildCtx).OutLines()
	for _, s := range strs {
		println(s)
	}
}
