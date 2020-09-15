package factorycontainermongodb

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	tools_garbage_collector "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ExampleNewSingleEphemeralInstanceMongoWithPort() {
	var err error
	var newPort nat.Port
	var pullStatusChannel = iotmakerdocker.NewImagePullStatusChannel()

	go func(c chan iotmakerdocker.ContainerPullStatusSendToChannel) {

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

	// stop and remove containers and garbage collector
	err = tools_garbage_collector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	newPort, err = nat.NewPort("tcp", "27015")
	if err != nil {
		panic(err)
	}

	err, _ = NewSingleEphemeralInstanceMongoWithPort(
		"container_a_delete_before_test",
		newPort,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	newPort, err = nat.NewPort("tcp", "27016")
	if err != nil {
		panic(err)
	}

	err, _ = NewSingleEphemeralInstanceMongoWithPort(
		"container_b_delete_before_test",
		newPort,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	newPort, err = nat.NewPort("tcp", "27017")
	if err != nil {
		panic(err)
	}

	err, _ = NewSingleEphemeralInstanceMongoWithPort(
		"container_c_delete_before_test",
		newPort,
		KMongoDBVersionTag_3,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 2)
	var mongoClient *mongo.Client
	var ctx context.Context
	mongoClient, err = mongo.NewClient(
		options.Client().ApplyURI(
			"mongodb://127.0.0.1:27015," +
				"127.0.0.1:27016," +
				"127.0.0.1:27017",
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

	err = tools_garbage_collector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	// Output:
	// image pull complete!
	// image pull complete!
	// image pull complete!
	// ping ok
}
