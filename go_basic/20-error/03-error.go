package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//fp,err := os.Open("./test.txt")
	fp, err := os.Open("./test1.txt")
	if err != nil {
		if instance, ok := err.(*os.PathError); ok {
			fmt.Printf("%v|%T", instance, instance)
		}
	}
	fmt.Println(fp)

	addr, err := net.LookupHost("www.baidu.com")
	if err != nil {
		if instance, ok := err.(*net.DNSError); ok {
			fmt.Printf("%v|%T", instance, instance)

		}
	}
	fmt.Println(addr)
}
