package youtube

import (
	"fmt"

	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func Download(entity *types.Input) error {

	err := getPath(entity)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	if entity.IsAudio() {
		return getAudio(entity)
	}

	if entity.IsQuality() {
		return getQuality(entity)
	}

	return getRegular(entity)

}
