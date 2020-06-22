package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongoWithPort(
	containerName string,
	port nat.Port,
	version MongoDBVersionTag,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:" + version.String()
	err, containerId = newMongoEphemeral(imageName, containerName, port, pullStatus)

	return
}
