package utils

import (
	"time"
)

// TimestampToTime 毫秒时间戳转为时间
func TimestampToTime(ts int64) time.Time {
	if ts == 0 {
		return time.Time{}
	}
	return time.Unix(0, ts*int64(time.Millisecond))
}
