package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongoWithPort(
	containerName string,
	port nat.Port,
	version MongoDBVersionTag,
	pullStatus *chan whaleAquarium.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:" + version.String()
	err, containerId = newMongoEphemeral(containerName, imageName, port, pullStatus)

	return
}
