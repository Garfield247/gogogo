package main

import "fmt"

func main() {
	// 字符串:多个byte的集合
	// 字符串的定义
	var s1 string
	s1 = "字符串"
	fmt.Printf("%T-%s\n", s1, s1) // string-字符串

	s2 := `这也是字符串`
	fmt.Printf("%T-%s\n", s2, s2) // string-这也是字符串

	// 单引号和双引号的区别
	v1 := 'a'
	v2 := "a"
	fmt.Printf("%T-%v\n", v1, v1) // int32-97
	fmt.Printf("%T-%v\n", v2, v2) // string-a
}
