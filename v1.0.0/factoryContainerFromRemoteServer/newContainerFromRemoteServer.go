package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

// English: Create a image and create and start a container from project inside into server
func NewContainerFromRemoteServer(
	newImageName,
	newContainerName,
	serverPath string,
	imageTags []string,
	buildStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, imageId, containerId string) {

	var containersVolumes []mount.Mount
	var imageVolumesList []string
	var containersVolumeTmpList = make([]iotmakerDocker.Mount, 0)

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

	err, containersVolumes = factoryDocker.NewVolumeMount(containersVolumeTmpList)
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateAndStart(
		newImageName,
		newContainerName,
		iotmakerDocker.KRestartPolicyUnlessStopped,
		containersVolumes,
		nil,
	)

	return
}
