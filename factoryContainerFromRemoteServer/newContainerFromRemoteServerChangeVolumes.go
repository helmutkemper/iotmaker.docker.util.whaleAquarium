package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"github.com/helmutkemper/iotmaker.docker/util"
)

// English: Create a image and create and start a container from project inside into server
// Warning: work in progress - buildStatus don't work yet
func NewContainerFromRemoteServerChangeVolumes(
	newImageName,
	newContainerName,
	networkName,
	serverPath string,
	imageTags []string,
	containersVolumes []mount.Mount,
	buildStatus chan whaleAquarium.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId string) {

	var nextNetworkConfig *network.NetworkingConfig
	var networkUtil util.NetworkGenerator

	err, networkUtil = factoryDocker.NewContainerNetworkGenerator(networkName, 10, 0, 0, 1)
	if err != nil {
		return
	}

	// init docker
	var dockerSys = whaleAquarium.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err = dockerSys.ImageBuildFromRemoteServer(serverPath, newImageName, imageTags, &buildStatus)
	if err != nil {
		return
	}

	err, nextNetworkConfig = networkUtil.GetNext()
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateAndStart(
		newImageName,
		newContainerName,
		whaleAquarium.KRestartPolicyUnlessStopped,
		containersVolumes,
		nextNetworkConfig,
	)

	return
}
