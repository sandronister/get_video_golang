package video

import (
	"fmt"
	"os"
	"os/exec"
)

func getRegular(url, path string) (string, error) {
	fmt.Printf("Baixando vídeo %s em qualidade regular...\n", url)
	cmd := exec.Command("yt-dlp", "-f", "mp4", "-o", path, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return "", fmt.Errorf("erro ao baixar vídeo: %v", err)
	}

	return path, nil
}
