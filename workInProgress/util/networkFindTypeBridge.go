package util

import (
	"errors"
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func NetworkFindTypeBridge() (err error, inspectList []types.NetworkResource) {
	var list []types.NetworkResource
	var inspect types.NetworkResource
	var netDriveToFind = iotmakerdocker.KNetworkDriveBridge

	inspectList = make([]types.NetworkResource, 0)

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
			inspect, err = ds.NetworkInspect(net.ID)
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
