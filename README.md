# iotmaker.docker.util.whaleAquarium

This package makes it easy to install ready-to-use containers.
> For usage examples, see the test functions within the code

#### MongoDB:

NewSingleEphemeralInstanceMongo()
NewSingleEphemeralInstanceMongoLatest()
NewSingleEphemeralInstanceMongoLatestWithPort()
NewSingleEphemeralInstanceMongoWithPort()
NewSingleEphemeralInstanceMongoWithPortAndNetworkConfiguration()

#### Project from a remote server, like github.com:

NewContainerFromRemoteServer()
NewContainerFromRemoteServerChangeExposedPortAndVolumes()
NewContainerFromRemoteServerChangeExposedPortAndVolumesWithNetworkConfiguration()
NewContainerFromRemoteServerChangeVolumes()
NewContainerFromRemoteServerChangeVolumesWithNetworkConfiguration()
NewContainerFromRemoteServerWithNetworkConfiguration()

**Sample**:

[iotmaker.docker.util.whaleAquarium.sample](https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample)

##### Server example:

```
server := "https://github.com/__USER__/__PROJECT__.git"
server := "https://x-access-token:__TOKEN__@github.com/__USER__/__PROJECT__.git"
```

##### To get a github token
**settings** > **Developer settings** > **Personal access tokens** > **Generate new token**

Mark [x]repo - Full control of private repositories
