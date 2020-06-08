# iotmaker.docker.util.whaleAquarium


this package makes it easy to install tools needed for my personal projects
> please, don't use. work in progress.

### MongoDB

This example install three containers MongoDB with ephemeral storage data
```golang
package main

import (
  "github.com/docker/go-connections/nat"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"
)

func main() {
  var err error
  var id string

  portA, _ := nat.NewPort("tcp", "27016")
  portB, _ := nat.NewPort("tcp", "27017")
  portC, _ := nat.NewPort("tcp", "27018")

  //todo: id network, id image
  err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoWithPort(
    "MongoDBTete",
    "mongodb_network",
    portA,
    factoryContainerMongoDB.KMongoDBVersionTag_latest,
  )
  if err != nil {
    panic(err)
  }

  err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoWithPort(
    "MongoDBTete2",
    "mongodb_network",
    portB,
    factoryContainerMongoDB.KMongoDBVersionTag_latest,
  )
  if err != nil {
    panic(err)
  }

  err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoWithPort(
    "MongoDBTete3",
    "mongodb_network",
    portC,
    factoryContainerMongoDB.KMongoDBVersionTag_latest,
  )
  if err != nil {
    panic(err)
  }

  _ = id
}
```
