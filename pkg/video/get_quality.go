package video

import (
	"fmt"
	"os"
	"os/exec"
)

func getQuality(url, path string) (string, error) {
	fmt.Printf("Baixando vídeo %s em qualidade 720p...\n", url)
	cmd := exec.Command("yt-dlp", "-f", "bestvideo[height=720]+bestaudio/best[height=720]", "-o", path, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return "", fmt.Errorf("erro ao baixar vídeo: %v", err)
	}

	fmt.Printf("Vídeo baixado com sucesso em %s\n", path)

	return path, nil
}
