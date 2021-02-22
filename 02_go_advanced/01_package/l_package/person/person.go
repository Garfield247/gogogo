// Package person provides ...
package person

import "fmt"

type Person struct {
	Name string // 大写开头可以被外部访问
	age  int    // 小写开头不能被外部访问
	Sex  string
}

func (p Person) Info() {
	fmt.Printf("Name:%s|age:%d|Sex:%s\n", p.Name, p.age, p.Sex)
}
