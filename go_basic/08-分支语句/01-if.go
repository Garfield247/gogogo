package main

import "fmt"

func main() {
	// if语句
	var num int = 100
	if num < 10 {
		fmt.Println("小于10")
	} else {
		fmt.Println("不小于10")
	}
	// 判断时声明要判断的内容
	if num2 := 100; num2 < 10 {
		fmt.Println("小于10")
	} else {
		fmt.Println("不小于10")
	}
}
