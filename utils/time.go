package utils

import "time"

var timeLocation, _ = time.LoadLocation("Asia/Shanghai")

func CurrentTime() time.Time {
	return time.Now().In(timeLocation)
}
