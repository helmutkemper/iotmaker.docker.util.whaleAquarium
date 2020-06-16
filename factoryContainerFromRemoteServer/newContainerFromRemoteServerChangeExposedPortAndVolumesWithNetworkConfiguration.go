package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration(
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
	currentPortList,
	newPortList []nat.Port,
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

	// image pull and wait (true)
	err = dockerSys.ImageBuildFromRemoteServer(serverPath, newImageName, imageTags, buildStatus)
	if err != nil {
		return
	}

	err, networkConfig = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateChangeExposedPortAndStart(
		newImageName,
		newContainerName,
		newContainerRestartPolicy,
		containersVolumes,
		networkConfig,
		currentPortList,
		newPortList,
	)

	return
}
