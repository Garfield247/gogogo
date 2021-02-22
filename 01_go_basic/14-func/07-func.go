package main

import "fmt"

// 延迟参数
func main() {
	a := 5
	defer fmt.Printf("defer a=%d\n", a)
	a = 10

	fmt.Printf(" a=%d\n", a)
}

//a=10
//defer a=5
