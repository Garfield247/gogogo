package main

import "fmt"

func main() {
	p1 := Person{name: "xiaoming", age: 22}
	p1.eat()
	s1 := Student{Person: Person{name: "tom", age: 10}, school: "beida"}
	s1.eat()
	s1.study()
}

type Person struct {
	name string
	age  int
}

type Student struct {
	Person // 通过 嵌套结构体模拟继承
	school string
}

func (p Person) eat() { // 父类方法
	fmt.Printf("%s正在吃东西\n", p.name)
}

func (s Student) eat() { // 子类重写的方法
	fmt.Printf("%s正在吃炸鸡\n", s.name)
}

func (s Student) study() { // 子类新增的方法
	fmt.Printf("%s正在%s学习\n", s.name, s.school)
}
