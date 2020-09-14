package toolsGarbageCollector

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"

func ImageUnreferencedRemove() (err error) {
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err = dockerSys.ImageGarbageCollector()
	return
}
