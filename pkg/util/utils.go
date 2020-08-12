package util

import "time"

func GenerateCurrentTimeStamp() int64 {
	return time.Now().Unix()
}
