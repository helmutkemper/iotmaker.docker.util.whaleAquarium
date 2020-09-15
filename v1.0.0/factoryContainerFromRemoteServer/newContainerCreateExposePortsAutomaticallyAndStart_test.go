package factorycontainerfromremoteserver

import (
	"fmt"
	toolsgarbagecollector "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/toolsGarbageCollector"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"io/ioutil"
	"net/http"
)

// Before run this example, you must create a file [windows: C:]/static/index.html and put
// "It's alive!" inside file, without html tags and line breaks.
func ExampleNewContainerCreateExposePortsAutomaticallyAndStart() {
	var err error
	var pullStatusChannel = iotmakerdocker.NewImagePullStatusChannel()

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

	err = toolsgarbagecollector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	_, _, err = NewContainerCreateExposePortsAutomaticallyAndStart(
		"server_delete_before_test:latest",
		"container_delete_before_test",
		"https://github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git",
		[]string{},
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

	err = toolsgarbagecollector.RemoveAllByNameContains("delete")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", site)
	// Output:
	// It's alive!
}
