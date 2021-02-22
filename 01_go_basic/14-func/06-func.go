package main

import "fmt"

// defer 延迟语句当函数执行到最后时这些defer语句会按逆序执行
func main() {
	a, b, c, d := 1, 2, 3, 4

	defer fmt.Println(a)
	defer fmt.Println(b)
	defer fmt.Println(c)
	fmt.Println(d)
}

// 4 3 2 1
