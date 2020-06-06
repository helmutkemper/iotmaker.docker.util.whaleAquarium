package iotmaker_docker_util_whaleAquarium

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
	"path/filepath"
	"sync"
)

//  mountType:
//     KVolumeMountTypeBind - TypeBind is the type for mounting host dir
//     KVolumeMountTypeVolume - TypeVolume is the type for remote storage
//     volumes
//     KVolumeMountTypeTmpfs - TypeTmpfs is the type for mounting tmpfs
//     KVolumeMountTypeNpipe - TypeNamedPipe is the type for mounting
//     Windows named pipes
//  source: relative file/dir path in computer
//  destination: full path inside container
type Mount struct {
	MountType   VolumeMountType
	Source      string
	Destination string
}

type Docker struct {
	cli       *client.Client
	ctx       context.Context
	networkId map[string]string
	imageId   map[string]string
	container map[string]container.ContainerCreateCreatedBody
}

// Must be first function call
func (el *Docker) Init() error {
	el.contextCreate()
	return el.clientCreate()
}

func (el *Docker) contextCreate() {
	el.ctx = context.Background()
}

// Negotiate best docker version
func (el *Docker) clientCreate() error {
	var err error

	el.cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	return err
}

// Mount nat por list by image config
func (el *Docker) ImageMountNatPortList(imageId string) (error, nat.PortMap) {
	var err error
	var portList []string
	var ret nat.PortMap = make(map[nat.Port][]nat.PortBinding)

	err, portList = el.ImageListExposedPorts(imageId)
	if err != nil {
		return err, nat.PortMap{}
	}

	for _, port := range portList {
		ret[nat.Port(port)] = []nat.PortBinding{
			{
				HostPort: port,
			},
		}
	}

	return err, ret
}

// Get an absolute path from file
func (el *Docker) FileMakeAbsolutePath(filePath string) (error, string) {
	fileAbsolutePath, err := filepath.Abs(filePath)
	return err, fileAbsolutePath
}

// Create a container
//   imageName: image name for download and pull
//   containerName: unique container name
//   RestartPolicy:
//      KRestartPolicyNo - Do not automatically restart the container. (the
//          default)
//      KRestartPolicyOnFailure - Restart the container if it exits due to an
//          error, which manifests as a non-zero exit code.
//      KRestartPolicyAlways - Always restart the container if it stops. If it is
//          manually stopped, it is restarted only when Docker daemon restarts or
//          the container itself is manually restarted. (See the second bullet
//          listed in restart policy details)
//      KRestartPolicyUnlessStopped - Similar to always, except that when the
//          container is stopped (manually or otherwise), it is not restarted
//          even after Docker daemon restarts.
//   mountVolumes: please use a factoryWhaleAquarium.NewVolumeMount()
//      for a complete list of volumes exposed by image, use
//      ImageListExposedVolumes(id) and ImageListExposedVolumesByName(name)
func (el *Docker) ContainerCreate(imageName, containerName string, restart RestartPolicy, mountVolumes []mount.Mount, net *network.NetworkingConfig) error {
	var err error
	var imageId string
	var portExposedList nat.PortMap
	var resp container.ContainerCreateCreatedBody

	err, imageId = el.ImageFindIdByName(imageName)
	err, portExposedList = el.ImageMountNatPortList(imageId)

	if len(el.container) == 0 {
		el.container = make(map[string]container.ContainerCreateCreatedBody)
	}

	resp, err = el.cli.ContainerCreate(
		el.ctx,
		&container.Config{
			Image: imageName,
		},
		&container.HostConfig{
			PortBindings: portExposedList,
			RestartPolicy: container.RestartPolicy{
				Name: restart.String(),
			},
			Resources: container.Resources{},
			Mounts:    mountVolumes,
		},
		net,
		nil,
		containerName,
	)
	if err != nil {
		return err
	}

	el.container[resp.ID] = resp

	return nil
}

// list image exposed ports by name
func (el *Docker) ImageListExposedPortsByName(name string) (error, []string) {
	var err error
	var id string
	err, id = el.ImageFindIdByName(name)
	if err != nil {
		return err, nil
	}

	return el.ImageListExposedPorts(id)
}

