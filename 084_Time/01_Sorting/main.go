package main

import (
	"fmt"
	"sort"
	"time"
)

// QtyAtTime struct
type QtyAtTime struct {
	date     time.Time
	quantity float64
}

func main() {
	var demands []QtyAtTime
	demands = append(demands,
		QtyAtTime{Date(2019, 1, 3), 80},
		QtyAtTime{Date(2019, 1, 1), 90},
		QtyAtTime{Date(2019, 1, 2), 70},
		QtyAtTime{Date(2019, 1, 4), 100},
	)

	fmt.Println("Before Sorting:")
	for _, demand := range demands {
		fmt.Println(demand.date, demand.quantity)
	}

	sort.Slice(demands, func(i, j int) bool {
		return demands[i].date.Before(demands[j].date)
	})

	fmt.Println("After Sorting:")
	for _, demand := range demands {
		fmt.Println(demand.date, demand.quantity)
	}
}

// Date returns date
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
