package factorycontainermongodb

import (
	"encoding/json"
	"fmt"
	factory_container_from_remote_server "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/factoryContainerFromRemoteServer"
	tools_garbage_collector "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"io/ioutil"
	"log"
	"net/http"
)

func ExampleNewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort() {
	var err error
	var pullStatusChannel = iotmakerdocker.NewImagePullStatusChannel()
	var networkAutoConfiguration *iotmakerdocker.NextNetworkAutoConfiguration

	go func(c chan iotmakerdocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				//fmt.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					//fmt.Println("image pull complete!")
				}
			}
		}

	}(*pullStatusChannel)

	// stop and remove containers and garbage collector
	err = tools_garbage_collector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	_, networkAutoConfiguration, err = iotmakerdocker.NewNetwork(
		"network_delete_before_test",
	)
	if err != nil {
		panic(err)
	}

	// address: 10.0.0.2
	err, _, _ = NewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort(
		"container_a_delete_before_test",
		iotmakerdocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	// address: 10.0.0.3
	err, _, _ = NewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort(
		"container_b_delete_before_test",
		iotmakerdocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	// address: 10.0.0.4
	err, _, _ = NewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort(
		"container_c_delete_before_test",
		iotmakerdocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	// address: 10.0.0.5:8080
	err, _, _, _ = factory_container_from_remote_server.NewContainerFromRemoteServerWithNetworkConfiguration(
		"image_server_delete_before_test:latest",
		"cont_server_delete_before_test",
		iotmakerdocker.KRestartPolicyUnlessStopped,
		networkAutoConfiguration,
		"https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.mongodb.test.git",
		[]string{},
		pullStatusChannel,
	)
	if err != nil && err.Error() != "there is a network with this name" {
		panic(err)
	}

	verifyServer("http://127.0.0.1:8080/?db=10.0.0.2:27017")
	verifyServer("http://127.0.0.1:8080/?db=10.0.0.3:27017")
	verifyServer("http://127.0.0.1:8080/?db=10.0.0.4:27017")

	// stop and remove containers and garbage collector
	err = tools_garbage_collector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	// Output:
	// map[error: ok:true]
	// map[error: ok:true]
	// map[error: ok:true]
}

func verifyServer(url string) {
	var err error
	var serverReader *http.Response
	var serverResponse []byte
	var dataToOutput interface{}

	serverReader, err = http.Get(url)
	if err != nil {
		log.Printf("http server get error: %v", err.Error())
		return
	}

	serverResponse, err = ioutil.ReadAll(serverReader.Body)
	if err != nil {
		log.Printf("http server readAll error: %v", err.Error())
		return
	}

	err = json.Unmarshal(serverResponse, &dataToOutput)
	if err != nil {
		log.Printf("http server json unmarshal error: %v", err.Error())
		return
	}

	fmt.Printf("%v\n", dataToOutput)
}
