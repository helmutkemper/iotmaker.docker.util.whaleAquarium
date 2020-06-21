package util

import (
	"errors"
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NetworkFindTypeHost() (err error, inspect types.NetworkResource) {
	var list []types.NetworkResource
	var netDriveToFind = iotmakerDocker.KNetworkDriveHost

	ds := iotmakerDocker.DockerSystem{}
	err = ds.Init()
	if err != nil {
		return
	}

	err, list = ds.NetworkList()
	if err != nil {
		return
	}

	for _, net := range list {
		if net.Driver == netDriveToFind.String() {
			return ds.NetworkInspect(net.ID)
		}
	}

	err = errors.New("network type host not found")

	return
}
