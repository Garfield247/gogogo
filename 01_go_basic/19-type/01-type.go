package main

import "fmt"

func main() {
	var s1 mystr = "hello"
	var n1 myint = 123
	info(s1)
	info(n1)
}

type myint int
type mystr string

func info(v interface{}) {
	fmt.Printf("%v|%T|%p\n", v, v, &v)
}
