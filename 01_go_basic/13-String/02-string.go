package main

import "fmt"

func main() {
	// 遍历字符串
	s1 := "hello world !"

	for i := 0; i < len(s1); i++ {

		fmt.Printf("[%d]'%c' |", i, s1[i])
	}
	fmt.Println("\n-------------")
	for i, c := range s1 {
		fmt.Printf("%d-%c   ", i, c)
	}
}
