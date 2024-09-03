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

func CreateMysql(containerName string, dbName string, rootPass string, pass string, mysqlUser string) {

	cli := CreateClient()

	ctx := context.Background()

	reader, err := cli.ImagePull(ctx, "docker.io/library/mysql", image.PullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	containerConfig := &container.Config{
		Image: "mysql",
		ExposedPorts: nat.PortSet{
			"3306/tcp": struct{}{},
		},
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + pass,
			"MYSQL_DATABASE=" + dbName,
			"MYSQL_USER=" + mysqlUser,
			"MYSQL_PASSWORD=" + rootPass,
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3306/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "3306",
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
