package factorycontainermongodb

import (
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// NewSingleEphemeralInstanceMongo (English): Install mongodb with the version of your
// choice with ephemeral data
//
// NewSingleEphemeralInstanceMongo (Português): Instala o MongoDB com a versão de sua
// escolha e dados efêmeros
func NewSingleEphemeralInstanceMongo(
	containerName string,
	version MongoDBVersionTag,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:" + version.String()
	port, _ := nat.NewPort("tcp", "27017")
	err, containerId = newMongoEphemeral(imageName, containerName, port, pullStatus)

	return
}
