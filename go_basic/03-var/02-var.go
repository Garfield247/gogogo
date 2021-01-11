package main

// 多变量声明
import "fmt"

func main() {
	// 1. 以逗号分隔,声明与赋值分开,若不赋值使用默认值
	var num1, num2, num3 int
	num1, num2, num3 = 1, 2, 3
	fmt.Println(num1, num2, num3)
	// 2. 直接赋值,变量的类型可以是不同类型
	var str1, num4, bol1 = "str1", 4, false
	fmt.Println(str1, num4, bol1)
}
