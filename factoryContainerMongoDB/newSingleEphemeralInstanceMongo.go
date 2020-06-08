package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.docker/util"
)

func NewSingleEphemeralInstanceMongo(containerName string, networkUtil util.NetworkGenerator, version MongoDBVersionTag) (error, string) {
	var imageName = "mongo:" + version.String()

	port, _ := nat.NewPort("tcp", "27017")
	return newMongoEphemeral(containerName, imageName, networkUtil, port)
}
