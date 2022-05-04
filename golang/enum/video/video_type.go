//go:generate stringer -type=VideoType

package video

type VideoType int64

const (
	Nexflix VideoType = iota + 1
	AmazonPrimeVideo
	DisneyPlus
	AppleTVPlus
	Hulu
)
