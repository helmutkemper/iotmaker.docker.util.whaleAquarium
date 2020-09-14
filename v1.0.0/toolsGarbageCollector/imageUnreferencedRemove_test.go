package toolsGarbageCollector

import (
	"fmt"
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func ExampleImageUnreferencedRemove() {
	var err error
	var list []types.ImageSummary

	err = ImageUnreferencedRemove()
	if err != nil {
		panic(err)
	}

	var dockerSys = iotmakerdocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(err)
	}

	list, err = dockerSys.ImageList()
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
