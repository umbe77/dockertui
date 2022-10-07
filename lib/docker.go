package lib

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func RefreshTable(ctx context.Context) []types.Container {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	return containers
}
