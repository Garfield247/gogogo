package main

import "fmt"

func main() {
	// var 类型是声明单个变量的语法
	// 声明变量的语法通常有三种
	// 1.指定变量类型
	var num1 int
	var str1 string
	//0未赋值将使用默认值
	fmt.Println(num1, "|", str1) //0| |0
	var num2 int = 10
	var str2 string = "字符串"
	fmt.Println(num2, "|", str2) //0| |0
}
