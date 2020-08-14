package factoryContainerMongoDB

import (
	"context"
	"fmt"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ExampleNewSingleEphemeralInstanceMongoLatest() {
	var err error
	var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()
	var containerID string

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

	err, containerID = NewSingleEphemeralInstanceMongoLatest(
		"container_delete_before_test",
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

	// stop and remove a container
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(err)
	}

	err = dockerSys.ContainerStopAndRemove(containerID, true, false, false)
	if err != nil {
		panic(err)
	}

	err = toolsGarbageCollector.ImageUnreferencedRemove()
	if err != nil {
		panic(err)
	}

	err = toolsGarbageCollector.VolumesUnreferencedRemove()
	if err != nil {
		panic(err)
	}

	// Output:
	// image pull complete!
	// ping ok
}
