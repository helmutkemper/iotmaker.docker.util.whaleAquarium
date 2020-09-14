# iotmaker.docker.util.whaleAquarium

<p align="center">
  <img src="https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/blob/master/image/Go-Logo_LightBlue.svg" width="500">
</p>
<p align="center">
  <a href="https://goreportcard.com/report/github.com/helmutkemper/iotmaker.docker.util.whaleAquarium">
    <img src="https://goreportcard.com/badge/github.com/helmutkemper/iotmaker.docker.util.whaleAquarium">
  </a>
  <a href="https://pkg.go.dev/github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0?tab=doc">
    <img src="https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/blob/master/image/godoc.svg">
  </a>
</p>

> work in progress! sorry...

This package makes it easy to install ready-to-use containers.
> For usage examples, see the test functions within the code

#### MongoDB:

|Function                                                                               |
|---------------------------------------------------------------------------------------|
| NewSingleEphemeralInstanceMongo()                                                     |
| NewSingleEphemeralInstanceMongoLatest()                                               |
| NewSingleEphemeralInstanceMongoLatestWithPort()                                       |
| NewSingleEphemeralInstanceMongoWithPort()                                             |
| NewSingleEphemeralInstanceMongoWithPortAndNetworkConfiguration()                      |

#### Project from a remote server, like github.com:

|Function                                                                               |
|---------------------------------------------------------------------------------------|
| NewContainerFromRemoteServer()                                                        |
| NewContainerFromRemoteServerChangeExposedPortAndVolumes()                             |
| NewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration()     |
| NewContainerFromRemoteServerChangeVolumes()                                           |
| NewContainerFromRemoteServerChangeVolumesWithNetworkConfiguration()                   |
| NewContainerFromRemoteServerWithNetworkConfiguration()                                |

**Sample**:

[iotmaker.docker.util.whaleAquarium.sample](https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample)

##### Server example:

```
server := "https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git"
server := "https://x-access-token:__TOKEN__@helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git"
```

##### To get a github token
**settings** > **Developer settings** > **Personal access tokens** > **Generate new token**

Mark [x]repo - Full control of private repositories
