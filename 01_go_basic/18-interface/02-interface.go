package main

import "fmt"

func main() {
	//所有的类型都实现了空接口,所以空接口就储存了任意的数据类型
	var a1 A = Cat{color: "black"}
	var a2 A = Person{name: "xiaoming", age: 3}
	var a3 A = "hello interface"
	var a4 A = 100

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	// 使用空接口定义可以存储任意类型的map
	m1 := make(map[string]interface{})
	m1["name"] = "map"
	m1["age"] = 1
	fmt.Println(m1)

	// 使用空接口定义可以存储任意类型的切片
	s1 := make([]interface{}, 0, 10)
	s1 = append(s1, a1, a2, a3, a4, m1)
	fmt.Println(s1)

}

type A interface {
	//定义一个接口 不包含任何的方法
}

type Cat struct {
	color string
}

type Person struct {
	name string
	age  int
}

// 将空接口作为函数的参数
func printObj(a interface{}) {
	fmt.Sprintln(a)
}
