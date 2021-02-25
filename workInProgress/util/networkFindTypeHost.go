package util

import (
	"errors"
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func NetworkFindTypeHost() (
	inspect types.NetworkResource,
	err error,
) {
	var list []types.NetworkResource
	var netDriveToFind = iotmakerdocker.KNetworkDriveHost

	ds := iotmakerdocker.DockerSystem{}
	err = ds.Init()
	if err != nil {
		return
	}

	list, err = ds.NetworkList()
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
