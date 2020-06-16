package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
)

// English: Create a image and create and start a container from project inside into server
func NewContainerFromRemoteServer(
	newImageName,
	newContainerName,
	serverPath string,
	imageTags []string,
	buildStatus *chan whaleAquarium.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId string) {

	var containersVolumes []mount.Mount
	var imageVolumesList []string
	var containersVolumeTmpList = make([]whaleAquarium.Mount, 0)

	// init docker
	var dockerSys = whaleAquarium.DockerSystem{}
	err = dockerSys.Init()
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
			whaleAquarium.Mount{
				MountType:   whaleAquarium.KVolumeMountTypeBind,
				Source:      volumePathInsideImage,
				Destination: volumePathInsideImage,
			},
		)
	}

	err, containersVolumes = factoryWhaleAquarium.NewVolumeMount(containersVolumeTmpList)
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateAndStart(
		newImageName,
		newContainerName,
		whaleAquarium.KRestartPolicyUnlessStopped,
		containersVolumes,
		nil,
	)

	return
}

func NewContainerFromRemoteServerWithNetworkConfiguration(
	newImageName,
	newContainerName,
	networkName string,
	networkDrive whaleAquarium.NetworkDrive,
	networkScope,
	networkSubnet,
	networkGateway string,
	serverPath string,
	imageTags []string,
	buildStatus *chan whaleAquarium.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId, networkId string) {

	var containersVolumes []mount.Mount
	var imageVolumesList []string
	var containersVolumeTmpList = make([]whaleAquarium.Mount, 0)
	var networkConfig whaleAquarium.NextNetworkConfiguration

	// init docker
	var dockerSys = whaleAquarium.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, networkId, networkConfig = dockerSys.NetworkCreate(
		networkName,
		whaleAquarium.KNetworkDriveBridge,
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
			whaleAquarium.Mount{
				MountType:   whaleAquarium.KVolumeMountTypeBind,
				Source:      volumePathInsideImage,
				Destination: volumePathInsideImage,
			},
		)
	}

	err, containersVolumes = factoryWhaleAquarium.NewVolumeMount(containersVolumeTmpList)
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateAndStart(
		newImageName,
		newContainerName,
		whaleAquarium.KRestartPolicyUnlessStopped,
		containersVolumes,
		networkConfig,
	)

	return
}
