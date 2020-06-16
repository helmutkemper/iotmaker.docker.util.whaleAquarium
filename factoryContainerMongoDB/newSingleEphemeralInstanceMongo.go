package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongo(
	containerName string,
	version MongoDBVersionTag,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:" + version.String()
	port, _ := nat.NewPort("tcp", "27017")
	err, containerId = newMongoEphemeral(containerName, imageName, port, pullStatus)

	return
}
