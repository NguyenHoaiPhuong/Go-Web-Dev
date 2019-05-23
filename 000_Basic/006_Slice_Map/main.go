package main

import (
	"fmt"
	"time"
)

// RemoveItemFromSlice : remove 1 specific item from slice
func RemoveItemFromSlice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	slice = append(slice[:2], slice[3:]...)
	for _, val := range slice {
		fmt.Println(val)
	}
}

func main() {
	university := new(University)
	university.Init()

	for i := 0; i < 102400; i++ {
		NewStudent(university)
	}

	ti := time.Now()
	for i := 102400; i > 0; i-- {
		university.RemoveStudentFromSlice(i)
	}
	seconds := time.Now().Sub(ti).Seconds()
	fmt.Println("Remove Student From Slice took", seconds)

	ti = time.Now()
	for i := 102400; i > 0; i-- {
		university.RemoveStudentFromMap(i)
	}
	seconds = time.Now().Sub(ti).Seconds()
	fmt.Println("Remove Student From Map took", seconds)
}
