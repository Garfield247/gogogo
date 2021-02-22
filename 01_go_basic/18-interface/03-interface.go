// Package main provides ...
package main

import "fmt"

// 接口的嵌套

func main() {
	var cat Cat
	fmt.Println("cat Cat")
	cat.test1()
	cat.test2()
	cat.test3()
	fmt.Println("_______________________")
	var a1 A = cat
	a1.test1()
	fmt.Println("_______________________")
	var a2 B = cat
	a2.test2()
	fmt.Println("_______________________")
	var a3 C = cat
	a3.test1()
	a3.test2()
	a3.test3()
}

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct {
}

func (c Cat) test1() {
	fmt.Println("test1")
}

func (c Cat) test2() {
	fmt.Println("test2")
}

func (c Cat) test3() {
	fmt.Println("test3")
}
