package toolsGarbageCollector

import (
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

// Remove unreferenced volumes
func VolumesUnreferencedRemove() (err error) {
	var dockerSys = iotmakerDocker.DockerSystem{}
	var volumes []types.Volume
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, volumes = dockerSys.VolumeList()
	if err != nil {
		return
	}

	for _, volumeData := range volumes {
		if volumeData.UsageData == nil || volumeData.UsageData.RefCount == -1 {
			// bug: volume do portainer tem volumeData.UsageData = nil
			_ = dockerSys.VolumeRemove(volumeData.Name)
		}
	}

	return
}
