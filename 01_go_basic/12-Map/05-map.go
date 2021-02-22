package main

import "fmt"

func main() {

	// 和切片相似映射是引用类型,将映射分配为一个新的变量时,它们都同时指向相同的内部数据结构,
	//因此对一个进行变更另外一个也会随之变化
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	printMap(m1)
	m2 := m1
	printMap(m2)
	m2["d"] = 400
	printMap(m1)
	printMap(m2)
}

func printMap(m map[string]int) {
	fmt.Printf("%v\n", m)
}
