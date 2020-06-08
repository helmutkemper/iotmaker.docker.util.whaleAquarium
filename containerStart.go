package iotmaker_docker_util_whaleAquarium

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerStart(id string) error {
	return el.cli.ContainerStart(el.ctx, id, types.ContainerStartOptions{})
}
