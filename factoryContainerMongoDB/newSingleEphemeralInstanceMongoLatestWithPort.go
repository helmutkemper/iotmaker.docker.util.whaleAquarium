package factoryContainerMongoDB

func NewSingleEphemeralInstanceMongoLatestWithPort(containerName, networkName string, port int) (error, string) {
	var imageName = "mongo:latest"

	return newMongoEphemeral(containerName, networkName, imageName, port)
}
