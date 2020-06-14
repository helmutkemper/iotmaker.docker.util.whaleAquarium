package factoryWhaleAquarium

import (
	"errors"
	"github.com/docker/docker/api/types/mount"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/util"
)

// en: Return a mount point after verify and correct source file/dir relative
//     path.
//     This function can't verify if file/dir exists inside on the image
func NewVolumeMount(list []whaleAquarium.Mount) (err error, mountList []mount.Mount) {
	var found bool
	var fileAbsolutePath string

	for _, v := range list {
		found = util.VerifyFileExists(v.Source)
		if found == false {
			return errors.New("source file not found"), nil
		}

		err, fileAbsolutePath = util.FileGetAbsolutePath(v.Source)
		if err != nil {
			return
		}

		mountList = append(
			mountList,
			mount.Mount{
				Type:   mount.Type(v.MountType.String()),
				Source: fileAbsolutePath,
				Target: v.Destination,
			},
		)
	}

	return
}
