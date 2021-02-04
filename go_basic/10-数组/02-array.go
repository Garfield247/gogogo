package main

import "fmt"

func main() {
	// 数组是值类型,不是引用类型
	a := [...]string{"aaa", "bbb", "ccc", "ddd"}
	b := a
	b[0] = "AAA"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
