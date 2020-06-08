package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/util"
)

func NewSingleEphemeralInstanceMongoWithPort(containerName string, networkUtil util.NetworkGenerator, port nat.Port, version MongoDBVersionTag) (error, string) {
	var imageName = "mongo:" + version.String()

	return newMongoEphemeral(containerName, imageName, networkUtil, port)
}
