package types

import (
	"github.com/sandronister/get_video_golang/pkg/text"
)

type Input struct {
	Url     string
	Quality string
	Path    string
	Kind    string
}

func (i *Input) IsQuality() bool {
	i.Quality = text.Sanitize(i.Quality)
	return i.Quality == "S" || i.Quality == "s"
}

func (i *Input) IsAudio() bool {
	i.Kind = text.Sanitize(i.Kind)
	return i.Kind == "A" || i.Kind == "a"
}

func (i *Input) SetNameAudio() {
	if i.IsAudio() {
		i.Quality = "S"
		i.Path = i.Path + ".mp3"
	}
}

func (i *Input) SetNameVideo() {
	i.Path = i.Path + ".mp4"
}
