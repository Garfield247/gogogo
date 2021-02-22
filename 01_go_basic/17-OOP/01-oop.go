package main

import "fmt"

// 结构体的方法

func main() {

	w1 := Worker{name: "xiaoming", age: 22, sex: "nan"}
	w1.info()
	w1.work()

}

type Worker struct {
	name string
	age  int
	sex  string
}

func (w Worker) info() {
	fmt.Printf("name:%s,\nage:%d,\nsex:%s\n", w.name, w.age, w.sex)
}

func (w Worker) work() {
	fmt.Printf("%s在工作\n", w.name)
}
