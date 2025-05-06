package video

import (
	"fmt"
	"os"
	"path"

	"github.com/sandronister/get_video_golang/pkg/text"
)

func getPath(url, folder string) (string, error) {
	title := getVideoTitle(url)

	folder = text.Sanitize(folder)

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		fmt.Print("Erro ao criar diretório: ", err)
		return "", err
	}

	path := path.Join(folder, title)

	if _, err := os.Stat(path); err == nil {
		return path, nil
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path, nil
	}

	return "", fmt.Errorf("erro ao verificar o arquivo")
}
