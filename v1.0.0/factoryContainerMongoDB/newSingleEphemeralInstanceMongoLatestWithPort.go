package factory_container_mongodb

import (
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

// NewSingleEphemeralInstanceMongoLatestWithPort (English): Install a MongoDB with
// ephemeral data and change exposed MongoDB port
//
// NewSingleEphemeralInstanceMongoLatestWithPort (Português): Instala o MongoDB com dados
// efêmeos e muda a porta do MongoDB exporta
func NewSingleEphemeralInstanceMongoLatestWithPort(
	containerName string,
	port nat.Port,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (err error, containerId string) {

	var imageName = "mongo:latest"
	err, containerId = newMongoEphemeral(imageName, containerName, port, pullStatus)

	return
}
