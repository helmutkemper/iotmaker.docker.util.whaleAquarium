package factoryContainerMongoDB

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.db.mongodb.config/factoryMongoDBConfig"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
	"github.com/helmutkemper/iotmaker.docker/util"
	"io/ioutil"
	"os"
)

func newMongoEphemeral(containerName, imageName string, networkUtil util.NetworkGenerator, newPort nat.Port) (error, string) {
	var err error
	var id string
	var file []byte
	var mountList []mount.Mount
	var nextNetworkConfig *network.NetworkingConfig

	var relativeConfigFilePathToSave = "./config.conf"

	// basic MongoDB configuration
	var conf = factoryMongoDBConfig.NewBasicConfigWithEphemeralData()
	err, file = conf.ToYaml(0)
	if err != nil {
		return err, ""
	}

	// save MongoDB configuration into disk
	err = ioutil.WriteFile(relativeConfigFilePathToSave, file, os.ModePerm)
	if err != nil {
		return err, ""
	}

	// init docker
	var dockerSys = whaleAquarium.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return err, ""
	}

	// image pull and wait (true)
	err = dockerSys.ImagePull(imageName, true)
	if err != nil {
		return err, ""
	}

	// define an external MongoDB config file path
	err, mountList = factoryWhaleAquarium.NewVolumeMount(
		[]whaleAquarium.Mount{
			{
				MountType:   whaleAquarium.KVolumeMountTypeBind,
				Source:      relativeConfigFilePathToSave,
				Destination: "/etc/mongo.conf",
			},
		},
	)
	if err != nil {
		return err, ""
	}

	err, nextNetworkConfig = networkUtil.GetNext()
	if err != nil {
		return err, ""
	}

	currentPort, _ := nat.NewPort("tcp", "27017")
	currentPortList := []nat.Port{
		currentPort,
	}

	newPortList := []nat.Port{
		newPort,
	}

	err, id = dockerSys.ContainerCreateChangeExposedPortAndStart(
		imageName,
		containerName,
		whaleAquarium.KRestartPolicyUnlessStopped,
		mountList,
		nextNetworkConfig,
		currentPortList,
		newPortList,
	)

	return err, id
}
