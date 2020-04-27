package main

import (
	"fmt"
	"time"
)

func toTimeString(t time.Time) {
	str := ToTimeStringRFC3339(t)
	fmt.Println(str)

	str = ToTimeStringISO8601(t)
	fmt.Println(str)
}

func formatTime(t time.Time) {
	str := t.Format("2006-01-02 15:04:05")
	fmt.Println(str)

	str = t.Format("2006-01-02")
	fmt.Println(str)
}

func main() {
	t := time.Now()

	toTimeString(t)
	formatTime(t)
}
