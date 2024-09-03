package docker

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

func CreatePostgres(containerName string, dbName string, rootPass string) {

	cli := CreateClient()

	ctx := context.Background()

	reader, err := cli.ImagePull(ctx, "docker.io/library/postgres", image.PullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	containerConfig := &container.Config{
		Image: "postgres",
		ExposedPorts: nat.PortSet{
			"5432/tcp": struct{}{},
		},
		Env: []string{
			"POSTGRES_PASSWORD=" + rootPass,
			"POSTGRES_DB=" + dbName,
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "5432",
				},
			},
		},
	}

	networkConfig := &network.NetworkingConfig{}

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, nil, containerName)
	if err != nil {
		log.Fatalf("Erro ao criar container: %v", err)
	}

	// Iniciar o container em modo detached
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Erro ao iniciar container: %v", err)
	}

	fmt.Printf("Container %s criado e iniciado em modo detached\n", resp.ID)
}
