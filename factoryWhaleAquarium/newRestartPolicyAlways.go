package factoryWhaleAquarium

import whaleAquarium "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium"

func NewKRestartPolicyAlwaysRestart() whaleAquarium.RestartPolicy {
	return whaleAquarium.KRestartPolicyOnFailure
}
