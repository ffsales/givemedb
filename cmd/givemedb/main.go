package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {

	fmt.Println("Hello World")

	cli := CreateClient()

	// CreateContainer()

	CreateMysql(cli)

	ListContainers(cli)

}

func ListContainers(cli *client.Client) {
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%s %s\n", ctr.ID, ctr.Image)
	}

}

func CreateContainer(cli *client.Client) {

	defer cli.Close()

	ctx := context.Background()

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

func CreateMysql(cli *client.Client) {
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
			"MYSQL_ROOT_PASSWORD=root1234",
			"MYSQL_DATABASE=aula-docker",
			"MYSQL_USER=userdocker",
			"MYSQL_PASSWORD=userdockerpwd",
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

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, nil, "my-mysql-container")
	if err != nil {
		log.Fatalf("Erro ao criar container: %v", err)
	}

	// Iniciar o container em modo detached
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Erro ao iniciar container: %v", err)
	}

	fmt.Printf("Container %s criado e iniciado em modo detached\n", resp.ID)

}

func CreateClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return cli
}
