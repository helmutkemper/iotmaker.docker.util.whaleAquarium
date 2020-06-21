package factoryContainerMongoDB

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
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

	newPort, err = nat.NewPort("tcp", "27017")
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

	newPort, err = nat.NewPort("tcp", "27017")
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

	var mongoClient *mongo.Client
	var ctx context.Context
	mongoClient, err = mongo.NewClient(
		options.Client().ApplyURI(
			"mongodb://" +
				"10.0.0.1:27017," +
				"10.0.0.2:27017," +
				"10.0.0.3:27017",
		),
	)
	if err != nil {
		panic(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 60*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		panic(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 60*time.Second)
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("ping ok")

	// Output:
	// image pull complete!
	// image pull complete!
	// image pull complete!
	// ping ok
}
