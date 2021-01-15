package main

import "fmt"

func main() {
	// int之间的转换
	var a int8
	a = 10
	var b int16
	// b=a // 这种是不允许的
	b = int16(a)
	fmt.Printf("%T-%v\n", a, a) // int8-10
	fmt.Printf("%T-%v\n", b, b) // int16-10

	// int/float转换 float->int向下取整
	f := 3.54
	var c int
	c = int(f)
	fmt.Printf("%T-%v\n", c, c) // int-3

	// bool不能和int转化
	//var d = true
	//var g int8
	// g = int8(d) // cannot convert d (type bool) to type int8

	// int和byte
	var b1 = 'a'
	var i1 int
	i1 = int(b1)
	fmt.Printf("%T-%v\n", i1, i1) // int-97
}
