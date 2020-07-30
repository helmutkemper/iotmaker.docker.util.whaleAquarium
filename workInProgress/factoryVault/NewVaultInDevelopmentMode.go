package factoryVault

import (
	"bytes"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"regexp"
)

func NewVaultInDevelopmentMode(
	containerName string,
	version VaultVersionTag,
	//newPort nat.Port,
	pullStatus *chan iotmakerDocker.ContainerPullStatusSendToChannel,
) (
	err error,
	containerId string,
	ApiAddress string,
	ClusterAddress string,
	vaultRootToken string,
	vaultUnsealKey string,
) {

	var ApiAddressRegExp *regexp.Regexp
	var ClusterAddressRegExp *regexp.Regexp
	var UnsealKeyRegExp *regexp.Regexp
	var RootTokenRegExp *regexp.Regexp
	var upgradingKeysFinishedRegExp *regexp.Regexp

	var imageName = "vault:" + version.String()
	//var mountList []mount.Mount
	var log []byte
	var logData [][]byte
	var tmp []byte

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, _, _ = dockerSys.ImagePull(imageName, pullStatus)
	if err != nil {
		return
	}

	// define an external MongoDB config file path
	//err, mountList = factoryDocker.NewVolumeMount(
	//  []iotmakerDocker.Mount{
	//    {
	//      MountType:   iotmakerDocker.KVolumeMountTypeBind,
	//      Source:      relativeConfigFilePathToSave,
	//      Destination: "/etc/mongo.conf",
	//    },
	//  },
	//)
	//if err != nil {
	//  return
	//}

	//defaultListenPort, _ := nat.NewPort("tcp", "8200")
	//unknownPort, _       := nat.NewPort("tcp", "8201")
	//currentPortList := []nat.Port{
	//  defaultListenPort,
	//  unknownPort,
	//}
	//
	//newPortList := []nat.Port{
	//  defaultListenPort,
	//  unknownPort,
	//}

	//err, containerId = dockerSys.ContainerCreateChangeExposedPortAndStart(
	//  imageName,
	//  containerName,
	//  iotmakerDocker.KRestartPolicyUnlessStopped,
	//  mountList,
	//  nil,
	//  currentPortList,
	//  newPortList,
	//)
	//if err != nil {
	//  return
	//}

	err, containerId = dockerSys.ContainerFindIdByName("vaultContainer")

	upgradingKeysFinishedRegExp = regexp.MustCompile("(?mi)upgrading keys finished") //todo: make const
	var pass = false
	for {
		err, log = dockerSys.ContainerLogs(containerId)
		if err != nil {
			return
		}

		logData = bytes.Split(log, []byte("\n"))
		for _, logLine := range logData {
			if upgradingKeysFinishedRegExp.Match(logLine) == true {
				pass = true
				break
			}
		}

		if pass == true {
			break
		}
	}

	ApiAddressRegExp = regexp.MustCompile("(?mi)(^.*\\(addr:\\s+[\"'])(?P<ApiAddress>[\\d.:]+)([\"'].*)")
	ClusterAddressRegExp = regexp.MustCompile("(?mi)(^.*cluster\\s+address:\\s+[\"'])(?P<ClusterAddress>[\\d.:]+)([\"'].*)")
	UnsealKeyRegExp = regexp.MustCompile("(?mi)(^.*Unseal\\s+Key:\\s+)(?P<UnsealKey>.{44})(.*)")
	RootTokenRegExp = regexp.MustCompile("(?mi)(^.*Root\\s+Token:\\s+)(?P<RootToken>.{26})(.*)")
	for _, logLine := range logData {
		if ApiAddressRegExp.Match(logLine) == true {
			tmp = ApiAddressRegExp.ReplaceAll(logLine, []byte("${ApiAddress}"))
			if len(tmp) != 0 {
				ApiAddress = string(tmp)
			}
		}

		if ClusterAddressRegExp.Match(logLine) == true {
			tmp = ClusterAddressRegExp.ReplaceAll(logLine, []byte("${ClusterAddress}"))
			if len(tmp) != 0 {
				ClusterAddress = string(tmp)
			}
		}

		if UnsealKeyRegExp.Match(logLine) == true {
			tmp = UnsealKeyRegExp.ReplaceAll(logLine, []byte("${UnsealKey}"))
			if len(tmp) != 0 {
				vaultUnsealKey = string(tmp)
				continue
			}
		}

		if RootTokenRegExp.Match(logLine) == true {
			tmp = RootTokenRegExp.ReplaceAll(logLine, []byte("${RootToken}"))
			if len(tmp) != 0 {
				vaultRootToken = string(tmp)
				break
			}
		}
	}

	return
}
