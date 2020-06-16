package factoryWhaleAquarium

import (
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/util"
)

func NewContainerNetworkGenerator(name string, a, b, c, d byte) (err error, generator util.NetworkGenerator) {
	var exists bool
	var networkId string
	var net = whaleAquarium.DockerSystem{}
	net.Init()

	err, exists = net.NetworkVerifyName(name)
	if err != nil {
		return
	}

	if exists == false {
		err, networkId = net.NetworkCreate(name)
		if err != nil {
			return
		}
	} else {
		err, networkId = net.NetworkFindIdByName(name)
		if err != nil {
			return
		}
	}

	generator.Init(networkId, name, a, b, c, d)

	return
}
