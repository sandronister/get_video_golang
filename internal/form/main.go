package form

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sandronister/get_video_golang/pkg/text"
	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func Start() *types.Input {
	reader := bufio.NewReader(os.Stdin)
	var kind string = "v"

	fmt.Println("Informe a url do youtube: ")
	urlPath, _ := reader.ReadString('\n')
	urlPath = urlPath[:len(urlPath)-1]

	fmt.Println("Informe a pasta que deseja salvar, se não existir será criada: ")
	folderPath, _ := reader.ReadString('\n')
	folderPath = folderPath[:len(folderPath)-1]

	fmt.Println("Deseja somente o áudio? (s/n): ")
	audio, _ := reader.ReadString('\n')
	audio = audio[:len(audio)-1]

	fmt.Println("Deseja alta qualidade? (s/n): ")
	quality, _ := reader.ReadString('\n')
	quality = quality[:len(quality)-1]

	if text.Sanitize(audio) == "s" {
		kind = "a"
	}

	return &types.Input{
		Url:     urlPath,
		Path:    text.Sanitize(folderPath),
		Kind:    kind,
		Quality: text.Sanitize(quality),
	}
}
