package main

import "fmt"

func main() {
	// 声明map
	// 默认map是nil map,不能用来存放键值对
	var map01 map[string]string
	fmt.Printf("map01=%v\n", map01)
	// 创建map并进行初始化
	map02 := map[string]string{"name": "小明"}
	fmt.Printf("map02=%v\n", map02)
	// map赋值
	map02["sex"] = "man"
	fmt.Printf("map02=%v\n", map02)
	// 使用make创建map
	map03 := make(map[string]int)
	map03["a"] = 1
	map03["b"] = 2
	fmt.Printf("map03=%v", map03)
}
