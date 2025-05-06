package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sandronister/get_video_golang/pkg/video"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Informe a url do youtube: ")
	urlPath, _ := reader.ReadString('\n')
	urlPath = urlPath[:len(urlPath)-1]

	fmt.Println("Informe a pasta que deseja salvar, se não existir será criada: ")
	folderPath, _ := reader.ReadString('\n')
	folderPath = folderPath[:len(folderPath)-1]

	fmt.Println("Deseja alta qualidade? (s/n): ")
	quality, _ := reader.ReadString('\n')
	quality = quality[:len(quality)-1]

	result, err := video.DownloadVideo(urlPath, folderPath, quality)
	if err != nil {
		fmt.Println("Error downloading video:", err)
		return
	}
	fmt.Println("Video downloaded successfully:", result)

}
