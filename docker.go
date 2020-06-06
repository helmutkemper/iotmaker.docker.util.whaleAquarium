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

//     mountType:
//          KVolumeMountTypeBind - TypeBind is the type for mounting host dir
//	        KVolumeMountTypeVolume - TypeVolume is the type for remote storage
//	        volumes
//	        KVolumeMountTypeTmpfs - TypeTmpfs is the type for mounting tmpfs
//	        KVolumeMountTypeNpipe - TypeNamedPipe is the type for mounting
//	        Windows named pipes
//     source: relative file/dir path in computer
//     destination: full path inside container
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
}

func (el *Docker) Prepare() error {
	el.contextCreate()
	return el.clientCreate()
}

func (el *Docker) contextCreate() {
	el.ctx = context.Background()
}

func (el *Docker) clientCreate() error {
	var err error

	el.cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	return err
}

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

func (el *Docker) FileMakeAbsolutePath(filePath string) (error, string) {
	fileAbsolutePath, err := filepath.Abs(filePath)
	return err, fileAbsolutePath
}

func (el *Docker) ContainerCreate(imageId string, restart RestartPolicy, mountVolumes []mount.Mount, net *network.NetworkingConfig) error {
	var err error
	var portExposedList nat.PortMap

	err, portExposedList = el.ImageMountNatPortList(imageId)

	resp, err = el.cli.ContainerCreate(
		el.ctx,
		&container.Config{
			Image: "mongo:latest",
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
		"MongoDBGolang",
	)
	if err != nil {
		panic(err)
	}
}

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

func (el *Docker) ImageList() (error, []types.ImageSummary) {
	ret, err := el.cli.ImageList(el.ctx, types.ImageListOptions{})
	return err, ret
}

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

func (el *Docker) NetworkRemove(name string) error {
	_, found := el.networkId[name]
	if found != false {
		return errors.New("network name not found in network created list")
	}

	return el.cli.NetworkRemove(el.ctx, el.networkId[name])
}

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
