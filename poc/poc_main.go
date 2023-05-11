package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/containerd/nerdctl/pkg/clientutil"
)

const Namespace_k8s = "k8s.io"
const UUID_SHORT_SIZE = 8

func main() {
	ns := Namespace_k8s //"default"
	ctx := context.Background()
	gopts := getGlobalOption(ns)

	client, ctx, cancel, err := clientutil.NewClient(ctx, ns, gopts.Address)
	if err != nil {
		println(err.Error())
		return
	}
	defer cancel()

	//image build
	tag := "build-test-tag:v0.1"
	dockerfile := fmt.Sprint(`FROM busybox 
CMD ["echo", "build-test-echo"]
`)
	buildCtx, err := createBuildContext(dockerfile)
	if err == nil {
		BuildImage(ctx, client, gopts, tag, buildCtx)
	}
	return

	//container list
	GetContainerList(ctx, client, gopts)

	//image list
	GetImageList(ctx, client, gopts)

	//container list with filter
	GetContainerListWithFilter(ctx, client, gopts, "name=samplepod")

	//container commit
	//CommitContainer(ctx, client, gopts, "samplepod-poc", "samplepod-commit/v0.2") //-> test, 1f76330
	//CommitContainer(ctx, client, gopts, "1f76330", "samplepod-commit/v0.2")	//-> test

	label_k := "app"
	label_v := "samplepod-poc"
	new_repo_tag := fmt.Sprintf("%s-commit:v0.1", label_v)

	filters := []string{
		fmt.Sprintf("labels.%s==%s", label_k, label_v),
	}
	containers, err := GetContainers(ctx, client, gopts, filters)
	var container_id_short string

	if err == nil && len(containers) == 1 {
		fmt.Println("founded", "id", containers[0].ID())
		container_id_short = containers[0].ID()[:UUID_SHORT_SIZE]

		CommitContainer(ctx, client, gopts, container_id_short, new_repo_tag)

		images, err := GetImages(ctx, client, gopts, new_repo_tag)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Print("query image => ")
			PrintImages(images)
		}
	}

}

func getGlobalOption(ns string) types.GlobalCommandOptions {
	debug := false
	debugFull := false
	address := "/run/containerd/containerd.sock"
	namespace := ns
	snapshotter := ""
	cniPath := ""
	cniConfigPath := ""
	dataRoot := ""
	cgroupManager := ""
	insecureRegistry := false
	hostsDir := []string{}
	experimental := false

	return types.GlobalCommandOptions{
		Debug:            debug,
		DebugFull:        debugFull,
		Address:          address,
		Namespace:        namespace,
		Snapshotter:      snapshotter,
		CNIPath:          cniPath,
		CNINetConfPath:   cniConfigPath,
		DataRoot:         dataRoot,
		CgroupManager:    cgroupManager,
		InsecureRegistry: insecureRegistry,
		HostsDir:         hostsDir,
		Experimental:     experimental,
	}
}

func createBuildContext(dockerfile string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "nerdctl-build-test")
	if err != nil {
		return "", err
	}
	if err = os.WriteFile(filepath.Join(tmpDir, "Dockerfile"), []byte(dockerfile), 0644); err != nil {
		return "", err
	}
	return tmpDir, nil
}
