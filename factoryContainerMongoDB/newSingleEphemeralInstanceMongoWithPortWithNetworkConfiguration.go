package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongoWithPortWithNetworkConfiguration(
	containerName string,
	newContainerRestartPolicy iotmakerDocker.RestartPolicy,
	port nat.Port,
	version MongoDBVersionTag,
	networkName string,
	networkDrive iotmakerDocker.NetworkDrive,
	networkScope,
	networkSubnet,
	networkGateway string,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId, networkId string) {

	var imageName = "mongo:" + version.String()
	err, containerId, networkId = newMongoEphemeralWithNetworkConfiguration(
		imageName,
		containerName,
		newContainerRestartPolicy,
		networkName,
		networkDrive,
		networkScope,
		networkSubnet,
		networkGateway,
		port,
		pullStatus,
	)

	return
}
