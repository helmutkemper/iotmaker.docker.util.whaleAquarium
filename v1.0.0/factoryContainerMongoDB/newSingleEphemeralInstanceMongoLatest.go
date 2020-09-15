package factorycontainermongodb

import (
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// NewSingleEphemeralInstanceMongoLatest (English): Install a last version of MongoDB
// with ephemeral data
//
// NewSingleEphemeralInstanceMongoLatest (Português): Instala a última versão do MongoDB
// com dados efêmeros
func NewSingleEphemeralInstanceMongoLatest(
	containerName string,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (
	containerId string,
	err error,
) {

	var imageName = "mongo:latest"
	port, _ := nat.NewPort("tcp", "27017")
	containerId, err = newMongoEphemeral(imageName, containerName, port, pullStatus)

	return
}
