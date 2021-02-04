// Package main  provides ...
package main

import "fmt"

func main() {
	m1 := make(map[string]string)
	m1["a"] = "aaaaa"
	m1["b"] = "bbbbb"
	m1["c"] = "ccccc"
	m1["d"] = "ddddd"
	fmt.Printf("m1:%v\n", m1)

	// 使用key 输出map值
	for key := range m1 {
		fmt.Printf("m1-%v:%v\n", key, m1[key])
	}
	// 遍历map
	for key, v := range m1 {
		fmt.Printf("m1-%v:%v\n", key, v)
	}

	// 判断元素是否存在
	//v, ok := m1["e"]
	v, ok := m1["c"]
	if ok {
		fmt.Printf("存在 ,v = %v", v)
	} else {
		fmt.Println("不存在")
	}
}
