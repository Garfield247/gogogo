package main

import "fmt"

// fmt 包实现了类似C语言printf和scanf的格式化I/O.
// 格式化动作('verb')源自C语言但更简单
func main() {
	// 普通打印
	fmt.Print("Print打印内容不换行,换行要加\\n\n")

	// 打印后换行
	fmt.Println("Println打印内容并自动换行")

	// 格式化打印,不会自动换行
	str1 := "字符串"
	num1 := 100
	flo1 := 3.14
	bool1 := true
	// %v原样输出
	fmt.Printf("str的值为:%v\n", str1)
	fmt.Printf("num的值为:%v\n", num1)
	fmt.Printf("flo的值为:%v\n", flo1)
	fmt.Printf("bool的值为:%v\n", bool1)
	// %T打印内容
	fmt.Printf("str的类型为:%T\n", str1)
	fmt.Printf("num的类型为:%T\n", num1)
	fmt.Printf("flo的类型为:%T\n", flo1)
	fmt.Printf("bool的类型为:%T\n", bool1)
	// %p打印内存地址,需要传入指针
	fmt.Printf("str的内存地址为:%p\n", &str1)
	fmt.Printf("num的内存地址为:%p\n", &num1)
	fmt.Printf("flo的内存地址为:%p\n", &flo1)
	fmt.Printf("bool的内存地址为:%p\n", &bool1)
	// %s打印字符串
	fmt.Printf("str的值为:%s\n", str1)
	// %d打印int
	fmt.Printf("num的值为:%d\n", num1)
	// %f打印浮点数
	fmt.Printf("flo的值为:%f\n", flo1)
	// %.{n}f打印浮点数,确定小数点位数
	fmt.Printf("flo的值为:%.5f\n", flo1)
	// %t打印布尔类型
	fmt.Printf("bool的值为:%t\n", bool1)
	// %c打印字符
	fmt.Printf("num的值为:%c\n", num1)
	//%b 二进制
	fmt.Printf("num的二进制值为:%b\n", num1)
	//%o 八进制
	fmt.Printf("num的八进制值为:%o\n", num1)
	//%x 16进制小写
	fmt.Printf("num的16进制小写值为:%x\n", num1*1000)
	//%X 16进制大写
	fmt.Printf("num的16进制大写值为:%X\n", num1*1000)
}
