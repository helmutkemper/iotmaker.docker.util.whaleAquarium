package factoryContainerMongoDB

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.db.mongodb.config/factoryMongoDBConfig"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"io/ioutil"
	"os"
)

func newMongoEphemeralWithNetworkConfiguration(
	imageName,
	containerName string,
	containerRestartPolicy iotmakerDocker.RestartPolicy,
	networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration,
	newPort nat.Port,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (err error, containerId, networkId string) {

	var file []byte
	var mountList []mount.Mount
	var networkConfig *network.NetworkingConfig

	var relativeConfigFilePathToSave = "./config.conf"

	// basic MongoDB configuration
	var conf = factoryMongoDBConfig.NewBasicConfigWithEphemeralData()
	err, file = conf.ToYaml(0)
	if err != nil {
		return
	}

	// save MongoDB configuration into disk
	err = ioutil.WriteFile(relativeConfigFilePathToSave, file, os.ModePerm)
	if err != nil {
		return
	}

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, networkConfig = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	// image pull and wait
	err, _, _ = dockerSys.ImagePull(imageName, pullStatus)
	if err != nil {
		return
	}

	// define an external MongoDB config file path
	err, mountList = factoryDocker.NewVolumeMount(
		[]iotmakerDocker.Mount{
			{
				MountType:   iotmakerDocker.KVolumeMountTypeBind,
				Source:      relativeConfigFilePathToSave,
				Destination: "/etc/mongo.conf",
			},
		},
	)
	if err != nil {
		return
	}

	currentPort, _ := nat.NewPort("tcp", "27017")
	currentPortList := []nat.Port{
		currentPort,
	}

	var IpAddress string
	err, IpAddress = networkAutoConfiguration.GetCurrentIpAddress()
	newPort, err = nat.NewPort(newPort.Proto(), IpAddress+":"+newPort.Port())
	newPortList := []nat.Port{
		newPort,
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
