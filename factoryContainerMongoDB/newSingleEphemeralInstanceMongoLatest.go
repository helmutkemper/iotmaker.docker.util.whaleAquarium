package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

func NewSingleEphemeralInstanceMongoLatest(containerName, networkName string, pullStatus *chan whaleAquarium.ContainerPullStatusSendToChannel) (error, string) {
	var imageName = "mongo:latest"

	err, netGenerator, _ := factoryDocker.NewContainerNetworkGenerator(networkName, 10, 0, 0, 1)
	if err != nil {
		return err, ""
	}

	port, _ := nat.NewPort("tcp", "27017")
	return newMongoEphemeral(containerName, imageName, netGenerator, port, pullStatus)
}
