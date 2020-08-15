package factoryContainerMongoDB

import (
	"bytes"
	"encoding/json"
	"fmt"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/factoryContainerFromRemoteServer"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func toFile(logLine []byte) {
	var err error
	var f *os.File
	f, err = os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if bytes.HasSuffix(logLine, []byte("\n")) == false {
		logLine = append(logLine, []byte("\n")...)
	}
	if _, err := f.Write(logLine); err != nil {
		_ = f.Close() // ignore error; Write error takes precedence
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}

func ExampleNewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort() {
	var err error
	var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()
	var networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				//fmt.Printf("image pull status: %+v\n", status)

				toFile([]byte(status.Stream))
				if status.Closed == true {
					fmt.Println("image pull complete!")
				}
			}
		}

	}(*pullStatusChannel)

	err, _, networkAutoConfiguration = factoryDocker.NewNetwork(
		"network_delete_before_test",
	)
	if err != nil {
		panic(err)
	}

	// address: 10.0.0.2
	err, _, _ = NewSingleEphemeralInstanceMongoWithNetworkConfigurationAndWithoutPort(
		"container_a_delete_before_test",
		iotmakerDocker.KRestartPolicyOnFailure,
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
		iotmakerDocker.KRestartPolicyOnFailure,
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
		iotmakerDocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	// address: 10.0.0.5:8080
	err, _, _, _ = factoryContainerFromRemoteServer.NewContainerFromRemoteServerWithNetworkConfiguration(
		"image_server_delete_before_test:latest",
		"cont_server_delete_before_test",
		iotmakerDocker.KRestartPolicyUnlessStopped,
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

	// stop and remove a container
	err = toolsGarbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	// Output:
	// image pull complete!
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
