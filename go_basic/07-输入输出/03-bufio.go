package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("输入字符串:")
	//os.Stdin系统的输入
	reader := bufio.NewReader(os.Stdin)
	// 传入byte结束符
	s1, _ := reader.ReadString('\n')
	fmt.Printf("读取到:%s", s1)
}
