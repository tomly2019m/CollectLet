package util

import "time"

func getTimeStamp() int64 {
	return time.Now().UnixMilli()
}
