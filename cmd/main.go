package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sandronister/get_video_golang/pkg/video"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the video URL: ")
	urlPath, _ := reader.ReadString('\n')
	urlPath = urlPath[:len(urlPath)-1]

	fmt.Print("Enter the folder path to save the video: ")
	folderPath, _ := reader.ReadString('\n')
	folderPath = folderPath[:len(folderPath)-1]

	result, err := video.DownloadVideo(urlPath, folderPath)
	if err != nil {
		fmt.Println("Error downloading video:", err)
		return
	}
	fmt.Println("Video downloaded successfully:", result)

}
