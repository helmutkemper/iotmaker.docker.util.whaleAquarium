package factoryContainerMongoDB

import (
	"context"
	"fmt"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ExampleNewSingleEphemeralInstanceMongo() {
	var err error
	var pullStatusChannel = iotmakerdocker.NewImagePullStatusChannel()

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
	err = toolsGarbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	err, _ = NewSingleEphemeralInstanceMongo(
		"container_delete_before_test",
		KMongoDBVersionTag_latest,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	var mongoClient *mongo.Client
	var ctx context.Context
	mongoClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
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

	err = toolsGarbageCollector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	// Output:
	// ping ok
}
