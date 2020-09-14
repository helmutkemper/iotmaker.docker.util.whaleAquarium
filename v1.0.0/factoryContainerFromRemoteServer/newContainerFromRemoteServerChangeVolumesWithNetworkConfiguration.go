package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func NewContainerFromRemoteServerChangeVolumesWithNetworkConfiguration(
	newImageName,
	newContainerName string,
	newContainerRestartPolicy iotmakerdocker.RestartPolicy,
	networkAutoConfiguration *iotmakerdocker.NextNetworkAutoConfiguration,
	serverPath string,
	imageTags []string,
	containersVolumes []mount.Mount,
	buildStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (
	imageId,
	containerId,
	networkId string,
	err error,
) {

	var networkConfig *network.NetworkingConfig

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

	err, networkConfig = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	containerId, err = dockerSys.ContainerCreateExposePortsAutomaticallyAndStart(
		newImageName,
		newContainerName,
		newContainerRestartPolicy,
		containersVolumes,
		networkConfig,
	)

	return
}
