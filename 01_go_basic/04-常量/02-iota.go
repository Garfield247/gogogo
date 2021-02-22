package main

import "fmt"

func main() {
	// iota 特殊的常量 可以被编译器"修改"的常量
	// 每当定义一个const ,iota的初始值为0
	// 每当定义一个常量 iota就自动累加1
	// 直到下一个const出现清0
	// b->f没有赋值默认跟上面的相同  = iota
	const (
		a = iota
		b
		c
		d
		e
		f
	)
	fmt.Println(a, b, c, d, f) // 0 1 2 3 4

	// iota 的变化
	const (
		A = iota  // 0
		B         // 1
		C         // 2
		D = "abc" // abc iota = 3
		E         // abc iota = 4
		F         // abc iota = 5
		G = 100   // 100 iota=6
		H = iota  // 7
		I         // 8
	)
	const (
		J = iota // 0
	)
	fmt.Println(A, B, C, D, E, F, G, H, I, J)

	// iota另类使用
	const (
		A1 = iota * 10 // 0
		B1             // 10
		C1             // 20
	)
	fmt.Println(A1, B1, C1)
	const (
		a1, b1 = iota + 1, iota + 2 // 1,2
		c1, d1                      // 2,3
		e1, f1                      // 3,4
	)
	fmt.Println(a1, b1, c1, d1, e1, f1)
}
