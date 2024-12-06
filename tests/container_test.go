package tests

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/image"
	"io"
	"log"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func TestCreateNewContainer(t *testing.T) {

	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	// Define the container configuration
	containerConfig := &container.Config{
		Image: "alpine",                          // Specify the image to use
		Cmd:   []string{"echo", "Hello, World!"}, // Command to run in the container
	}

	// Define host configuration (optional)
	hostConfig := &container.HostConfig{
		AutoRemove: true, // Automatically remove the container when it exits
	}

	// Context for the API call
	ctx := context.Background()

	// Pull the image if it doesn't exist locally
	data, err := cli.ImagePull(ctx, "alpine", image.PullOptions{})
	if err != nil {
		log.Fatalf("Failed to pull Docker image: %v", err)
	}
	defer data.Close()

	// Print the pull response
	dataBytes, err := io.ReadAll(data)
	if err != nil {
		log.Fatalf("Failed to read pull response: %v", err)
	}
	fmt.Println(string(dataBytes))

	fmt.Println("Image pulled successfully!")

	// Create the container
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "my-container")
	if err != nil {
		log.Fatalf("Failed to create Docker container: %v", err)
	}
	fmt.Printf("Container created with ID: %s\n", resp.ID)

	// Start the container
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Failed to start Docker container: %v", err)
	}
	fmt.Println("Container started successfully!")

	// Optionally, wait for the container to finish
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case status := <-statusCh:
		fmt.Printf("Container exited with status code: %d\n", status.StatusCode)
	case err := <-errCh:
		log.Fatalf("Error while waiting for container: %v", err)
	}
}
