package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongoLatestWithPort(
	containerName string,
	port nat.Port,
	pullStatus *chan whaleAquarium.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:latest"
	err, containerId = newMongoEphemeral(containerName, imageName, port, pullStatus)

	return
}
