package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {

	fmt.Println("Hello World")

	CreateContainer()

	ListContainers()

}

func ListContainers() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%s %s\n", ctr.ID, ctr.Image)
	}

}

func CreateContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containerConfig := &container.Config{
		Image: "nginx",
		ExposedPorts: nat.PortSet{
			"80/tcp": struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "80",
				},
			},
		},
	}

	networkConfig := &network.NetworkingConfig{}

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, nil, "my-nginx-container")
	if err != nil {
		log.Fatalf("Erro ao criar container: %v", err)
	}

	// Iniciar o container em modo detached
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Erro ao iniciar container: %v", err)
	}

	fmt.Printf("Container %s criado e iniciado em modo detached\n", resp.ID)

}
