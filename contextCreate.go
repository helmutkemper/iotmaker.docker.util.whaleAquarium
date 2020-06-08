package iotmaker_docker_util_whaleAquarium

import "context"

func (el *DockerSystem) contextCreate() {
	el.ctx = context.Background()
}
