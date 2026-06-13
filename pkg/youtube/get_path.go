package youtube

import (
	"fmt"
	"os"
	"path"

	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func getPath(entity *types.Input) error {
	title := getVideoTitle(entity.Url, entity.Browser, entity.Profile)

	if err := os.MkdirAll(entity.Path, os.ModePerm); err != nil {
		fmt.Print("Erro ao criar diretório: ", err)
		return err
	}

	entity.Path = path.Join(entity.Path, "/", title)

	fmt.Println("Pasta criada com sucesso: ", entity.Path)

	return nil
}
