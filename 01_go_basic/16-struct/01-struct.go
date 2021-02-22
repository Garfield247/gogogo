package main

import "fmt"

func main() {

	type Person struct {
		name    string
		age     int
		sex     bool
		address string
	}

	// 方式一
	var p1 Person
	p1.name = "小明"
	p1.age = 10
	p1.sex = true
	p1.address = "beijing"

	fmt.Printf("name:%s,age:%d,sex:%t,address:%s\n", p1.name, p1.age, p1.sex, p1.address)

	//方式二
	p2 := Person{}
	p2.name = "小红"
	p2.age = 13
	p2.sex = false
	p2.address = "shanghai"

	fmt.Printf("name:%s,age:%d,sex:%t,address:%s\n", p2.name, p2.age, p2.sex, p2.address)
	//方式三
	p3 := Person{name: "小刚", age: 23, sex: true, address: "shenzhen"}

	fmt.Printf("name:%s,age:%d,sex:%t,address:%s\n", p3.name, p3.age, p3.sex, p3.address)

}
