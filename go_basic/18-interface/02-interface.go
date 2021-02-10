package main

import "fmt"

func main() {
	var a1 A = Cat{color: "black"}
	var a2 A = Person{name: "xiaoming", age: 3}
	var a3 A = "hello interface"
	var a4 A = 100

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

}

type A interface {
	//定义一个接口 不包含任何的方法
	//所有的类型都实现了空接口,所以空接口就储存了任意的数据类型
}

type Cat struct {
	color string
}

type Person struct {
	name string
	age  int
}
