package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongoLatestWithPort(
	containerName string,
	port nat.Port,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:latest"
	err, containerId = newMongoEphemeral(imageName, containerName, port, pullStatus)

	return
}
