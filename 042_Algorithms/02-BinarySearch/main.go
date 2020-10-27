package main

import "fmt"

var (
	arr = []int{1, 3, 5, 7, 9, 10, 11}
)

func main() {
	found := -1
	x := 11
	begin := 0
	end := len(arr) - 1
	for end >= begin {
		mid := begin + int((end-begin)/2)
		if x == arr[mid] {
			found = mid
			break
		} else if x > arr[mid] {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Println(found)
}
