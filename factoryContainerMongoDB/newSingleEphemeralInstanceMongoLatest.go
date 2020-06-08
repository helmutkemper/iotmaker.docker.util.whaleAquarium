package factoryContainerMongoDB

import "github.com/docker/go-connections/nat"

func NewSingleEphemeralInstanceMongoLatest(containerName, networkName string) (error, string) {
	var imageName = "mongo:latest"

	port, _ := nat.NewPort("tcp", "27017")
	return newMongoEphemeral(containerName, networkName, imageName, port)
}
