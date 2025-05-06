package main

import (
	"fmt"

	"github.com/sandronister/get_video_golang/internal/form"
	"github.com/sandronister/get_video_golang/pkg/youtube"
)

func main() {

	entity := form.Start()

	err := youtube.Download(entity)
	if err != nil {
		fmt.Println("Error downloading video:", err)
		return
	}
	fmt.Println("Download efetuado com sucesso:", entity)

}
