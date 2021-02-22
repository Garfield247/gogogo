package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("自定义错误")
	fmt.Printf("%T%v\n", err1, err1)
	err2 := fmt.Errorf("错误:%d", 100)
	fmt.Printf("%T%v\n", err2, err2)
}
