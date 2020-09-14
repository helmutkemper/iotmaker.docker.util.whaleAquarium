package factoryContainerMongoDB

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"

func NewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort(
	containerName string,
	newContainerRestartPolicy iotmakerdocker.RestartPolicy,
	networkAutoConfiguration *iotmakerdocker.NextNetworkAutoConfiguration,
	version MongoDBVersionTag,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
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
