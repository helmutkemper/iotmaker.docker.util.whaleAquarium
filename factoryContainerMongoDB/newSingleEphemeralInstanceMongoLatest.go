package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/util"
)

func NewSingleEphemeralInstanceMongoLatest(containerName string, networkUtil util.NetworkGenerator, pullStatus chan whaleAquarium.ContainerPullStatusSendToChannel) (error, string) {
	var imageName = "mongo:latest"

	port, _ := nat.NewPort("tcp", "27017")
	return newMongoEphemeral(containerName, imageName, networkUtil, port, pullStatus)
}
