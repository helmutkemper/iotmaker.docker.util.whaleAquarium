package factoryContainerMongoDB

import (
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort(
	containerName string,
	newContainerRestartPolicy iotmakerDocker.RestartPolicy,
	networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration,
	version MongoDBVersionTag,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId, networkId string) {

	var imageName = "mongo:" + version.String()
	err, containerId, networkId = newMongoEphemeralWithNetworkConfigurationAndWithoutPort(
		imageName,
		containerName,
		newContainerRestartPolicy,
		networkAutoConfiguration,
		pullStatus,
	)

	return
}
