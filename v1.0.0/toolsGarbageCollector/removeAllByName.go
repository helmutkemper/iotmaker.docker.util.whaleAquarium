package toolsGarbageCollector

import (
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

// Use this function to remove trash after test.
// This function removes container, image and network by name and unlinked volumes and
// imagens
func RemoveAllByNameContains(name string) (err error) {
	var nameAndId []iotmakerDocker.NameAndId
	var container types.ContainerJSON
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, nameAndId = dockerSys.ContainerFindIdByNameContains(name)
	if err != nil && err.Error() != "container name not found" {
		return err
	}

	for _, data := range nameAndId {
		err, container = dockerSys.ContainerInspect(data.ID)
		if err != nil {
			return
		}

		if container.State != nil && container.State.Running == true {
			err = dockerSys.ContainerStopAndRemove(data.ID, true, false, false)
			if err != nil {
				return
			}
		}

		if container.State != nil && container.State.Running == false {
			err = dockerSys.ContainerRemove(data.ID, true, false, false)
			if err != nil {
				return
			}
		}
	}

	err, nameAndId = dockerSys.ImageFindIdByNameContains(name)
	if err != nil && err.Error() != "image name not found" {
		return err
	}
	for _, data := range nameAndId {
		err = dockerSys.ImageRemove(data.ID, false, false)
		if err != nil {
			return
		}
	}

	err, nameAndId = dockerSys.NetworkFindIdByNameContains(name)
	if err != nil && err.Error() != "network name not found" {
		return err
	}
	for _, data := range nameAndId {
		err = dockerSys.NetworkRemove(data.ID)
		if err != nil {
			return
		}
	}

	err = VolumesUnreferencedRemove()
	if err != nil {
		return
	}

	err = dockerSys.ImageGarbageCollector()
	if err != nil {
		return
	}

	return
}
