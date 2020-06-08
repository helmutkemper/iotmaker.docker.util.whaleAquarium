package iotmaker_docker_util_whaleAquarium

import "github.com/docker/docker/client"

// Negotiate best docker version
func (el *DockerSystem) clientCreate() error {
	var err error

	el.cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	return err
}
