package factorycontainerfromremoteserver

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// NewContainerFromRemoteServerWithNetworkConfiguration (English): Create a image and
// create and start a container from project inside into server
//
// NewContainerFromRemoteServerWithNetworkConfiguration (Português): Cria uma imagem e
// inicializa o container baseado no conteúdo do servidor
func NewContainerFromRemoteServerWithNetworkConfiguration(
	newImageName,
	newContainerName string,
	newContainerRestartPolicy iotmakerdocker.RestartPolicy,
	networkAutoConfiguration *iotmakerdocker.NextNetworkAutoConfiguration,
	serverPath string,
	imageTags []string,
	buildStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId, networkId string) {

	var containersVolumes []mount.Mount
	var imageVolumesList []string
	var containersVolumeTmpList = make([]iotmakerdocker.Mount, 0)
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

	imageId, err = dockerSys.ImageFindIdByName(newImageName)
	if err != nil {
		return
	}

	imageVolumesList, err = dockerSys.ImageListExposedVolumesByName(newImageName)
	if err != nil {
		return
	}

	for _, volumePathInsideImage := range imageVolumesList {
		containersVolumeTmpList = append(
			containersVolumeTmpList,
			iotmakerdocker.Mount{
				MountType:   iotmakerdocker.KVolumeMountTypeBind,
				Source:      volumePathInsideImage,
				Destination: volumePathInsideImage,
			},
		)
	}

	containersVolumes, err = iotmakerdocker.NewVolumeMount(containersVolumeTmpList)
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
