package main

import "fmt"

func main() {
	// 嵌套结构体,类似于继承
	type Person struct {
		name    string
		age     int
		sex     bool
		address string
	}

	type Student struct {
		Person
		sex    string
		school string
	}

	s1 := Student{}
	s1.name = "xiaoming"
	s1.age = 19
	s1.sex = "nan"
	s1.address = "beijing"
	s1.school = "beida"
	fmt.Println(s1)

	s2 := Student{}
	s2.Person.name = "xiaogang"
	s2.Person.age = 19
	s2.Person.sex = true
	//s2.sex = "nan"
	s2.Person.address = "shanghai"
	s2.school = "jiaoda"
	fmt.Println(s2)
}
