package factorycontainerfromremoteserver

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// NewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration
// (English): Create a image and create and start a container from project inside into
// server
//
// NewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration
// (Português): Cria uma imagem e inicializa o container baseado no conteúdo do servidor
func NewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration(
	newImageName,
	newContainerName string,
	newContainerRestartPolicy iotmakerdocker.RestartPolicy,
	networkAutoConfiguration *iotmakerdocker.NextNetworkAutoConfiguration,
	serverPath string,
	imageTags []string,
	portList nat.PortMap,
	containersVolumes []mount.Mount,
	buildStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (
	imageId string,
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

	// image pull and wait (true)
	_, err = dockerSys.ImageBuildFromRemoteServer(serverPath, newImageName, imageTags, buildStatus)
	if err != nil {
		return
	}

	networkConfig, err = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	containerId, err = dockerSys.ContainerCreateAndStart(
		newImageName,
		newContainerName,
		newContainerRestartPolicy,
		portList,
		containersVolumes,
		networkConfig,
	)

	return
}
