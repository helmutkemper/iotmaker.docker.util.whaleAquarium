package factoryContainerMongoDB

func NewSingleEphemeralInstanceMongoLatest(containerName, networkName string) (error, string) {
	var imageName = "mongo:latest"

	return newMongoEphemeral(containerName, networkName, imageName, 27017)
}
