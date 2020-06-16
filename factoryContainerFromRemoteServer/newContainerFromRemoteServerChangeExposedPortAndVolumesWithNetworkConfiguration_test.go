package factoryContainerFromRemoteServer

import (
	"fmt"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"io/ioutil"
	"net/http"
)

func ExampleNewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration() {
	var err error
	var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()
	var volumesList []mount.Mount

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				//fmt.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					fmt.Println("image pull complete!")
				}
			}
		}

	}(*pullStatusChannel)

	currentPort, err := nat.NewPort("tcp", "3000")
	if err != nil {
		panic(err)
	}

	newPort, err := nat.NewPort("tcp", "8080")
	if err != nil {
		panic(err)
	}

	currentPortList := []nat.Port{
		currentPort,
	}

	newPortList := []nat.Port{
		newPort,
	}

	err, volumesList = factoryWhaleAquarium.NewVolumeMount(
		[]iotmakerDocker.Mount{
			{
				MountType:   iotmakerDocker.KVolumeMountTypeBind,
				Source:      "C:\\static", //absolute or relative to this code
				Destination: "/static",    //folder inside container
			},
		},
	)
	if err != nil {
		panic(err)
	}

	err, _, _, _ = NewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration(
		"server_delete_before_test:latest",
		"container_delete_before_test",
		iotmakerDocker.KRestartPolicyOnFailure,
		"network_delete_before_test",
		iotmakerDocker.KNetworkDriveBridge,
		"local",
		"10.0.0.0/16",
		"10.0.0.1",
		"https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git",
		[]string{},
		currentPortList,
		newPortList,
		volumesList,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	var resp *http.Response
	var site []byte
	resp, err = http.Get("http://localhost:8080")
	site, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", site)
	// Output:
	// image pull complete!
	// It's alive!
}
