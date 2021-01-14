package main

import "fmt"

func main() {

	// 变量在声明时会获取一块内存空间,后续改变变量的值的时候更改的是该变量内存中存储的内容
	var num int
	fmt.Printf("num-value:%d,memory address:%p \n", num, &num) // num-value:0,memory address:0xc00001a0d0

	num = 100
	fmt.Printf("num-value:%d,memory address:%p \n", num, &num) // num-value:100,memory address:0xc00001a0d0

	num = 135
	fmt.Printf("num-value:%d,memory address:%p \n", num, &num) // num-value:135,memory address:0xc00001a0d0

}
