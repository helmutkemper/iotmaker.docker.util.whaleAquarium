package toolsgarbagecollector

import (
	"fmt"
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func ExampleVolumesUnreferencedRemove() {
	var err error
	var volumes []types.Volume
	var counterUnreferencedVolumes = 0

	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(err)
	}

	volumes, err = dockerSys.VolumeList()
	if err != nil {
		panic(err)
	}

	for _, volumeData := range volumes {
		if volumeData.UsageData == nil || volumeData.UsageData.RefCount == -1 {
			counterUnreferencedVolumes += 1
		}
	}

	if counterUnreferencedVolumes == 0 {
		return
	}

	err = VolumesUnreferencedRemove()
	if err != nil {
		panic(err)
	}

	counterUnreferencedVolumes = 0
	volumes, err = dockerSys.VolumeList()
	if err != nil {
		panic(err)
	}

	for _, volumeData := range volumes {
		if volumeData.UsageData == nil || volumeData.UsageData.RefCount == -1 {
			counterUnreferencedVolumes += 1
		}
	}

	if counterUnreferencedVolumes != 0 {
		fmt.Println("test fail")
	}

	// Output:
	//
}
