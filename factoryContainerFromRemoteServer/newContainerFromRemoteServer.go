package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"github.com/helmutkemper/iotmaker.docker/util"
)

// English: Create a image and create and start a container from project inside into server
// Warning: work in progress - buildStatus don't work yet
func NewContainerFromRemoteServer(
	newImageName,
	newContainerName,
	networkName,
	serverPath string,
	imageTags []string,
	buildStatus chan whaleAquarium.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId string) {

	var nextNetworkConfig *network.NetworkingConfig
	var networkUtil util.NetworkGenerator
	var containersVolumes []mount.Mount
	var imageVolumesList []string
	var containersVolumeTmpList = make([]whaleAquarium.Mount, 0)

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
