package factoryContainerMongoDB

import "github.com/docker/go-connections/nat"

func NewSingleEphemeralInstanceMongoWithPort(containerName, networkName string, port nat.Port, version MongoDBVersionTag) (error, string) {
	var imageName = "mongo:" + version.String()

	return newMongoEphemeral(containerName, networkName, imageName, port)
}
