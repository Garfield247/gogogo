package main

import "fmt"

func main() {
	// 常量:程序执行过程中数值不能改变
	// - 常量定义后不使用不会报错
	// - 常量定义后不能更改
	// - 常量名使用大大写
	//常数:固定的数值
	fmt.Println(100)
	fmt.Println("adc")
	//显示定义
	const NAME string = "constant"
	//NAME = "aaa" // 04-constant/01-canstant.go:15:7: cannot assign to NAME
	// 隐式定义
	const PI = 3.1415926
	// 批量定义
	const C1, C2, C3 = 1, 2, 3
	const (
		C4 int    = 5
		C5 string = "abc"
		C6 bool   = false
	)
	// 在定义一组常量时若未赋值,则默认和上一个相同
	const (
		C7 int = 100
		C8
		C9
		C10 string = "qwer"
		C11
		C12
	)

	fmt.Printf("%T-%v\n", C7, C7)
	fmt.Printf("%T-%v\n", C8, C8)
	fmt.Printf("%T-%v\n", C9, C9)
	fmt.Printf("%T-%v\n", C10, C10)
	fmt.Printf("%T-%v\n", C11, C11)
	fmt.Printf("%T-%v\n", C12, C12)
	// 输出值:
	// int-100
	// int-100
	// int-100
	// string-qwer
	// string-qwer
	// string-qwer

	// 枚举类型
	const (
		A = 1
		B = 2
		C = 3
		D = 4
	)
}
