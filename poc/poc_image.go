package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/images"
	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/containerd/nerdctl/pkg/cmd/image"
	"github.com/containerd/nerdctl/pkg/imgutil"
)

func GetImageList(ctx context.Context, client *containerd.Client, gopts types.GlobalCommandOptions) {
	options := getImageGlobalOption(gopts)
	image.ListCommandHandler(ctx, client, options)
}

func GetImages(ctx context.Context, client *containerd.Client, gopts types.GlobalCommandOptions, ref string) ([]images.Image, error) {
	filters := []string{
		fmt.Sprintf("reference=%s", ref),
	}
	nameAndRefFilter := []string{}

	return image.List(ctx, client, filters, nameAndRefFilter)
}

func PrintImages(images []images.Image) {
	for _, i := range images {
		repo, tag := imgutil.ParseRepoTag(i.Name)
		fmt.Printf("%s\t%s\t%s\n", repo, tag, i.CreatedAt.String())
	}
}

func getImageGlobalOption(gopts types.GlobalCommandOptions) types.ImageListOptions {

	var filters []string

	quiet := false
	noTrunc := false
	format := ""
	var inputFilters []string
	digests := false
	names := false

	return types.ImageListOptions{
		GOptions:         gopts,
		Quiet:            quiet,
		NoTrunc:          noTrunc,
		Format:           format,
		Filters:          inputFilters,
		NameAndRefFilter: filters,
		Digests:          digests,
		Names:            names,
		All:              true,
		Stdout:           os.Stdout,
	}
}
