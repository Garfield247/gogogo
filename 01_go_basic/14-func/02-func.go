package main

import "fmt"

// 可变参数(不定长参数S
func main() {
	num := sum(10, 32, 34, 45, 453, 4)
	fmt.Printf("sum=%d\n", num)
}

func sum(arg ...int) int {
	fmt.Printf("args:%v", arg)
	var num int = 0

	for _, n := range arg {
		num += n
	}

	return num
}
