package iotmaker_docker_util_whaleAquarium

func (el *DockerSystem) ContainerCreateAndStart(imageName, containerName string, restart RestartPolicy, mountVolumes []mount.Mount, net *network.NetworkingConfig) (error, string) {
	err, id := el.ContainerCreate(imageName, containerName, restart, mountVolumes, net)
	if err != nil {
		return err, ""
	}

	err = el.ContainerStart(id)
	return err, id
}
