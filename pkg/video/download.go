package video

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
)

func sanitizeFolderName(folder string) string {
	invalidChars := regexp.MustCompile(`[<>:"/\\|?*]`)
	folder = invalidChars.ReplaceAllString(folder, "_")

	folder = regexp.MustCompile(`^[\s.]+|[\s.]+$`).ReplaceAllString(folder, "")

	reservedNames := map[string]bool{
		"CON": true, "PRN": true, "AUX": true, "NUL": true,
		"COM1": true, "COM2": true, "COM3": true, "COM4": true,
		"COM5": true, "COM6": true, "COM7": true, "COM8": true,
		"COM9": true, "LPT1": true, "LPT2": true, "LPT3": true,
		"LPT4": true, "LPT5": true, "LPT6": true, "LPT7": true,
		"LPT8": true, "LPT9": true,
	}
	if reservedNames[folder] {
		folder = folder + "_folder"
	}

	return folder
}

func getPath(url, folder string) (string, error) {
	title, err := getVideoTitle(url)
	if err != nil {
		return "", fmt.Errorf("erro ao obter título do vídeo: %v", err)
	}

	folder = sanitizeFolderName(folder)

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

func getVideoTitle(url string) (string, error) {
	cmd := exec.Command("yt-dlp", "--get-title", url)

	// Captura a saída diretamente
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("erro ao obter título do vídeo: %v", err)
	}

	// Remove espaços em branco extras
	title := string(output)
	title = sanitizeFolderName(title)

	return title, nil
}

func getVideo(url, path string) (string, error) {

	cmd := exec.Command("yt-dlp", "-f", "mp4", "-o", path, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return "", fmt.Errorf("erro ao baixar vídeo: %v", err)
	}

	return path, nil
}

func DownloadVideo(url, folder string) (string, error) {

	path, err := getPath(url, folder)

	if err != nil {
		return "", fmt.Errorf("erro ao criar diretório: %v", err)
	}

	fmt.Printf("Baixando vídeo de %s para %s\n", url, path)

	return getVideo(url, path)

}
