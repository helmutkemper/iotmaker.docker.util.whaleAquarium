package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongoWithPortAndNetworkConfiguration(
	containerName string,
	newContainerRestartPolicy iotmakerDocker.RestartPolicy,
	networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration,
	port nat.Port,
	version MongoDBVersionTag,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId, networkId string) {

	var imageName = "mongo:" + version.String()
	err, containerId, networkId = newMongoEphemeralWithNetworkConfiguration(
		imageName,
		containerName,
		newContainerRestartPolicy,
		networkAutoConfiguration,
		port,
		pullStatus,
	)

	return
}
