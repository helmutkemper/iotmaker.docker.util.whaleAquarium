package iotmaker_docker_util_whaleAquarium

// list exposed volumes from image by name
func (el *DockerSystem) ImageListExposedVolumesByName(name string) (error, []string) {
	var err error
	var id string
	err, id = el.ImageFindIdByName(name)
	if err != nil {
		return err, nil
	}

	return el.ImageListExposedVolumes(id)
}
