package helper

import "time"

func GetTimeNow() int {
	return int(time.Now().Unix())
}
