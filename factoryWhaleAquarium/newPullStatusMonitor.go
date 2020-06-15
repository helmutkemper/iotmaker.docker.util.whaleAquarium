package factoryWhaleAquarium

import (
	"fmt"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

func NewPullStatusMonitor() (pullStatusChannel *chan iotmakerDocker.ContainerPullStatusSendToChannel) {
	pullStatusChannel = factoryDocker.NewImagePullStatusChannel()

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				fmt.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					fmt.Println("image pull complete!")
				}
			}
		}

	}(*pullStatusChannel)

	return
}
