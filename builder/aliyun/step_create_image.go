package aliyun

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/mitchellh/packer/packer"
)

type stepCreateImage struct {
}

func (s *stepCreateImage) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*ecs.Client)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("config").(Config)
	snapshotId := state.Get("snapshot_id").(string)

	imageId, err := client.CreateImage(&ecs.CreateImageArgs{
		RegionId: c.RegionId,
		SnapshotId: snapshotId,
		ImageName: c.ImageName,
		Description: c.ImageDescription,
	})


	if err != nil {
		err := fmt.Errorf("Error creating image: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if imageId == "" {
		err := fmt.Errorf("Error creating image, imageId is empty")
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Message("Image has been created!")

	// Store the image id for later
	state.Put("image_id", imageId)

	return multistep.ActionContinue
}

func (s *stepCreateImage) Cleanup(state multistep.StateBag) {}

