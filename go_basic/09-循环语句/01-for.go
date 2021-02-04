package main

import "fmt"

func line() {
	fmt.Println("------------------------------------")
}

func main() {
	// 普通forxunhuan
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)

	}
	line()
	// 类while用法
	a := 1
	for a < 10 {
		a++
		fmt.Println(a)
	}
	line()
	// for range
	nums := [6]int{1, 2, 3, 5}
	for i, x := range nums {
		fmt.Printf("index:%d,value:%d\n", i, x)
	}
}
