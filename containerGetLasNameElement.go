package iotmaker_docker_util_whaleAquarium

import "strings"

func ContainerGetLasNameElement(name string) string {
	names := strings.Split(name, "/")

	l := len(names) - 1

	if l > -1 {
		return names[l]
	}

	return name
}
