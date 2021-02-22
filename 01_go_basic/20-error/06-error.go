// Package main provides ...
package main

import "fmt"

func main() {

	defer myPrint("main defer1")
	funcA()
	defer myPrint("main defer2")
	funcB()
	defer myPrint("main defer3")

}

func funcA() {
	fmt.Println("函数A")
}

func funcB() {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println(msg)
		}
	}()
	fmt.Println("函数B")
	defer myPrint("funcB defer1")
	for i := 1; i <= 10; i++ {
		fmt.Printf("i:%d\n", i)
		if i == 5 {
			panic("funcB panic")
		}
	}
	defer myPrint("funcB defer2")
}

func myPrint(v interface{}) {
	fmt.Println(v)
}
