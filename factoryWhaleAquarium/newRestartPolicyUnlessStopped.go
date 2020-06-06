package factoryWhaleAquarium

import (
	whaleAquarium "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium"
)

func NewRestartPolicyUnlessStopped() whaleAquarium.RestartPolicy {
	return whaleAquarium.KRestartPolicyUnlessStopped
}
