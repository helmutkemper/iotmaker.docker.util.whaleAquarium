package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongoLatest(
	containerName string,
	pullStatus *chan whaleAquarium.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:latest"
	port, _ := nat.NewPort("tcp", "27017")
	err, containerId = newMongoEphemeral(imageName, containerName, port, pullStatus)

	return
}
