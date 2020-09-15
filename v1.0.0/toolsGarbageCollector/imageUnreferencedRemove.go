package tools_garbage_collector

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"

// ImageUnreferencedRemove (English): Image garbage collector
//
// ImageUnreferencedRemove (Português): Coletor de lixo das imagens
func ImageUnreferencedRemove() (err error) {
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err = dockerSys.ImageGarbageCollector()
	return
}
