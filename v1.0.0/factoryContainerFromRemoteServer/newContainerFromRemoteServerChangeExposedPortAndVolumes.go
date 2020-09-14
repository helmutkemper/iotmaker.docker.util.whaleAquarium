package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// English: Create a image and create and start a container from project inside into server
// Warning: work in progress - buildStatus don't work yet
func NewContainerFromRemoteServerChangeExposedPortAndVolumes(
	newImageName,
	newContainerName,
	serverPath string,
	imageTags []string,
	portList nat.PortMap,
	containersVolumes []mount.Mount,
	buildStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId string) {

	// init docker
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	// image pull and wait (true)
	_, err = dockerSys.ImageBuildFromRemoteServer(serverPath, newImageName, imageTags, buildStatus)
	if err != nil {
		return
	}

	containerId, err = dockerSys.ContainerCreateAndStart(
		newImageName,
		newContainerName,
		iotmakerdocker.KRestartPolicyUnlessStopped,
		portList,
		containersVolumes,
		nil,
	)

	return
}
