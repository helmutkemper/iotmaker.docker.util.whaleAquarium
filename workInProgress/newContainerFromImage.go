package workInProgress

import (
	"github.com/docker/docker/api/types/mount"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

func NewContainerFromImage(
	imageName,
	containerName string,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var containersVolumes []mount.Mount
	var imageVolumesList []string
	var containersVolumeTmpList = make([]iotmakerDocker.Mount, 0)

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	// image pull and wait (true)
	err, _, _ = dockerSys.ImagePull(imageName, pullStatus)
	if err != nil {
		return
	}

	err, imageVolumesList = dockerSys.ImageListExposedVolumesByName(imageName)
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
		imageName,
		containerName,
		iotmakerDocker.KRestartPolicyUnlessStopped,
		containersVolumes,
		nil,
	)

	return
}
