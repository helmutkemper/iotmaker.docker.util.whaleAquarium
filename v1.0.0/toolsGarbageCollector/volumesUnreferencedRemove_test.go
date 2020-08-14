package toolsGarbageCollector

import (
	"fmt"
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func ExampleVolumesUnreferencedRemove() {
	var err error
	var volumes []types.Volume
	var counterUnreferencedVolumes = 0

	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(err)
	}

	err, volumes = dockerSys.VolumeList()
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
	err, volumes = dockerSys.VolumeList()
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
