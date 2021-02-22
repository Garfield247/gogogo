package main

import "fmt"

//堆栈的延迟
// 当一个函数有多个延迟调用时,他们将会被添加到一个堆栈当中,并以后进先出的方式执行
func main() {
	str1 := "hello world !"
	for _, n := range []rune(str1) {
		defer fmt.Printf("%c", n)
	}
}