// list image exposed ports by id
func (el *Docker) ImageListExposedPorts(id string) (error, []string) {
	var err error
	var imageData types.ImageInspect
	var ret = make([]string, 0)

	imageData, _, err = el.cli.ImageInspectWithRaw(el.ctx, id)
	if err != nil {
		return err, []string{}
	}
	for port := range imageData.ContainerConfig.ExposedPorts {
		ret = append(ret, port.Port()+"/"+port.Proto())
	}

	return nil, ret
}

// list exposed volumes from image by name
func (el *Docker) ImageListExposedVolumesByName(name string) (error, []string) {
	var err error
	var id string
	err, id = el.ImageFindIdByName(name)
	if err != nil {
		return err, nil
	}

	return el.ImageListExposedVolumes(id)
}

// list exposed volumes from image by id
func (el *Docker) ImageListExposedVolumes(id string) (error, []string) {
	var err error
	var imageData types.ImageInspect
	var ret = make([]string, 0)

	imageData, _, err = el.cli.ImageInspectWithRaw(el.ctx, id)
	if err != nil {
		return err, []string{}
	}
	for volume := range imageData.ContainerConfig.Volumes {
		ret = append(ret, volume)
	}

	return nil, ret
}

// verify if exposed volume (folder only) defined by user is exposed
// in image
func (el *Docker) ImageVerifyVolume(id, path string) (error, bool) {
	err, list := el.ImageListExposedVolumes(id)
	if err != nil {
		return err, false
	}

	for _, volume := range list {
		if volume == path {
			return nil, true
		}
	}

	return nil, false
}

// find image id by name
func (el *Docker) ImageFindIdByName(name string) (error, string) {
	err, list := el.ImageList()
	if err != nil {
		return err, ""
	}

	if len(el.imageId) == 0 {
		el.imageId = make(map[string]string)
	}

	for _, data := range list {
		for _, dataTag := range data.RepoTags {
			if dataTag == name {
				el.imageId[name] = data.ID
				return nil, data.ID
			}
		}
	}

	return errors.New("image name not found"), ""
}

// list images
func (el *Docker) ImageList() (error, []types.ImageSummary) {
	ret, err := el.cli.ImageList(el.ctx, types.ImageListOptions{})
	return err, ret
}

// wait image pull be completed
func (el *Docker) ImageWaitPull(name string) error {
	var wg sync.WaitGroup

	_, found := el.imageId[name]
	if found == false {
		return errors.New("image name not found in id list")
	}

	wg.Add(1)
	go func(el *Docker, wg *sync.WaitGroup, name string) {

		for {
			err, id := el.ImageFindIdByName(name)
			if err != nil {
				panic(err)
			}

			if id != "" {
				wg.Done()
				return
			}
		}

	}(el, &wg, name)

	wg.Wait()

	return nil
}

// image pull
func (el *Docker) ImagePull(name string, attachStdOut bool) error {
	reader, err := el.cli.ImagePull(el.ctx, name, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	if len(el.imageId) == 0 {
		el.imageId = make(map[string]string)
	}

	el.imageId[name] = ""

	if attachStdOut == true {
		io.Copy(os.Stdout, reader)
	}

	return err
}

// verify if network name exists
func (el *Docker) NetworkVerifyName(name string) (error, bool) {
	resp, err := el.cli.NetworkList(el.ctx, types.NetworkListOptions{})
	if err != nil {
		return err, false
	}

	for _, v := range resp {
		if v.Name == name {
			return nil, true
		}
	}

	return nil, false
}

// remove network by name
func (el *Docker) NetworkRemove(name string) error {
	_, found := el.networkId[name]
	if found != false {
		return errors.New("network name not found in network created list")
	}

	return el.cli.NetworkRemove(el.ctx, el.networkId[name])
}

// create network
func (el *Docker) NetworkCreate(name string) error {
	resp, err := el.cli.NetworkCreate(el.ctx, name, types.NetworkCreate{
		Labels: map[string]string{
			"name": name,
		},
	})

	if err != nil {
		return err
	}

	if len(el.networkId) == 0 {
		el.networkId = make(map[string]string)
	}

	el.networkId[name] = resp.ID

	return err
}
