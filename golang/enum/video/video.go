package video

import (
	"errors"
	"fmt"
)

var ErrUnknownVideo = errors.New("Unknown Video")

type Video interface {
	String() string
}

type (
	netflix          struct{}
	amazonPrimeVideo struct{}
	disneyPlus       struct{}
	appleTVPlus      struct{}
	hulu             struct{}
)

var (
	Netflix          netflix
	AmazonPrimeVideo amazonPrimeVideo
	DisneyPlus       disneyPlus
	AppleTVPlus      appleTVPlus
	Hulu             hulu
)

func New(video string) (Video, error) {
	switch video {
	case Netflix.String():
		return Netflix, nil
	case AmazonPrimeVideo.String():
		return AmazonPrimeVideo, nil
	case DisneyPlus.String():
		return DisneyPlus, nil
	case AppleTVPlus.String():
		return AppleTVPlus, nil
	case Hulu.String():
		return Hulu, nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownVideo, video)
	}
}

func (n netflix) String() string {
	return "Netflix"
}

func (a amazonPrimeVideo) String() string {
	return "AmazonPrimeVideo"
}

func (d disneyPlus) String() string {
	return "DisneyPlus"
}

func (a appleTVPlus) String() string {
	return "AppleTVPlus"
}

func (h hulu) String() string {
	return "Hulu"
}
