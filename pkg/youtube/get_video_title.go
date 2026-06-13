package youtube

import (
	"github.com/sandronister/get_video_golang/pkg/text"
)

func getVideoTitle(url, browser, profile string) string {
	cmd := newCommand(browser, profile, "--get-title", url)

	output, err := cmd.Output()
	if err != nil {
		return "video-title"
	}

	title := string(output)
	title = text.Sanitize(title)

	return title
}
