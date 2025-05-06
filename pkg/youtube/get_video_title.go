package youtube

import (
	"fmt"
	"os/exec"

	"github.com/sandronister/get_video_golang/pkg/text"
)

func getVideoTitle(url string) (string, error) {
	cmd := exec.Command("yt-dlp", "--get-title", url)

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("erro ao obter título do vídeo: %v", err)
	}

	title := string(output)
	title = text.Sanitize(title)

	return title, nil
}
