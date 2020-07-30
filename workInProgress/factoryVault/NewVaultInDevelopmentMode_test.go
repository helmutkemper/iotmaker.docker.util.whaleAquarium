package factoryVault

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"net/http"
	"time"
)

func ExampleNewVaultInDevelopmentMode() {
	var err error
	var containerID string
	var ApiAddress string
	var ClusterAddress string
	var vaultRootToken string
	var vaultUnsealKey string

	//var vaultAddr  = "0.0.0.0:8200"
	var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				//fmt.Printf("image pull status: %+v\n", status.Stream)

				if status.Closed == true {
					fmt.Println("image pull complete!")
				}
			}
		}

	}(*pullStatusChannel)

	err, containerID, ApiAddress, ClusterAddress, vaultRootToken, vaultUnsealKey = NewVaultInDevelopmentMode(
		"vaultContainer",
		KVaultVersionTag_latest,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}
	_ = ClusterAddress

	fmt.Printf("container ID: %v\n", containerID)

	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	var client *api.Client
	var data *api.Secret
	client, err = api.NewClient(&api.Config{Address: "http://" + ApiAddress, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}

	_ = vaultUnsealKey
	client.SetToken(vaultRootToken)

	_, err = client.Logical().Write("secret/data/my-secret", map[string]interface{}{"teste": "testado ok"})
	if err != nil {
		panic(err)
	}

	data, err = client.Logical().Read("secret/data/my-secret")
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(data.Data)
	fmt.Println(string(b))

	// Output:
	//
}
