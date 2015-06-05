package main

import (
	"time"
)


func ToDate(d int64) string{
	str_time := time.Unix(d, 0).Format("2006-01-02 15:04:05")
	return str_time
}
