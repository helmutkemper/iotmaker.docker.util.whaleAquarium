package util

import (
	"errors"
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NetworkFindTypeBridge() (err error, inspectList []types.NetworkResource) {
	var list []types.NetworkResource
	var inspect types.NetworkResource
	var netDriveToFind = iotmakerDocker.KNetworkDriveBridge

	inspectList = make([]types.NetworkResource, 0)

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
			err, inspect = ds.NetworkInspect(net.ID)
			if err != nil {
				return
			}
			inspectList = append(inspectList, inspect)
		}
	}

	if len(inspectList) == 0 {
		err = errors.New("network type bridge not found")
	}

	return
}
