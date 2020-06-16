package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
)

func NewContainerFromRemoteServerWithNetworkConfiguration(
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
	buildStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId, networkId string) {

	var containersVolumes []mount.Mount
	var imageVolumesList []string
	var containersVolumeTmpList = make([]iotmakerDocker.Mount, 0)
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

	err, imageVolumesList = dockerSys.ImageListExposedVolumesByName(newImageName)
	if err != nil {
		return
	}

	for _, volumePathInsideImage := range imageVolumesList {
		containersVolumeTmpList = append(
			containersVolumeTmpList,
			iotmakerDocker.Mount{
				MountType:   iotmakerDocker.KVolumeMountTypeBind,
				Source:      volumePathInsideImage,
				Destination: volumePathInsideImage,
			},
		)
	}

	err, containersVolumes = factoryWhaleAquarium.NewVolumeMount(containersVolumeTmpList)
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