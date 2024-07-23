package util

import "time"

func GetTimeStamp() int64 {
	return time.Now().UnixMilli()
}

func GetCurrentTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}
