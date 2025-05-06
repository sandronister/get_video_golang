package youtube

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func getAudio(entity *types.Input) error {
	entity.SetNameAudio()
	fmt.Printf("Baixando áudio de %s em qualidade regular...\n", entity.Url)
	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-o", entity.Path, entity.Url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("erro ao baixar áudio: %v", err)
	}

	return nil
}
