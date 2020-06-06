package factoryWhaleAquarium

import (
	"errors"
	"github.com/docker/docker/api/types/mount"
	whaleAquarium "github.com/helmutkemper/iotmaker.docker.util.whaleAquarium"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/util"
)

// en: Return a mount point after verify and correct source file/dir relative
//     path.
//     This function can't verify if file/dir exists inside on the image
func NewVolumeMount(list []whaleAquarium.Mount) (error, []mount.Mount) {
	var err error
	var found bool
	var fileAbsolutePath string
	var ret = make([]mount.Mount, 0)

	for _, v := range list {
		found = util.VerifyFileExists(v.Source)
		if found == false {
			return errors.New("source file not found"), nil
		}

		err, fileAbsolutePath = util.FileGetAbsolutePath(v.Source)
		if err != nil {
			return err, nil
		}

		ret = append(
			ret,
			mount.Mount{
				Type:   v.MountType.String(),
				Source: fileAbsolutePath,
				Target: v.Destination,
			},
		)
	}

	return nil, ret
}
