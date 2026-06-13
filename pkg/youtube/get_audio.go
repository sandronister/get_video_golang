package youtube

import (
	"fmt"

	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func getAudio(entity *types.Input) error {
	entity.SetNameAudio()
	fmt.Printf("Baixando áudio de %s em qualidade regular...\n", entity.Url)
	fmt.Printf("Salvando em %s\n", entity.Path)
	cmd := newCommand(
		entity.Browser,
		entity.Profile,
		"-f", "bestaudio",
		"-x", "--audio-format", "mp3",
		"-o", entity.Path,
		entity.Url,
	)

	err := runCommand(cmd)

	if err != nil {
		return fmt.Errorf("erro ao baixar áudio: %v", err)
	}

	return nil
}
