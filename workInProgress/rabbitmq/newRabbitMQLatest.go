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
	var currentPortHttpAPI nat.Port
	var currentPortAMPQ nat.Port
	var currentPortInterNodeAndCli nat.Port

	var imageName = "rabbitmq:" + version.String()
	currentPortHttpAPI, err = nat.NewPort("tcp", "15672")
	currentPortAMPQ, err = nat.NewPort("tcp", "5672")
	currentPortInterNodeAndCli, err = nat.NewPort("tcp", "25676")

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

	currentPortList := []nat.Port{
		currentPortHttpAPI,
		currentPortAMPQ,
		currentPortInterNodeAndCli,
	}

	newPortList := []nat.Port{
		portHttpAPI,
		portAMPQ,
		portInterNodeAndCli,
	}

	err, networkConfig = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	err, containerId = dockerSys.ContainerCreateChangeExposedPortAndStart(
		imageName,
		containerName,
		containerRestartPolicy,
		mountList,
		networkConfig,
		currentPortList,
		newPortList,
	)

	return
}
