package main

// 多变量声明
import "fmt"

func main() {
	// 1. 以逗号分隔,声明与赋值分开,若不赋值使用默认值
	// 多个同类型的变量声明
	var num1, num2, num3 int
	num1, num2, num3 = 1, 2, 3
	fmt.Println(num1, num2, num3)
	// 同时声明多种类型
	var (
		a1 int
		b1 string
		c1 bool
	)
	a1, b1, c1 = 1, "b1", true
	fmt.Println(a1, b1, c1)
	// 2. 直接赋值,变量的类型可以是不同类型
	var str1, num4, bol1 = "str1", 4, false
	fmt.Println(str1, num4, bol1)
	var (
		a = 10
		b = "sdfasas"
		c = false
	)
	fmt.Println(a, b, c)
}
