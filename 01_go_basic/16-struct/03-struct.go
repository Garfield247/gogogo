package main

import "fmt"

func main() {

	// 匿名结构体

	s1 := struct {
		name string
		age  int
	}{
		name: "xiaoming",
		age:  10,
	}

	fmt.Println(s1)
}
