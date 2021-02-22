package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)
	// 允许追加空的切片
	numbers = append(numbers, 0)
	printSlice(numbers)
	// 向切片追加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers)
	numbers = append(numbers, 2, 3, 4, 5)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d | cap=%d | slice=%v\n", len(x), cap(x), x)
}
