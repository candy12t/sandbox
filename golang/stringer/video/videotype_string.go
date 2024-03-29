// Code generated by "stringer -type=VideoType"; DO NOT EDIT.

package video

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Nexflix-1]
	_ = x[AmazonPrimeVideo-2]
	_ = x[DisneyPlus-3]
	_ = x[AppleTVPlus-4]
	_ = x[Hulu-5]
}

const _VideoType_name = "NexflixAmazonPrimeVideoDisneyPlusAppleTVPlusHulu"

var _VideoType_index = [...]uint8{0, 7, 23, 33, 44, 48}

func (i VideoType) String() string {
	i -= 1
	if i < 0 || i >= VideoType(len(_VideoType_index)-1) {
		return "VideoType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _VideoType_name[_VideoType_index[i]:_VideoType_index[i+1]]
}
