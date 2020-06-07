package factoryContainerMongoDB

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/helmutkemper/iotmaker.db.mongodb.config/factoryMongoDBConfig"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/util"
	"io/ioutil"
	"os"
)

func NewSingleEphemeralInstanceMongoLatest(containerName, networkName string) (error, string) {
	var imageName = "mongo:latest"

	return newMongoEphemeral(containerName, networkName, imageName)
}

func NewSingleEphemeralInstanceMongo(containerName, networkName string, version MongoDBVersionTag) (error, string) {
	var imageName = "mongo:" + version.String()

	return newMongoEphemeral(containerName, networkName, imageName)
}

func newMongoEphemeral(containerName, networkName, imageName string) (error, string) {
	var err error
	var id string
	var file []byte
	var networkUtil util.NetworkGenerator
	var mountList []mount.Mount
	var nextNetworkConfig *network.NetworkingConfig

	var relativeConfigFilePathToSave = "./config.conf"

	// basic MongoDB configuration
	var conf = factoryMongoDBConfig.NewBasicConfig()
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

	// create a network
	err, networkUtil = factoryWhaleAquarium.NewContainerNetworkGenerator(networkName, 10, 0, 0, 1)
	if err != nil {
		return err, ""
	}

	// image pull and wait (true)
	err = dockerSys.ImagePull(imageName, true)
	if err != nil {
		return err, ""
	}

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

	err, id = dockerSys.ContainerCreateAndStart(
		imageName,
		containerName,
		whaleAquarium.KRestartPolicyUnlessStopped,
		mountList,
		nextNetworkConfig,
	)

	return err, id
}
