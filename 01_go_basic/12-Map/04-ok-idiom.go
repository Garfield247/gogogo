package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Printf("m1=%v\n", m1)
	// 我们可以通过map[key]的方法获取value对应的key
	a := m1["a"]
	fmt.Printf("a=%v\n", a)
	// 如果该key不存在时则获取到该value数据类型的默认值
	e := m1["e"]
	fmt.Printf("e=%v\n", e)
	// 可以通过 ok-idiom 获取值,并知道该key是否存在
	f, ok := m1["f"]
	fmt.Printf("f=%v | ok = %v\n", f, ok)

	//map的长度
	length := len(m1)
	fmt.Printf("m1的长度为%d", length)
}
