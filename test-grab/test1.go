package main

import (
	"fmt"
	"sort"
)

type Ints []int

// O(nlog(n))

// O(klog(k))

// {2, 5}

// O(k)

// 1 -> ko bo vo

// 4 bo vo

// (n - k)klog(k)

// k = n/2

var (
	mainArr = Ints{2, 5, 1, 4, 3}
	subArr  Ints
	k       = 2
)

O(1) O(log2(n)) O(n) O(n*log2(n)) O(n2)

// Len func
func (arr Ints) Len() int {
	return len(arr)
}

// Less func
func (arr Ints) Less(i, j int) bool {
	return arr[i] < arr[j]
}

// Swap func
func (arr Ints) Swap(i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// Sort func
func (arr Ints) Sort() Ints {
	sort.Sort(arr)
	return arr
}

input = []int{1, 5, 5, 3}
expectedOutput

func MaxSum(arr Ints, k int) int {

}

func main() {
	mainArr.Sort()

	sum := 0
	num := mainArr.Len()
	if k > num {
		fmt.Println("Error")
		return
	}
	for i := 0; i < k; i++ {
		sum += mainArr[num-i-1]
	}
	fmt.Println(sum)
}
