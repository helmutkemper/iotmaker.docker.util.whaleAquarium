package factoryContainerFromRemoteServer

import (
	"fmt"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"io/ioutil"
	"net/http"
)

func ExampleNewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration() {
	var err error
	var pullStatusChannel = iotmakerdocker.NewImagePullStatusChannel()
	var volumesList []mount.Mount
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

	err = toolsGarbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	serverPort := nat.PortMap{
		// container port number/protocol [tpc/udp]
		"3000/tcp": []nat.PortBinding{ // server original port
			{
				// server output port number
				HostPort: "8080",
			},
		},
	}

	volumesList, err = iotmakerdocker.NewVolumeMount(
		[]iotmakerdocker.Mount{
			{
				MountType:   iotmakerdocker.KVolumeMountTypeBind,
				Source:      "C:\\static", //absolute or relative to this code
				Destination: "/static",    //folder inside container
			},
		},
	)
	if err != nil {
		panic(err)
	}

	_, networkAutoConfiguration, err = iotmakerdocker.NewNetwork("network_delete_before_test")
	if err != nil {
		panic(err)
	}

	err, _, _, _ = NewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration(
		"server_delete_before_test:latest",
		"container_delete_before_test",
		iotmakerdocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		"https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git",
		[]string{},
		serverPort,
		volumesList,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	var resp *http.Response
	var site []byte
	resp, err = http.Get("http://localhost:8080")
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

	err = toolsGarbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", site)
	// Output:
	// It's alive!
}
