package main

import "fmt"

// 比较两个整数的大小

func main() {
	num1, num2 := 12, 12
	max(num1, num2)
}

func max(num1, num2 int) {
	if num1 == num2 {
		fmt.Println("num1 = num2 ")
		return
	}
	if num1 > num2 {
		fmt.Printf("max:num1-%d\n", num1)
		return
	} else {
		fmt.Printf("max:num2-%d\n", num2)
		return
	}
}
