package main

import "fmt"

func main() {
	/*
		算数运算法 + - * / % ++ --


		在golang中 a/b取商 在Python中a/b是取值 a//b才是取商   a%b都相同
	*/

	a := 10
	b := 3

	sum := a + b
	fmt.Printf("%T-%v\n", sum, sum)
	sub := a - b
	fmt.Printf("%T-%v\n", sub, sub)
	mul := a * b
	fmt.Printf("%T-%v\n", mul, mul)
	div := a / b
	fmt.Printf("%T-%v\n", div, div)
	mod := a % b
	fmt.Printf("%T-%v\n", mod, mod)

	f1 := 6.28
	f2 := 3.14
	f3 := f1 / f2
	fmt.Printf("%T-%v\n", f3, f3)

	// 不同类型之间不能直接运算
	//f4 := a / f1
	//fmt.Printf("%T-%v\n", f4, f4)
	f4 := float64(a) / f1
	fmt.Printf("%T-%v\n", f4, f4)

	c := 3
	c++
	fmt.Println(c)
	c--
	fmt.Println(c)

}

// int-7
// int-30
// int-3
// int-1
// float64-2
// float64-1.592356687898089
