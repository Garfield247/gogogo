package main

import "fmt"

// 多返回值
func main() {
	count, num := sum(10, 32, 213, 312, 213, 321)
	fmt.Printf("count=%d,sum=%d\n", count, num)
}

// 返回值声明名称及类型
func sum(arg ...int) (count int, num int) {
	fmt.Printf("args:%v", arg)
	count = len(arg)
	num = 0

	for _, n := range arg {
		num += n
	}

	return count, num
}

//返回值只声明类型
func sum2(arg ...int) (int, int) {
	fmt.Printf("args:%v", arg)
	count := len(arg)
	num := 0

	for _, n := range arg {
		num += n
	}

	return count, num
}
