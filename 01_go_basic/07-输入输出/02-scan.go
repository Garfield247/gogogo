package main

import "fmt"

func main() {
	// scan
	var x int
	var y float64
	var z string
	fmt.Println("分别输入int,float,String;以空格间隔:")
	// 以换行为结尾
	//fmt.Scanln(&x, &y, &z)
	fmt.Scan(&x, &y, &z)
	fmt.Printf("输入的值为:%v,%v,%v\n", x, y, z)
}
