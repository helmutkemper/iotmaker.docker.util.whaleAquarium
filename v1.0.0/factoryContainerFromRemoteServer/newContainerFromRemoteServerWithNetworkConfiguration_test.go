package factory_container_from_remote_server

import (
	"fmt"
	tools_garbage_collector "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"io/ioutil"
	"net/http"
)

func ExampleNewContainerFromRemoteServerWithNetworkConfiguration() {
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

	err = tools_garbage_collector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	_, networkAutoConfiguration, err = iotmakerdocker.NewNetwork("network_delete_before_test")
	if err != nil {
		panic(err)
	}

	err, _, _, _ = NewContainerFromRemoteServerWithNetworkConfiguration(
		"server_delete_before_test:latest",
		"container_delete_before_test",
		iotmakerdocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		"https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git",
		[]string{},
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	var resp *http.Response
	var site []byte
	resp, err = http.Get("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	site, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	err = tools_garbage_collector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", site)
	// Output:
	// It's alive!
}
