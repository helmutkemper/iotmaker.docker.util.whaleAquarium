package util

import (
	"github.com/docker/docker/api/types/network"
)

type NetworkGenerator struct {
	ip   IPv4Generator
	name string
}

func (el *NetworkGenerator) Init(name string, a, b, c, d byte) {
	el.name = name
	el.ip.Init(a, b, c, d)
}

func (el *NetworkGenerator) GetNext() (error, *network.NetworkingConfig) {
	var err = el.ip.Inc()
	return err, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			el.name: {
				IPAddress: el.ip.String(),
			},
		},
	}
}
