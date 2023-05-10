package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containerd/containerd"
	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/containerd/nerdctl/pkg/cmd/container"
)

func GetContainerList(ctx context.Context, client *containerd.Client, gopts types.GlobalCommandOptions) {
	options := getContainerOptionList(gopts)
	container.List(ctx, client, options)
}

func GetContainerListWithFilter(ctx context.Context, client *containerd.Client, gopts types.GlobalCommandOptions, filter string) {
	options := getContainerOptionList(gopts)
	options.Filters = append(options.Filters, filter)
	container.List(ctx, client, options)
}

func CommitContainer(ctx context.Context, client *containerd.Client, gopts types.GlobalCommandOptions, req string, rawref string) {
	options := getContainerCommitOption(gopts)
	err := container.Commit(ctx, client, rawref, req, options)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("commit complete => %s\n", rawref)
	}
}

func GetContainers(ctx context.Context, client *containerd.Client, gopts types.GlobalCommandOptions, filters []string) ([]containerd.Container, error) {
	containers, err := client.Containers(ctx, filters...)
	if err != nil {
		fmt.Println(err.Error())
		//return containers, err
	}

	/*
		for i, c := range containers {
			//id
			fmt.Println(i, "container.id", c.ID())

			//images
			image, err := c.Image(ctx)
			if err == nil {
				fmt.Println(i, "image.name", image.Name())
			}

			//labels
			labels, err := c.Labels(ctx)
			if err == nil {
				fmt.Println(i, "labels - start")
				for k, v := range labels {
					fmt.Println(k, v)
				}
				fmt.Println(i, "labels - end")
			}
		}
	*/

	return containers, err
}

func getContainerOptionList(gopts types.GlobalCommandOptions) types.ContainerListOptions {

	// psCommand.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	// psCommand.Flags().IntP("last", "n", -1, "Show n last created containers (includes all states)")
	// psCommand.Flags().BoolP("latest", "l", false, "Show the latest created container (includes all states)")
	// psCommand.Flags().Bool("no-trunc", false, "Don't truncate output")
	// psCommand.Flags().BoolP("quiet", "q", false, "Only display container IDs")
	// psCommand.Flags().BoolP("size", "s", false, "Display total file sizes")

	// // Alias "-f" is reserved for "--filter"
	// psCommand.Flags().String("format", "", "Format the output using the given Go template, e.g, '{{json .}}', 'wide'")
	// psCommand.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// 	return []string{"json", "table", "wide"}, cobra.ShellCompDirectiveNoFileComp
	// })
	// psCommand.Flags().StringSliceP("filter", "f", nil, "Filter matches containers based on given conditions")
	all := false
	lastN := -1
	trunc := false
	quiet := false
	size := false
	format := "" //"wide"
	filters := []string{}

	return types.ContainerListOptions{
		Stdout:   os.Stdout,
		GOptions: gopts,
		All:      all,
		LastN:    lastN,
		Truncate: trunc,
		Quiet:    quiet,
		Size:     size,
		Format:   format,
		Filters:  filters,
	}
}

const PocCommitMsg = "poc-commit-message"

func getContainerCommitOption(gopts types.GlobalCommandOptions) types.ContainerCommitOptions {
	author := ""
	message := PocCommitMsg
	pause := false
	change := []string{}

	return types.ContainerCommitOptions{
		Stdout:   os.Stdout,
		GOptions: gopts,
		Author:   author,
		Message:  message,
		Pause:    pause,
		Change:   change,
	}
}
