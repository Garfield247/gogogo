package main

import "fmt"

func main() {
	/*
		逻辑运算符
	*/

	b1, b2, b3 := true, false, true
	// 逻辑与 &&
	r1 := b1 && b2
	r2 := b1 && b3
	fmt.Println(r1, r2)
	// 逻辑与 ||
	r3 := b1 || b2
	r4 := b1 || b3
	fmt.Println(r3, r4)
	// 逻辑与 !
	r5 := !b1
	r6 := !b2
	fmt.Println(r5, r6)
}
