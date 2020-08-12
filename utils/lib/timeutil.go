package lib

import "time"

const FORMAT_1 = "2006-01-02 15:04:05"

/**
获取当前时间
*/
func GetTimeStr() string {
	now := time.Now().Format(FORMAT_1)
	return now
}
