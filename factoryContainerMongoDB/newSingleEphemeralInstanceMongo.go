package factoryContainerMongoDB

func NewSingleEphemeralInstanceMongo(containerName, networkName string, version MongoDBVersionTag) (error, string) {
	var imageName = "mongo:" + version.String()

	return newMongoEphemeral(containerName, networkName, imageName, 27017)
}
