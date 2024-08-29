package containers

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ListAllContainers() ([]string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to Docker: %w", err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("unable to list containers: %w", err)
	}

	var containerDetails []string
	for _, container := range containers {
		containerDetails = append(containerDetails, fmt.Sprintf("%s: %s", container.ID[:10], container.Image))
	}

	return containerDetails, nil
}

func StartContainer(containerID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("unable to connect to Docker: %w", err)
	}

	if err := cli.ContainerStart(context.Background(), containerID, container.StartOptions{}); err != nil {
		return fmt.Errorf("unable to start container: %w", err)
	}

	return nil
}

func StopContainer(containerID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("unable to connect to Docker: %w", err)
	}

	if err := cli.ContainerStop(context.Background(), containerID, container.StopOptions{}); err != nil {
		return fmt.Errorf("unable to stop container: %w", err)
	}

	return nil
}

func RemoveContainer(containerID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("unable to connect to Docker: %w", err)
	}

	if err := cli.ContainerRemove(context.Background(), containerID, container.RemoveOptions{Force: true}); err != nil {
		return fmt.Errorf("unable to remove container: %w", err)
	}

	return nil
}
