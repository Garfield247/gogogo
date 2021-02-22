package main

import "fmt"

func main() {

	//创建map

	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	printMap(m1)

	// 删除元素
	fmt.Println("删除元素")
	delete(m1, "a")
	printMap(m1)
}

func printMap(m map[string]int) {

	//遍历打印
	for k, v := range m {
		fmt.Printf("%v:%v\n", k, v)
	}
}
