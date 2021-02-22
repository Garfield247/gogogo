package utils

import "fmt"

// 当同一个包下有多个文件中存在init函数时会按照文件名的字母排序执行
func init() {
	fmt.Println("utils/test1 init 1")

}

func Test() {
	fmt.Println("utils/test1 Test ")

}
