package video

import (
	"fmt"

	"github.com/sandronister/get_video_golang/pkg/text"
)

func DownloadVideo(url, folder, quality string) (string, error) {

	path, err := getPath(url, folder)

	if err != nil {
		return "", fmt.Errorf("erro ao criar diretório: %v", err)
	}

	quality = text.Sanitize(quality)

	if quality == "S" || quality == "s" {
		return getQuality(url, path)
	}

	return getRegular(url, path)

}
