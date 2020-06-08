package factoryContainerMongoDB

func NewSingleEphemeralInstanceMongoWithPort(containerName, networkName string, port int, version MongoDBVersionTag) (error, string) {
	var imageName = "mongo:" + version.String()

	return newMongoEphemeral(containerName, networkName, imageName, port)
}
