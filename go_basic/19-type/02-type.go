// Package main provides ...
package main

import (
	"fmt"
	"strconv"
)

// 定义函数类型
func main() {

	res1 := func1()
	fmt.Println(res1)
	res2 := res1(10, 20)
	fmt.Println(res2)
}

type myfun func(int, int) string

func func1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}
