package gofuck

import (
	"log"
	"time"
)

func GetCurrentTimeIndonesia() time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("Error loading location:", err)
		return time.Now()
	}
	return time.Now().In(loc)
}
