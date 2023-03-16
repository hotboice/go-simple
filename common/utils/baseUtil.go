package utils

import "time"

func Timetmp() int {
	return time.Now().Second()
}
