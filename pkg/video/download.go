package video

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func DownloadVideo(url, folder string) (string, error) {

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return "", err
	}

	path := path.Join(folder, "video.mp4")

	cmd := exec.Command("yt-dlp", "-f", "mp4", "-o", path, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erro ao baixar vídeo: %v\n", err)
	} else {
		fmt.Println("Download concluído com sucesso.")
	}

	return fmt.Sprintf("Video downloaded to %s", path), nil

}
