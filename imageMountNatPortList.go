package iotmaker_docker_util_whaleAquarium

import "github.com/docker/go-connections/nat"

// Mount nat por list by image config
func (el *DockerSystem) ImageMountNatPortList(imageId string) (error, nat.PortMap) {
	var err error
	var portList []string
	var ret nat.PortMap = make(map[nat.Port][]nat.PortBinding)

	err, portList = el.ImageListExposedPorts(imageId)
	if err != nil {
		return err, nat.PortMap{}
	}

	for _, port := range portList {
		ret[nat.Port(port)] = []nat.PortBinding{
			{
				HostPort: port,
			},
		}
	}

	return err, ret
}
