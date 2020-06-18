package rabbitmq

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewRabbitMQLatest(
	containerName string,
	containerRestartPolicy iotmakerDocker.RestartPolicy,
	networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration,
	portHttpAPI nat.Port,
	portAMPQ nat.Port,
	portInterNodeAndCli nat.Port,
	version RabbitMQVersionTag,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId, networkId string) {

	var mountList []mount.Mount
	var networkConfig *network.NetworkingConfig

	var imageName = "rabbitmq:" + version.String()

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	// image pull and wait
	err, _, _ = dockerSys.ImagePull(imageName, pullStatus)
	if err != nil {
		return
	}

	newPortList := nat.PortMap{
		portHttpAPI: []nat.PortBinding{
			{
				HostPort: portHttpAPI.Port(),
			},
		},
		portAMPQ: []nat.PortBinding{
			{
				HostPort: portAMPQ.Port(),
			},
		},
		portInterNodeAndCli: []nat.PortBinding{
			{
				HostPort: portInterNodeAndCli.Port(),
			},
		},
	}

	err, networkConfig = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateExposePortAndStart(
		imageName,
		containerName,
		containerRestartPolicy,
		mountList,
		networkConfig,
		newPortList,
	)

	return
}
