package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/images"
	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/containerd/nerdctl/pkg/buildkitutil"
	"github.com/containerd/nerdctl/pkg/cmd/builder"
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

func BuildImage(ctx context.Context, client *containerd.Client, gopts types.GlobalCommandOptions, tag string, buildCtx string) {
	options, err := getBuilderBuildOption(gopts, tag, buildCtx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := builder.Build(ctx, client, options); err != nil {
		fmt.Println(err.Error())
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

func getBuilderBuildOption(gopts types.GlobalCommandOptions, tag string, buildCtx string) (types.BuilderBuildOptions, error) {
	buildKitHost, err := buildkitutil.GetBuildkitHost(gopts.Namespace)
	if err != nil {
		return types.BuilderBuildOptions{}, err
	}

	platform := []string{"amd64"}
	buildContext := buildCtx
	output := ""
	tagValue := []string{tag}
	progress := ""
	filename := ""
	target := ""
	buildArgs := []string{}
	label := []string{}
	noCache := false
	secret := []string{}
	ssh := []string{}
	cacheFrom := []string{}
	cacheTo := []string{}
	rm := false
	iidfile := ""
	quiet := false

	return types.BuilderBuildOptions{
		GOptions:     gopts,
		BuildKitHost: buildKitHost,
		BuildContext: buildContext,
		Output:       output,
		Tag:          tagValue,
		Progress:     progress,
		File:         filename,
		Target:       target,
		BuildArgs:    buildArgs,
		Label:        label,
		NoCache:      noCache,
		Secret:       secret,
		SSH:          ssh,
		CacheFrom:    cacheFrom,
		CacheTo:      cacheTo,
		Rm:           rm,
		IidFile:      iidfile,
		Quiet:        quiet,
		Platform:     platform,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
		Stdin:        os.Stdin,
	}, nil
}
