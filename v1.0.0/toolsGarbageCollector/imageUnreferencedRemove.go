package toolsGarbageCollector

import (
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func ImageUnreferencedRemove() (err error) {
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err = dockerSys.ImageGarbageCollector()
	return
}
