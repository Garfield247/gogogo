package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串的数据类型转换

	// string <-> bool
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T-%v\n", b1, b1)

	bb1 := true
	ss1 := strconv.FormatBool(bb1)
	fmt.Printf("%T-%v\n", ss1, ss1)

	// string <-> int

	s2 := "100"
	// 要转化的字符串,进制,位数
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T-%v\n", i2, i2)
	var ii2 int64 = 100
	ss2 := strconv.FormatInt(ii2, 10)

	fmt.Printf("%T-%v\n", ss2, ss2)

	// 大多数时候使用Atoi/Itoa
	s3 := "100"
	i3, err := strconv.Atoi(s3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T-%v\n", i3, i3)
	ss3 := strconv.Itoa(i3)
	fmt.Printf("%T-%v\n", ss3, ss3)

}
