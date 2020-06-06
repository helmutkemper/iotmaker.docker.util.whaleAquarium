package factoryWhaleAquarium

import (
	whaleAquarium "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/util"
)

func NewContainerNetworkGenerator(name string, a, b, c, d byte) (error, util.NetworkGenerator) {
	var err error
	var exists bool
	var net = whaleAquarium.Docker{}
	var ret = util.NetworkGenerator{}
	ret.Init(name, a, b, c, d)

	err, exists = net.NetworkVerifyName(name)
	if err != nil {
		return err, ret
	}

	if exists == false {
		err = net.NetworkCreate(name)
		if err != nil {
			return err, ret
		}
	}

	return err, ret
}
