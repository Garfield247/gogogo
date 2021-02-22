package main

import "fmt"

func main() {
	s1 := "abcdef"
	// 字符串转换为切片
	slice1 := []byte(s1)
	fmt.Printf("%v\n", slice1)

	slice2 := make([]byte, len(slice1))
	for i, b := range slice1 {
		slice2[i] = b - 26
	}
	fmt.Printf("%v\n", slice2)

	s2 := string(slice2)

	fmt.Printf("%v\n", s2)
}
