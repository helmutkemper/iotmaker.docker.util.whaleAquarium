# iotmaker.docker.util.whaleAquarium


this package makes it easy to install tools needed for my personal projects


###MongoDB

How to install MongoDB with ephemeral storage data
```golang
package main

import "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"

func main() {
	var err error
	var id string

	err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoLatest(
		"./config.conf",
		"MongoDBTete",
		"mongodb_network",
	)
	if err != nil {
		panic(err)
	}

	_ = id
}

```