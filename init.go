package iotmaker_docker_util_whaleAquarium

// Must be first function call
func (el *DockerSystem) Init() error {
	el.contextCreate()
	return el.clientCreate()
}
