# iotmaker.docker.util.whaleAquarium

this package makes it easy to install tools needed for my personal projects
> Work in progress
>
> The functions listed below are ready for use:
> NewSingleEphemeralInstanceMongoWithPort()
> NewSingleEphemeralInstanceMongoLatestWithPort()
> NewSingleEphemeralInstanceMongoLatest()
> NewSingleEphemeralInstanceMongo()

### MongoDB

This example install three containers MongoDB with ephemeral storage data
```golang
package main

import (
  "github.com/docker/go-connections/nat"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
)

func main() {
  var err error
  var id string

  portA, _ := nat.NewPort("tcp", "27016")
  portB, _ := nat.NewPort("tcp", "27017")
  portC, _ := nat.NewPort("tcp", "27018")

  ch := factoryWhaleAquarium.NewPullStatusMonitor()

  err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoWithPort(
    "MongoDBTete",
    "mongodb_network",
    portA,
    factoryContainerMongoDB.KMongoDBVersionTag_latest,
    ch,
  )
  if err != nil {
    panic(err)
  }

  err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoWithPort(
    "MongoDBTete2",
    "mongodb_network",
    portB,
    factoryContainerMongoDB.KMongoDBVersionTag_latest,
    ch,
  )
  if err != nil {
    panic(err)
  }

  err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoWithPort(
    "MongoDBTete3",
    "mongodb_network",
    portC,
    factoryContainerMongoDB.KMongoDBVersionTag_latest,
    ch,
  )
  if err != nil {
    panic(err)
  }

  _ = id
}
```
This example install one container MongoDB with ephemeral storage data
```golang
package main

import (
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
)

func main() {

  ch := factoryWhaleAquarium.NewPullStatusMonitor()
  err, _ := factoryContainerMongoDB.NewSingleEphemeralInstanceMongoLatest(
    "mongoLocal",
    "dockerNetwork",
    ch,
  )
  if err != nil {
    panic(err)
  }
}
```

How to monitoring image pull download
```golang
package main

import (
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
)

func main() {
  var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()

  go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {
    for {
      select {
      case status := <-c:
        fmt.Printf("image pull status: %+v\n", status)

        if status.Closed == true {
          fmt.Println("image pull complete!")
        }
      }
    }
  }(pullStatusChannel)

  err, _ := factoryContainerMongoDB.NewSingleEphemeralInstanceMongoLatest(
    "mongoLocal",
    "dockerNetwork",
    pullStatusChannel,
  )
  if err != nil {
    panic(err)
  }
}
```
