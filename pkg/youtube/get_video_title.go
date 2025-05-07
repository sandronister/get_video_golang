package youtube

import (
	"os/exec"

	"github.com/sandronister/get_video_golang/pkg/text"
)

func getVideoTitle(url string) string {
	cmd := exec.Command("yt-dlp", "--get-title", url)

	output, err := cmd.Output()
	if err != nil {
		return "video-title"
	}

	title := string(output)
	title = text.Sanitize(title)

	return title
}
