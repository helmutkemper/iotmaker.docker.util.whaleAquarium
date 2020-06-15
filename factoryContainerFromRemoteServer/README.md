```golang
package main

import (
  "fmt"
  whaleAquarium "github.com/helmutkemper/iotmaker.docker"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerFromRemoteServer"
  "github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()

func init() {
  go func(c chan whaleAquarium.ContainerPullStatusSendToChannel) {

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
}

func main() {
  var err error

  err, _, _ = factoryContainerFromRemoteServer.NewContainerFromRemoteServer(
    "server:latest",
    "serverLocal",
    "server_network",
    "https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git",
    []string{},
    pullStatusChannel,
  )
  if err != nil {
    panic(err)
  }
}
```

```golang
package main

import (
  "fmt"
  whaleAquarium "github.com/helmutkemper/iotmaker.docker"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerFromRemoteServer"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
  "github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()

func init() {
  go func(c chan whaleAquarium.ContainerPullStatusSendToChannel) {

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
}

func main() {

  // define an external MongoDB config file path
  err, volumesList := factoryWhaleAquarium.NewVolumeMount(
    []whaleAquarium.Mount{
      {
        MountType:   whaleAquarium.KVolumeMountTypeBind,
        Source:      "C:\\static", //relative to this code
        Destination: "/static",     //container folder
      },
    },
  )
  if err != nil {
    panic(err)
  }

  err, _, _ = factoryContainerFromRemoteServer.NewContainerFromRemoteServerChangeVolumes(
    "server:latest",
    "serverLocal",
    "server_network",
    "https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git",
    []string{},
    volumesList,
    pullStatusChannel,
  )
  if err != nil {
    panic(err)
  }
}
```

```golang
package main

import (
  "github.com/docker/go-connections/nat"
  whaleAquarium "github.com/helmutkemper/iotmaker.docker"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryContainerFromRemoteServer"
  "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/factoryWhaleAquarium"
)

func main() {

  currentPort, err := nat.NewPort("tcp", "3000")
  if err != nil {
    panic(err)
  }

  newPort, err := nat.NewPort("tcp", "8080")
  if err != nil {
    panic(err)
  }

  currentPortList  := []nat.Port{
    currentPort,
  }

  newPortList  := []nat.Port{
    newPort,
  }

  // define an external MongoDB config file path
  err, volumesList := factoryWhaleAquarium.NewVolumeMount(
    []whaleAquarium.Mount{
      {
        MountType:   whaleAquarium.KVolumeMountTypeBind,
        Source:      "./server/static", //relative to this code
        Destination: "/app/static",     //container folder
      },
    },
  )
  if err != nil {
    panic(err)
  }

  err, _, _ = factoryContainerFromRemoteServer.NewContainerFromRemoteServerChangeExposedPortAndVolumes(
    "server:latest",
    "serverLoval",
    "server_network",
    "https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git",
    []string{
      "server:latest",
    },
    currentPortList,
    newPortList,
    volumesList,
    nil,
  )
  if err != nil {
    panic(err)
  }
}
```