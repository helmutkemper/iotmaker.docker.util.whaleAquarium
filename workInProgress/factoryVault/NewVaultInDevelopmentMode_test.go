package factoryVault

import (
	"encoding/json"
	"errors"
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

	yourRealData, _ := json.Marshal(dataFromValt.Data)
	fmt.Printf("vault data: %s\n", yourRealData)

	// Output:
	// vault data: {"data":{"your_key":"please understand: put your data inside the \"data\" key or the vault will return the error \u003e\u003eno data provided\u003c\u003c"},"metadata":{"created_time":"2020-07-30T19:56:08.675428254Z","deletion_time":"","destroyed":false,"version":4}}
}
