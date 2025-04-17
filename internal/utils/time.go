package utils

import "time"

func NowTime() string {
	return time.Now().Format(time.DateTime)
}

func NowUnix() int64 {
	return time.Now().Unix()
}
