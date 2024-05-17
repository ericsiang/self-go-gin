package time_test

import (
	"api/util/time_relate"
	"testing"
	"time"
)

func TestTrackTime(t *testing.T) {
	defer time_relate.TrackTime(time.Now())

	time.Sleep(1 * time.Second)
}
