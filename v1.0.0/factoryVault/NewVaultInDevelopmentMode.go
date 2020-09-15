package factory_vault

import (
	"bytes"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"regexp"
)

// NewVaultInDevelopmentMode (English): Manage Secrets and Protect Sensitive Data
// Site: https://www.vaultproject.io/
//
// NewVaultInDevelopmentMode (Português): Gerencia segredos e protege dados sensíveis
// Site: https://www.vaultproject.io/
func NewVaultInDevelopmentMode(
	containerName string,
	version VaultVersionTag,
	//newPort nat.Port,
	pullStatus *chan iotmakerdocker.ContainerPullStatusSendToChannel,
) (
	err error,
	containerId string,
	ApiAddress string,
	ClusterAddress string,
	vaultRootToken string,
	vaultUnsealKey string,
) {

	var upgradingKeysFinishedRegExp *regexp.Regexp

	var imageName = "vault:" + version.String()
	var mountList []mount.Mount
	var log []byte
	var logData [][]byte

	// init docker
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	_, _, err = dockerSys.ImagePull(imageName, pullStatus)
	if err != nil {
		return
	}

	portMap := nat.PortMap{
		// container port number/protocol [tpc/udp]
		"8200/tcp": []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: "8200",
			},
		},
		// container port number/protocol [tpc/udp]
		"8201/tcp": []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: "8201",
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
	if err != nil {
		return
	}

	containerId, err = dockerSys.ContainerFindIdByName("vaultContainer")

	upgradingKeysFinishedRegExp = regexp.MustCompile("(?mi)upgrading keys finished") //todo: make const
	var pass = false
	for {
		log, err = dockerSys.ContainerLogs(containerId)
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

	ApiAddress, ClusterAddress, vaultUnsealKey, vaultRootToken = filterLogData(logData)
	return
}

func filterLogData(
	logData [][]byte,
) (
	ApiAddress string,
	ClusterAddress string,
	vaultUnsealKey string,
	vaultRootToken string,
) {

	var ApiAddressRegExp *regexp.Regexp
	var ClusterAddressRegExp *regexp.Regexp
	var UnsealKeyRegExp *regexp.Regexp
	var RootTokenRegExp *regexp.Regexp

	var tmp []byte

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
