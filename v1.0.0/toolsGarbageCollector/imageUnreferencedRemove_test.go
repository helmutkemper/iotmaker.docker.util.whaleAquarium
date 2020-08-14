package toolsGarbageCollector

import (
	"fmt"
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func ExampleImageUnreferencedRemove() {
	var err error
	var list []types.ImageSummary

	err = ImageUnreferencedRemove()
	if err != nil {
		panic(err)
	}

	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(err)
	}

	err, list = dockerSys.ImageList()
	for _, image := range list {
		if len(image.RepoTags) > 0 {
			if image.RepoTags[0] == "<none>:<none>" {
				fmt.Println("test fail")
				return
			}
		}
	}

	// Output:
	//
}
