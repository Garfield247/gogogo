package main

import "fmt"

func main() {

	// 匿名结构体只需要声明数据类型就可以了
	// 默认以数据类型作为字段名称
	// 数据类型不允许重复
	p1 := Person{"xiaoming", 11, true}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	fmt.Println(p1.bool)
	p1.string = "xiaogang"
	fmt.Println(p1.string)

}

type Person struct {
	string
	int
	bool
}
