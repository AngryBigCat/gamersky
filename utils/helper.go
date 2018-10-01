package utils

import (
	"time"
	"log"
)

func DatetimeToUnix(datetime string) int64 {
	parse, err := time.Parse("2006-01-02 15:04", datetime)
	if err != nil {
		log.Printf("datetime is error")
		return 0
	}
	return parse.Unix()
}