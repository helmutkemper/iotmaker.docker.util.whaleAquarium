package factoryWhaleAquarium

import (
	whaleAquarium "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium"
)

func NewRestartPolicyOnFailureRestart() whaleAquarium.RestartPolicy {
	return whaleAquarium.KRestartPolicyOnFailure
}
