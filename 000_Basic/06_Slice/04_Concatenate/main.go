package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("address of arr %p\n", &arr)
	fmt.Printf("arr point to %p\n", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("address of element index %d: %p\n", i, &arr[i])
	}

	a := arr[0:1]
	fmt.Println("a:", a, ", size:", len(a), ", cap:", cap(a))
	fmt.Printf("address of a %p\n", &a)
	fmt.Printf("a point to %p\n", a)

	b := arr[2:]
	fmt.Println("b:", b, ", size:", len(b), ", cap:", cap(b))
	fmt.Printf("address of b %p\n", &b)
	fmt.Printf("b point to %p\n", b)

	newArr := append(a, b...)
	fmt.Println("newArr:", newArr, ", size:", len(newArr), ", cap:", cap(newArr))

	fmt.Println("a:", a, ", size:", len(a), ", cap:", cap(a))
	fmt.Printf("address of a %p\n", &a)
	fmt.Printf("a point to %p\n", a)

	fmt.Println("b:", b, ", size:", len(b), ", cap:", cap(b))
	fmt.Printf("address of b %p\n", &b)
	fmt.Printf("b point to %p\n", b)
}
