package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func NewSingleEphemeralInstanceMongoLatestWithPort(
	containerName string,
	port nat.Port,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:latest"
	err, containerId = newMongoEphemeral(imageName, containerName, port, pullStatus)

	return
}
