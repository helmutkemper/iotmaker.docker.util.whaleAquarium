package factoryWhaleAquarium

import "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/util"

func NewIPv4Generator(a, b, c, d byte) util.IPv4Generator {
	var ret util.IPv4Generator
	ret.Init(a, b, c, d)

	return ret
}
