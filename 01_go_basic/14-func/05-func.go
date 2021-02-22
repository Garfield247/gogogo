package main

import "fmt"

// 引用传递
func main() {
	addOne1 := func(a int) {
		a += 1
	}
	addOne2 := func(a *int) {
		*a += 1
	}

	addOne3 := func(b int) int {
		b += 1
		return b
	}
	addOne4 := func(b *int) int {
		*b += 1
		return *b
	}
	num1, num2, num3, num4 := 10, 10, 10, 10

	addOne1(num1)
	fmt.Printf("num1=%d\n", num1)
	addOne2(&num2)
	fmt.Printf("num2=%d\n", num2)
	num5 := addOne3(num3)
	fmt.Printf("num3=%d,num5=%d\n", num3, num5)
	num6 := addOne4(&num4)
	fmt.Printf("num4=%d,num6=%d\n", num4, num6)
}

//num1=10
//num2=11
//num3=10,num5=11
//num4=11,num6=11
