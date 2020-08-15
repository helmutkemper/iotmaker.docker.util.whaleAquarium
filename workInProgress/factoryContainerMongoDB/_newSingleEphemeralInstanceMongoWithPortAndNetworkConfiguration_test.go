package factoryContainerMongoDB

import (
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/factoryContainerFromRemoteServer"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

func ExampleNewSingleEphemeralInstanceMongoWithPortAndNetworkConfiguration() {
	var err error
	var newPort nat.Port
	var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()
	var networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {

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

	err, _, networkAutoConfiguration = factoryDocker.NewNetwork(
		"network_delete_before_test",
	)
	if err != nil {
		panic(err)
	}

	newPort, err = nat.NewPort("tcp", "27016")
	if err != nil {
		panic(err)
	}

	err, _, _ = NewSingleEphemeralInstanceMongoWithPortAndNetworkConfiguration(
		"container_a_delete_before_test",
		iotmakerDocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		newPort,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil && err.Error() != "there is a network with this name" {
		panic(err)
	}

	newPort, err = nat.NewPort("tcp", "27017")
	if err != nil {
		panic(err)
	}

	err, _, _ = NewSingleEphemeralInstanceMongoWithPortAndNetworkConfiguration(
		"container_b_delete_before_test",
		iotmakerDocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		newPort,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil && err.Error() != "there is a network with this name" {
		panic(err)
	}

	newPort, err = nat.NewPort("tcp", "27018")
	if err != nil {
		panic(err)
	}

	err, _, _ = NewSingleEphemeralInstanceMongoWithPortAndNetworkConfiguration(
		"container_c_delete_before_test",
		iotmakerDocker.KRestartPolicyOnFailure,
		networkAutoConfiguration,
		newPort,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil && err.Error() != "there is a network with this name" {
		panic(err)
	}

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

	verifyServer("http://127.0.0.1:8080/?db=10.0.0.2:27016")
	verifyServer("http://127.0.0.1:8080/?db=10.0.0.3:27017")
	verifyServer("http://127.0.0.1:8080/?db=10.0.0.4:27018")

	err = toolsGarbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	// Output:
	// map[error: ok:true]
	// map[error: ok:true]
	// map[error: ok:true]
}
