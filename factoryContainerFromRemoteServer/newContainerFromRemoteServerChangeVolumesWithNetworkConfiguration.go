package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewContainerFromRemoteServerChangeVolumesWithNetworkConfiguration(
	newImageName,
	newContainerName string,
	newContainerRestartPolicy iotmakerDocker.RestartPolicy,
	networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration,
	serverPath string,
	imageTags []string,
	containersVolumes []mount.Mount,
	buildStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId, networkId string) {

	var networkConfig *network.NetworkingConfig

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

	err, networkConfig = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateExposePortsAutomaticallyAndStart(
		newImageName,
		newContainerName,
		newContainerRestartPolicy,
		containersVolumes,
		networkConfig,
	)

	return
}
