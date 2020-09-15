package factoryContainerMongoDB

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"

// NewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort (English):
// Install a MongoDB with ephemeral data, network configuration and don't expose MongoDB
// port
//
// NewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort (Português):
// Instala o MongoDB com dado efêmeos, configuração de rede e não expõe a porta do
// MongoDB
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
