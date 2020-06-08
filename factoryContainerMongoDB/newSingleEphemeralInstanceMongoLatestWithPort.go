package factoryContainerMongoDB

import "github.com/docker/go-connections/nat"

func NewSingleEphemeralInstanceMongoLatestWithPort(containerName, networkName string, port nat.Port) (error, string) {
	var imageName = "mongo:latest"

	return newMongoEphemeral(containerName, networkName, imageName, port)
}
