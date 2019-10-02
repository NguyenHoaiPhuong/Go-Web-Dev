package main

import (
	"fmt"
)

func print(arr *[]int) {
	slice := *arr
	fmt.Println(slice)
	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))
	fmt.Printf("address of slice %p\n", arr)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("address of element index %d: %p\n", i, &slice[i])
	}
}

func main() {
	arr := []int{1}
	print(&arr)

	arr = append(arr, 2)
	print(&arr)

	arr = append(arr, 3)
	print(&arr)

	arr = append(arr, 4)
	print(&arr)

	arr = append(arr, 5)
	print(&arr)

	arr = append(arr, 6)
	print(&arr)

	arr = append(arr, 7)
	print(&arr)

	arr = append(arr, 8)
	print(&arr)
}
