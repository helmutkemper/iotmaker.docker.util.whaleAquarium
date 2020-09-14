package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func NewSingleEphemeralInstanceMongoWithPort(
	containerName string,
	port nat.Port,
	version MongoDBVersionTag,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:" + version.String()
	err, containerId = newMongoEphemeral(imageName, containerName, port, pullStatus)

	return
}
