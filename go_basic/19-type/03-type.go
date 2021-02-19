package main

import "fmt"

func main() {

	var n1 myint = 10
	var n2 int64
	//使用类型别名时可以直接进行赋值
	n2 = n1
	fmt.Println(n1, n2)
}

// 起类型别名   go1.9版本开始
//1.9之前
//type byte uint8
//type rune int32
//1.9开始
//type byte = uint8
//type rune = int32
type myint = int64
