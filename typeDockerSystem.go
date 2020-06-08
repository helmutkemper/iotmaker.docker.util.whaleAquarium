package iotmaker_docker_util_whaleAquarium

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerSystem struct {
	cli       *client.Client
	ctx       context.Context
	networkId map[string]string
	imageId   map[string]string
	container map[string]container.ContainerCreateCreatedBody
}
