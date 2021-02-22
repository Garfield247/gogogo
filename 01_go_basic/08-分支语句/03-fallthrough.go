package main

import "fmt"

// 如果需要贯通后续的case,就添加fallthrough 并且fallthrough只能出现case的最后一行
func main() {
	switch x := 5; x {
	case 5:
		x += 1
		fmt.Println(x)
		fallthrough
	case 6:
		x += 1
		fmt.Println(x)

	case 7:
		x += 1
		fmt.Println(x)
	}
}
