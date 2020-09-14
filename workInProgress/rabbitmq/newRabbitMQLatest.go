package rabbitmq

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func NewRabbitMQLatest(
	containerName string,
	containerRestartPolicy iotmakerdocker.RestartPolicy,
	networkAutoConfiguration *iotmakerdocker.NextNetworkAutoConfiguration,
	portHttpAPI nat.Port,
	portAMPQ nat.Port,
	portInterNodeAndCli nat.Port,
	version RabbitMQVersionTag,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
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
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	// image pull and wait
	_, _, err = dockerSys.ImagePull(imageName, pullStatus)
	if err != nil {
		return
	}

	newPortMap := nat.PortMap{
		// container port number/protocol [tpc/udp]
		currentPortHttpAPI: []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: currentPortHttpAPI.Port(),
			},
		},
		currentPortAMPQ: []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: currentPortAMPQ.Port(),
			},
		},
		currentPortInterNodeAndCli: []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: currentPortInterNodeAndCli.Port(),
			},
		},
	}

	err, networkConfig = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	containerId, err = dockerSys.ContainerCreateAndStart(
		imageName,
		containerName,
		containerRestartPolicy,
		newPortMap,
		mountList,
		networkConfig,
	)

	return
}
