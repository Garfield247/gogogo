// Package main provides ...
package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)
	numbers = append(numbers, 0, 1, 2, 3, 4, 5)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)
	printSlice(numbers1)
	// numbers1和numbers两者不存在联系,其中一个发生变化时另一个不会发生变化
}

func printSlice(x []int) {
	fmt.Printf("len=%d | cap=%d | slice=%v\n", len(x), cap(x), x)
}
