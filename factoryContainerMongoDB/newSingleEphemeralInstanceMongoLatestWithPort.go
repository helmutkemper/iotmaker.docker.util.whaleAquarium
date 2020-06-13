package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/util"
)

func NewSingleEphemeralInstanceMongoLatestWithPort(containerName string, networkUtil util.NetworkGenerator, port nat.Port, pullStatus chan whaleAquarium.ContainerPullStatusSendToChannel) (error, string) {
	var imageName = "mongo:latest"

	return newMongoEphemeral(containerName, imageName, networkUtil, port, pullStatus)
}
