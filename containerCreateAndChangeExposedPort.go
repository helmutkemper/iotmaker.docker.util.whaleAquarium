package iotmaker_docker_util_whaleAquarium

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

// Create a container
//   imageName: image name for download and pull
//   containerName: unique container name
//   RestartPolicy:
//      KRestartPolicyNo - Do not automatically restart the container. (the
//          default)
//      KRestartPolicyOnFailure - Restart the container if it exits due to an
//          error, which manifests as a non-zero exit code.
//      KRestartPolicyAlways - Always restart the container if it stops. If it is
//          manually stopped, it is restarted only when Docker daemon restarts or
//          the container itself is manually restarted. (See the second bullet
//          listed in restart policy details)
//      KRestartPolicyUnlessStopped - Similar to always, except that when the
//          container is stopped (manually or otherwise), it is not restarted
//          even after Docker daemon restarts.
//   mountVolumes: please use a factoryWhaleAquarium.NewVolumeMount()
//      for a complete list of volumes exposed by image, use
//      ImageListExposedVolumes(id) and ImageListExposedVolumesByName(name)
func (el *DockerSystem) ContainerCreateAndChangeExposedPort(imageName, containerName string, restart RestartPolicy, mountVolumes []mount.Mount, net *network.NetworkingConfig, currentPort, changeToPort []nat.Port) (error, string) {
	var err error
	var imageId string
	var portExposedList nat.PortMap
	var resp container.ContainerCreateCreatedBody

	err, imageId = el.ImageFindIdByName(imageName)
	if err != nil {
		return err, ""
	}

	err, portExposedList = el.ImageMountNatPortListChangeExposed(imageId, currentPort, changeToPort)
	if err != nil {
		return err, ""
	}

	if len(el.container) == 0 {
		el.container = make(map[string]container.ContainerCreateCreatedBody)
	}

	resp, err = el.cli.ContainerCreate(
		el.ctx,
		&container.Config{
			Image: imageName,
		},
		&container.HostConfig{
			PortBindings: portExposedList,
			RestartPolicy: container.RestartPolicy{
				Name: restart.String(),
			},
			Resources: container.Resources{},
			Mounts:    mountVolumes,
		},
		net,
		nil,
		containerName,
	)
	if err != nil {
		return err, ""
	}

	el.container[resp.ID] = resp

	return nil, resp.ID
}
