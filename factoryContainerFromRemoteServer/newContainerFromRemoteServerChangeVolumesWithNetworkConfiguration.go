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
	networkName string,
	networkDrive iotmakerDocker.NetworkDrive,
	networkScope,
	networkSubnet,
	networkGateway string,
	serverPath string,
	imageTags []string,
	containersVolumes []mount.Mount,
	buildStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId, networkId string) {

	var networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration
	var networkConfig *network.NetworkingConfig

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, networkId, networkAutoConfiguration = dockerSys.NetworkCreate(
		networkName,
		networkDrive,
		networkScope,
		networkSubnet,
		networkGateway,
	)
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

	err, containerId = dockerSys.ContainerCreateAndStart(
		newImageName,
		newContainerName,
		newContainerRestartPolicy,
		containersVolumes,
		networkConfig,
	)

	return
}
