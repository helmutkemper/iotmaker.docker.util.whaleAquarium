package factorycontainerfromremoteserver

import (
	"fmt"
	"github.com/docker/docker/api/types/mount"
	tools_garbage_collector "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"io/ioutil"
	"net/http"
)

// Before run this example, you must create a file [windows: C:]/static/index.html and put
// "It's alive!" inside file, without html tags and line breaks.
func ExampleNewContainerFromRemoteServerChangeVolumes() {
	var err error
	var pullStatusChannel = iotmakerdocker.NewImagePullStatusChannel()
	var volumesList []mount.Mount

	go func(c chan iotmakerdocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				//fmt.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					//fmt.Println("image pull complete!")
				}
			}
		}

	}(*pullStatusChannel)

	err = tools_garbage_collector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	volumesList, err = iotmakerdocker.NewVolumeMount(
		[]iotmakerdocker.Mount{
			{
				MountType:   iotmakerdocker.KVolumeMountTypeBind,
				Source:      "C:\\static", //absolute or relative to this code
				Destination: "/static",    //folder inside container
			},
		},
	)
	if err != nil {
		panic(err)
	}

	_, _, err = NewContainerFromRemoteServerChangeVolumes(
		"server_delete_before_test:latest",
		"container_delete_before_test",
		"https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git",
		[]string{},
		volumesList,
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}

	var resp *http.Response
	var site []byte
	resp, err = http.Get("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	site, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	err = tools_garbage_collector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", site)
	// Output:
	// It's alive!
}
