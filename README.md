# iotmaker.docker.util.whaleAquarium


this package makes it easy to install tools needed for my personal projects


### MongoDB

How to install MongoDB with ephemeral storage data
```golang
package main

import "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"

func main() {
	var err error
	var id string

	err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoLatest(
		"MongoDBTete",
		"mongodb_network",
	)
	if err != nil {
		panic(err)
	}

	_ = id
}
```

```golang
package main

import "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"

func main() {
	var err error
	var id string

	err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongo(
		"MongoDBTete",
		"mongodb_network",
		factoryContainerMongoDB.KMongoDBVersionTag_3_3_15,
	)
	if err != nil {
		panic(err)
	}

	_ = id
}
```

```golang
package main

import "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"

func main() {
	var err error
	var id string

	err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoLatestWithPort(
		"MongoDBTete",
		"mongodb_network",
        27017
	)
	if err != nil {
		panic(err)
	}

	_ = id
}
```

```golang
package main

import "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerMongoDB"

func main() {
	var err error
	var id string

	err, id = factoryContainerMongoDB.NewSingleEphemeralInstanceMongoWithPort(
		"MongoDBTete",
		"mongodb_network",
        27017
		factoryContainerMongoDB.KMongoDBVersionTag_3_3_15,
	)
	if err != nil {
		panic(err)
	}

	_ = id
}
```
