package toolsgarbagecollector

import (
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// Remove unreferenced volumes
func VolumesUnreferencedRemove() (err error) {
	var dockerSys = iotmakerdocker.DockerSystem{}
	var volumes []types.Volume
	err = dockerSys.Init()
	if err != nil {
		return
	}

	volumes, err = dockerSys.VolumeList()
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
