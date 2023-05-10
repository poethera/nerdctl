package main

import (
	"testing"

	"github.com/containerd/nerdctl/pkg/testutil"
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