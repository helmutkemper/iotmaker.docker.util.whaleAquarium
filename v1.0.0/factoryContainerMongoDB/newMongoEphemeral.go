package factorycontainermongodb

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.db.mongodb.config/factoryMongoDBConfig"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"io/ioutil"
	"os"
)

// newMongoEphemeral (English):
//
// newMongoEphemeral (Português):
func newMongoEphemeral(
	imageName,
	containerName string,
	newPort nat.Port,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (
	containerId string,
	err error,
) {

	var file []byte
	var mountList []mount.Mount
	var defaultMongoDbPort nat.Port

	defaultMongoDbPort, err = nat.NewPort("tcp", "27017")
	if err != nil {
		return
	}

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

	// image pull and wait (true)
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
		defaultMongoDbPort: []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: newPort.Port(),
			},
		},
	}

	containerId, err = dockerSys.ContainerCreateAndStart(
		imageName,
		containerName,
		iotmakerdocker.KRestartPolicyUnlessStopped,
		portMap,
		mountList,
		nil,
	)

	return
}
