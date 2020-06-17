package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

// English: Create a image and create and start a container from project inside into server
// Warning: work in progress - buildStatus don't work yet
func NewContainerFromRemoteServerChangeExposedPortAndVolumes(
	newImageName,
	newContainerName,
	serverPath string,
	imageTags []string,
	currentPortList,
	newPortList []nat.Port,
	containersVolumes []mount.Mount,
	buildStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId string) {

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	// image pull and wait (true)
	err = dockerSys.ImageBuildFromRemoteServer(serverPath, newImageName, imageTags, buildStatus)
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateChangeExposedPortAndStart(
		newImageName,
		newContainerName,
		iotmakerDocker.KRestartPolicyUnlessStopped,
		containersVolumes,
		nil,
		currentPortList,
		newPortList,
	)

	return
}
