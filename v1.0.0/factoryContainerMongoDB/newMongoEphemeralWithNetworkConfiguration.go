package factory_container_mongodb

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.db.mongodb.config/factoryMongoDBConfig"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"io/ioutil"
	"os"
)

// newMongoEphemeralWithNetworkConfiguration (English):
//
// newMongoEphemeralWithNetworkConfiguration (Português):
func newMongoEphemeralWithNetworkConfiguration(
	imageName,
	containerName string,
	containerRestartPolicy iotmakerdocker.RestartPolicy,
	networkAutoConfiguration *iotmakerdocker.NextNetworkAutoConfiguration,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
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
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, networkConfig = networkAutoConfiguration.GetNext()
	if err != nil {
		return
	}

	// image pull and wait
	_, _, err = dockerSys.ImagePull(imageName, pullStatus)
	if err != nil {
		return
	}

	// define an external MongoDB config file path
	mountList, err = iotmakerdocker.NewVolumeMount(
		[]iotmakerdocker.Mount{
			{
				MountType:   iotmakerdocker.KVolumeMountTypeBind,
				Source:      relativeConfigFilePathToSave,
				Destination: "/etc/mongo.conf",
			},
		},
	)
	if err != nil {
		return
	}

	portMap := nat.PortMap{
		// container port number/protocol [tpc/udp]
		"27017/tcp": []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: "27017",
			},
		},
	}

	containerId, err = dockerSys.ContainerCreateAndStart(
		imageName,
		containerName,
		containerRestartPolicy,
		portMap,
		mountList,
		networkConfig,
	)

	return
}
