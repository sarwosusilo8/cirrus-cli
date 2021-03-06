package instance

import (
	"context"
	"github.com/cirruslabs/cirrus-ci-agent/api"
	"github.com/cirruslabs/cirrus-cli/internal/executor/instance/runconfig"
	"github.com/cirruslabs/cirrus-cli/internal/executor/platform"
	"path"
)

type ContainerInstance struct {
	Image                string
	CPU                  float32
	Memory               uint32
	AdditionalContainers []*api.AdditionalContainer
	Platform             platform.Platform
}

func (inst *ContainerInstance) Run(ctx context.Context, config *runconfig.RunConfig) (err error) {
	workingVolume, err := CreateWorkingVolumeFromConfig(ctx, config, inst.Platform)
	if err != nil {
		return err
	}
	defer func() {
		if config.ContainerOptions.NoCleanup {
			config.Logger.Infof("not cleaning up working volume %s, don't forget to remove it with \"docker volume rm %s\"",
				workingVolume.Name(), workingVolume.Name())

			return
		}

		cleanupErr := workingVolume.Close(config.ContainerBackend)
		if err == nil {
			err = cleanupErr
		}
	}()

	params := &Params{
		Image:                inst.Image,
		CPU:                  inst.CPU,
		Memory:               inst.Memory,
		AdditionalContainers: inst.AdditionalContainers,
		Platform:             inst.Platform,
		WorkingVolumeName:    workingVolume.Name(),
	}

	return RunContainerizedAgent(ctx, config, params)
}

func (inst *ContainerInstance) WorkingDirectory(projectDir string, dirtyMode bool) string {
	return path.Join(inst.Platform.WorkingVolumeMountpoint(), platform.WorkingVolumeWorkingDir)
}
