package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

// English: Create a image and create and start a container from project inside into server
// Warning: work in progress - buildStatus don't work yet
func NewContainerFromRemoteServerChangeVolumes(
	newImageName,
	newContainerName,
	serverPath string,
	imageTags []string,
	containersVolumes []mount.Mount,
	buildStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId string) {

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err = dockerSys.ImageBuildFromRemoteServer(serverPath, newImageName, imageTags, buildStatus)
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateExposePortsAutomaticallyAndStart(
		newImageName,
		newContainerName,
		iotmakerDocker.KRestartPolicyUnlessStopped,
		containersVolumes,
		nil,
	)

	return
}
