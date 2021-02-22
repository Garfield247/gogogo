package main

import "fmt"

func main() {
	/* 基本数据类型
	布尔类型: bool
		取值:true false
	数值类型: 整数/浮点/复数
	字符串: string
	*/

	// 布尔类型 bool 默认值 false
	var b1 bool
	fmt.Printf("%T-%v\n", b1, b1) // bool-false
	b1 = true
	fmt.Printf("%T-%v\n", b1, b1) // bool-true
	b2 := false
	fmt.Printf("%T-%v\n", b2, b2) // bool-false

	// 在golang中bool类型不能通过0,1来进行赋值
	// b2 = 0 // cannot use 0 (type untyped int) as type bool in assignment

	// 数值类型
	// int/uint 根据底层平台代表32位或者64位整数
	// 整数 int8/int16/int32/int64 uint8/uint16/uint32/uint64 int有符号unint无符号
	// byte<->uint8 rune<->int32
	// int8    -128~127 | uin8 0~255 |2^8
	var numint8 int8
	numint8 = 100
	fmt.Printf("%T-%v\n", numint8, numint8) // int8-100

	// int16 -32768~32767 | uint16 0~65535 |2^16
	var numint16 int16
	numint16 = 22222
	fmt.Printf("%T-%v\n", numint16, numint16) // int16-22222

	// int32 -2147483648~2147483647 | uint32 0~4294967295 | 2^32
	// int64 -2^63~2^63-1 | uint64  0~2^64-1 | 2^64
	// numint16 = numint8 // cannot use numint8 (type int8) as type int16 in assignment

	// golang 是强类型,通情况下intN之间不能相互赋值,但是(byte,uint8)/(rune/int32)可以互换

	var b3 byte
	b3 = 10
	var b4 uint8
	b4 = b3
	fmt.Printf("%T-%v\n", b4, b4) // uin8-10

	// 浮点类型
	//float32
	var f32 float32
	f32 = 3.14
	var f64 float64
	f64 = 6.28
	fmt.Printf("%T-%v\n", f32, f32)   // float32-3.14
	fmt.Printf("%T-%.3f\n", f32, f32) // float32-3.140
	fmt.Printf("%T-%.6f\n", f32, f32) // float32-3.140000
	fmt.Printf("%T-%v\n", f64, f64)   // float64-6.28

}
