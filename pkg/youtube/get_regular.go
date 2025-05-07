package youtube

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func getRegular(entity *types.Input) error {
	entity.SetNameVideo()
	fmt.Printf("Baixando vídeo %s em qualidade regular...\n", entity.Url)
	cmd := exec.Command("yt-dlp", "-f", "mp4", "-o", entity.Path, entity.Url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("erro ao baixar vídeo: %v", err)
	}

	return nil
}
