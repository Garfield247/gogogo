// Package utils provides ...
package utils

import "fmt"

func init() {
	fmt.Println("utils/init FUNC1")
}

// init 函数可以在一个文件中存在多个,按照顺序执行
func init() {
	fmt.Println("utils/init FUNC2")
}

func Count() {
	fmt.Println("utils/util.go Count() function")
}
