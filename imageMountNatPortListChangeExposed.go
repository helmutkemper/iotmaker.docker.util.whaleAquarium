package iotmaker_docker_util_whaleAquarium

import (
	"github.com/docker/go-connections/nat"
)

// Mount nat por list by image config
func (el *DockerSystem) ImageMountNatPortListChangeExposed(imageId string, currentPortList, changeToPortList []nat.Port) (error, nat.PortMap) {
	var err error
	var portList []string
	var ret nat.PortMap = make(map[nat.Port][]nat.PortBinding)

	err, portList = el.ImageListExposedPorts(imageId)
	if err != nil {
		return err, nat.PortMap{}
	}

	for _, port := range portList {
		inPort := port
		for k, currPort := range currentPortList {
			if currPort.Port()+"/"+currPort.Proto() == port {
				inPort = changeToPortList[k].Port() + "/" + changeToPortList[k].Proto()
				break
			}
		}

		ret[nat.Port(port)] = []nat.PortBinding{
			{
				HostPort: inPort,
			},
		}
	}

	return err, ret
}
