package toolsGarbageCollector

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// Use this function to remove trash after test.
// This function removes container, image and network by name and unlinked volumes and
// imagens
func RemoveAllByNameContains(name string) (err error) {
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err = dockerSys.RemoveAllByNameContains(name)

	return
}
