package main

//单变量声明
import "fmt"

func main() {
	// var 类型是声明单个变量的语法
	// 声明变量的语法通常有三种
	// 1.指定变量类型
	var num1 int
	var str1 string
	//0未赋值将使用默认值
	fmt.Println(num1, "|", str1) //0 |
	var num2 int = 10
	var str2 string = "字符串"
	fmt.Println(num2, "|", str2) //10 | 字符串
	// 2. 根据值自行判断数据类型
	var num3 = 10
	var str3 = "str3"
	fmt.Println(num3, "|", str3) // 10 | str3
	//类型推导
	fmt.Printf("num3-type:%T-value:%d\n", num3, num3)
	fmt.Printf("str3-type:%T-value:%s\n", str3, str3)
	// 3. 省略var !这种方式只能被用在函数体内不能用于全局变量的声明与赋值
	num4 := 4
	str4 := "str4"
	fmt.Println(num4, "|", str4) // 4 | str4
}
