package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// English: Create a image and create and start a container from project inside into server
// Warning: work in progress - buildStatus don't work yet
func NewContainerFromRemoteServerChangeVolumes(
	newImageName,
	newContainerName,
	serverPath string,
	imageTags []string,
	containersVolumes []mount.Mount,
	buildStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (
	imageId,
	containerId string,
	err error,
) {

	// init docker
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	_, err = dockerSys.ImageBuildFromRemoteServer(serverPath, newImageName, imageTags, buildStatus)
	if err != nil {
		return
	}

	containerId, err = dockerSys.ContainerCreateExposePortsAutomaticallyAndStart(
		newImageName,
		newContainerName,
		iotmakerdocker.KRestartPolicyUnlessStopped,
		containersVolumes,
		nil,
	)

	return
}
