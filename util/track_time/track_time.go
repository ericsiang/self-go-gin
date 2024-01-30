package track_time

import (
	"fmt"
	"time"
)

func TrackTime(pre time.Time) time.Duration {
	expend := time.Since(pre)
	fmt.Println("expend : ", expend)
	return expend
}
