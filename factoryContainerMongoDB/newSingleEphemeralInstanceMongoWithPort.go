package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/util"
)

func NewSingleEphemeralInstanceMongoWithPort(containerName string, networkUtil util.NetworkGenerator, port nat.Port, version MongoDBVersionTag, pullStatus chan whaleAquarium.ContainerPullStatusSendToChannel) (error, string) {
	var imageName = "mongo:" + version.String()

	return newMongoEphemeral(containerName, imageName, networkUtil, port, pullStatus)
}
