package track_time_test

import (
	"api/util/track_time"
	"testing"
	"time"
)

func TestTrackTime(t *testing.T) {
	defer track_time.TrackTime(time.Now())

	time.Sleep(1 * time.Second)
}
