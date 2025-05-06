package youtube

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func getQuality(entity *types.Input) error {
	fmt.Printf("Baixando vídeo %s em qualidade superior...\n", entity.Url)
	cmd := exec.Command("yt-dlp", "-f", "bestvideo+bestaudio/best", "-o", entity.Path, entity.Url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("erro ao baixar vídeo: %v", err)
	}

	fmt.Printf("Vídeo baixado com sucesso em %s\n", entity.Path)

	return nil
}
