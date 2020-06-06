package factoryWhaleAquarium

import "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/util"

func NewContainerNetworkGenerator(name string, a, b, c, d byte) util.NetworkGenerator {
	var ret = util.NetworkGenerator{}
	ret.Init(name, a, b, c, d)

	return ret
}
