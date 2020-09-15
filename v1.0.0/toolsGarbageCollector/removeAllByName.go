package toolsGarbageCollector

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// RemoveAllByNameContains (English): Use this function to remove unnecessary after
// testing.
// This function removes container, image and network by name and unlinked volumes and
// imagens
//
// RemoveAllByNameContains (Português): Use esta função para remover o desnecessário
// depois dos testes.
// Esta função remove containers, imagens e rede por nome e volumes e imagens sem uso.
func RemoveAllByNameContains(name string) (err error) {
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err = dockerSys.RemoveAllByNameContains(name)

	return
}
