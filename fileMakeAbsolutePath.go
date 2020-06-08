package iotmaker_docker_util_whaleAquarium

import "path/filepath"

// Get an absolute path from file
func (el *DockerSystem) FileMakeAbsolutePath(filePath string) (error, string) {
	fileAbsolutePath, err := filepath.Abs(filePath)
	return err, fileAbsolutePath
}
