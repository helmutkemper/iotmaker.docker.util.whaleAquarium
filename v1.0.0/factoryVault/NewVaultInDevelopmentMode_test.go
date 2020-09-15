package factory_vault

import (
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	tools_garbage_collector "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"net/http"
	"time"
)

// this example install an vault container and test
// https://www.vaultproject.io/
func ExampleNewVaultInDevelopmentMode() {

	var err error
	var containerID string
	var ApiAddress string
	var ClusterAddress string
	var vaultRootToken string
	var vaultUnsealKey string

	var pullStatusChannel = iotmakerdocker.NewImagePullStatusChannel()

	go func(c chan iotmakerdocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				//fmt.Printf("image pull status: %+v\n", status.Stream)

				if status.Closed == true {
					//fmt.Println("image pull complete!")
				}
			}
		}

	}(*pullStatusChannel)

	err = tools_garbage_collector.RemoveAllByNameContains("vaultContainer")
	if err != nil {
		panic(err)
	}

	err, containerID, ApiAddress, ClusterAddress, vaultRootToken, vaultUnsealKey = NewVaultInDevelopmentMode(
		"vaultContainer",
		KVaultVersionTag_latest,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}
	_ = ClusterAddress
	_ = vaultUnsealKey

	if containerID == "" {
		err = errors.New("container id not found")
		panic(err)
	}

	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	var client *api.Client
	var dataFromValt *api.Secret
	client, err = api.NewClient(&api.Config{Address: "http://" + ApiAddress, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}

	client.SetToken(vaultRootToken)

	var dataToVault = make(map[string]interface{})
	var yourData = make(map[string]interface{})
	yourData["your_key"] = "please understand: put your data inside the \"data\" key or the vault will return the error >>no data provided<<"
	dataToVault["data"] = yourData

	_, err = client.Logical().Write("secret/data/my-secret", dataToVault)
	if err != nil {
		panic(err)
	}

	dataFromValt, err = client.Logical().Read("secret/data/my-secret")
	if err != nil {
		panic(err)
	}

	if dataFromValt != nil {
		d := dataFromValt.Data
		fmt.Printf("vault data: %v\n", d["data"])
	}

	// stop and remove a container
	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(err)
	}

	err = dockerSys.ContainerStopAndRemove(containerID, true, false, false)
	if err != nil {
		panic(err)
	}

	err = tools_garbage_collector.RemoveAllByNameContains("vaultContainer")
	if err != nil {
		panic(err)
	}

	// Output:
	// vault data: map[your_key:please understand: put your data inside the "data" key or the vault will return the error >>no data provided<<]
}
