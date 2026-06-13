package youtube

import (
	"fmt"

	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func getQuality(entity *types.Input) error {
	entity.SetNameVideo()
	fmt.Printf("Baixando vídeo %s em qualidade superior...\n", entity.Url)
	cmd := newCommand(
		entity.Browser,
		entity.Profile,
		"-f", "bestvideo[vcodec^=avc1]+bestaudio[ext=m4a]/best[vcodec^=avc1][ext=mp4]/best",
		"--merge-output-format", "mp4",
		"-o", entity.Path,
		entity.Url,
	)

	err := runCommand(cmd)

	if err != nil {
		return fmt.Errorf("erro ao baixar vídeo: %v", err)
	}

	fmt.Printf("Vídeo baixado com sucesso em %s\n", entity.Path)

	return nil
}
