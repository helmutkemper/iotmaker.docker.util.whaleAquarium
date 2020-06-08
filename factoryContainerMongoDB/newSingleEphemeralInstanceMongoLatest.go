package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/util"
)

func NewSingleEphemeralInstanceMongoLatest(containerName string, networkUtil util.NetworkGenerator) (error, string) {
	var imageName = "mongo:latest"

	port, _ := nat.NewPort("tcp", "27017")
	return newMongoEphemeral(containerName, imageName, networkUtil, port)
}
