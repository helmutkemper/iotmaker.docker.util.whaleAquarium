package factoryContainerFromRemoteServer

import (
	"github.com/docker/docker/api/types/mount"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// English: Create a image and create and start a container from project inside into server
func NewContainerCreateExposePortsAutomaticallyAndStart(
	newImageName,
	newContainerName,
	serverPath string,
	imageTags []string,
	buildStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (
	imageId string,
	containerId string,
	err error,
) {

	var containersVolumes []mount.Mount
	var imageVolumesList []string
	var containersVolumeTmpList = make([]iotmakerdocker.Mount, 0)

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

	containerId, err = dockerSys.ContainerCreateExposePortsAutomaticallyAndStart(
		newImageName,
		newContainerName,
		iotmakerdocker.KRestartPolicyUnlessStopped,
		containersVolumes,
		nil,
	)

	return
}
