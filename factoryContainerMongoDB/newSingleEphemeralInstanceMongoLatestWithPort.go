package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

func NewSingleEphemeralInstanceMongoLatestWithPort(containerName, networkName string, port nat.Port, pullStatus chan whaleAquarium.ContainerPullStatusSendToChannel) (error, string) {
	var imageName = "mongo:latest"

	err, netGenerator := factoryDocker.NewContainerNetworkGenerator(networkName, 10, 0, 0, 1)
	if err != nil {
		return err, ""
	}

	return newMongoEphemeral(containerName, imageName, netGenerator, port, pullStatus)
}
