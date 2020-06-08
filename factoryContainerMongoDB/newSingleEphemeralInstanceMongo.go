package factoryContainerMongoDB

import "github.com/docker/go-connections/nat"

func NewSingleEphemeralInstanceMongo(containerName, networkName string, version MongoDBVersionTag) (error, string) {
	var imageName = "mongo:" + version.String()

	port, _ := nat.NewPort("tcp", "27017")
	return newMongoEphemeral(containerName, networkName, imageName, port)
}
