package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongo(
	containerName string,
	version MongoDBVersionTag,
	pullStatus *chan whaleAquarium.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:" + version.String()
	port, _ := nat.NewPort("tcp", "27017")
	err, containerId = newMongoEphemeral(containerName, imageName, port, pullStatus)

	return
}
