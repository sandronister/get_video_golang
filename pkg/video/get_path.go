package video

import (
	"fmt"
	"os"
	"path"

	"github.com/sandronister/get_video_golang/pkg/text"
)

func getPath(url, folder string) (string, error) {
	title, err := getVideoTitle(url)
	if err != nil {
		return "", fmt.Errorf("erro ao obter título do vídeo: %v", err)
	}

	folder = text.Sanitize(folder)

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		fmt.Print("Erro ao criar diretório: ", err)
		return "", err
	}

	path := path.Join(folder, title+".mp4")

	if _, err := os.Stat(path); err == nil {
		return path, nil
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path, nil
	}

	return "", fmt.Errorf("erro ao verificar o arquivo")
}
