package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/util"
)

func NewSingleEphemeralInstanceMongoLatestWithPort(containerName string, networkUtil util.NetworkGenerator, port nat.Port) (error, string) {
	var imageName = "mongo:latest"

	return newMongoEphemeral(containerName, imageName, networkUtil, port)
}
