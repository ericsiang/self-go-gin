package time_relate

import (
	"fmt"
	"time"
)

// 取得當下的台北時間
func TimeNow() time.Time {
	now := time.Now()
	local, _ := time.LoadLocation("Asia/Taipei") //修改成台北時間
	//local, _ := time.LoadLocation("") //修改成台北時間
	return now.In(local)
}

func TrackTime(pre time.Time) time.Duration {
	expend := time.Since(pre)
	fmt.Println("expend : ", expend)
	return expend
}
