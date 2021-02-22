package main

import "fmt"

func main() {
	//声明数组 需指明数组的大小和存储的数据类型
	var l1 [10]int
	fmt.Println(l1)
	//初始化数组中{}中的元素不能大于[]中的数字
	//var l2 = [5]int{1, 2, 4, 5}
	l2 := [5]int{1, 2, 4, 5}
	fmt.Println(l2)
	// 如果忽略[]中的数字,Go语言会根据元素的个数来设置数组的大小
	//var l3 = []int{1, 2, 3, 45, 67, 87}
	l3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(l3)

	// 指定索引
	l4 := [...]int{1: 1, 5: 2, 8: 0}
	fmt.Println(l4)

	//数组的访问
	a := l2[3]
	fmt.Println(a)

	// 数字遍历
	for i := 0; i < len(l3); i++ {
		fmt.Printf("l3[%d]:%d \n", i, l3[i])
	}
	for i, x := range l2 {
		fmt.Printf("l2[%d]:%d \n", i, x)

	}

	//多维数组
	ll1 := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(ll1)
}
