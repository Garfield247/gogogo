package main

import "fmt"

func main() {
	// 定义切片
	//var s0 []int
	// 使用make创建切片
	//var s1 []int = make([]int,2)
	//s2:= make([]int,2)

	a := [5]int{1, 2, 3, 4, 5}
	var s1 []int = a[1:4]
	fmt.Println(s1)
	//切片是引用类型,对切片进行的操作会作用在源数组上
	for i := range s1 {
		s1[i] += 100
	}
	fmt.Println(s1)
	fmt.Println(a)

	// len获取切片长度 cap获取切片最大长度
	s2 := make([]int, 3, 5)

	PrintSlice(s2)

	PrintSlice(s1)
}

func PrintSlice(slice []int) {
	fmt.Printf("len:%d;cap:%d;slice:%v \n", len(slice), cap(slice), slice)
}
