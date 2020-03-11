package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	str := ToTimeStringRFC3339(t)
	fmt.Println(str)

	str = ToTimeStringISO8601(t)
	fmt.Println(str)
}
