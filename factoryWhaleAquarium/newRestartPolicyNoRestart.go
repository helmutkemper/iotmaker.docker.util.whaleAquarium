package factoryWhaleAquarium

import whaleAquarium "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium"

func NewRestartPolicyNoRestart() whaleAquarium.RestartPolicy {
	return whaleAquarium.KRestartPolicyNo
}
