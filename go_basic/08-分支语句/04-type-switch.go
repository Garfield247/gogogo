package main

import "fmt"

func main() {
	// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}
	switch i := x.(type) {

	case nil:
		fmt.Printf("x 的类型是:%T", i)
	case int:

		fmt.Printf("x 的类型是:%T", i)
	case float64:

		fmt.Printf("x 的类型是:%T", i)
	case bool, string:

		fmt.Printf("x 的类型是:%T", i)
	default:

		fmt.Printf("x 的类型是:%T", i)
	}
}
