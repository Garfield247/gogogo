package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filepath := "./typescript"
	//fp, err := os.Open(filepath)
	//if err != nil {
	//fmt.Println("err:", err)
	//return
	//}
	//defer fp.Close()
	//bs := make([]byte, 4, 4)
	////n, _ := fp.Read(bs)
	////fmt.Println("n", n)
	////fmt.Println(string(bs))
	//n := -1
	//for {
	//n, err = fp.Read(bs)
	//if n == 0 || err == io.EOF {
	//break
	//}
	//fmt.Println(string(bs[:n]))
	//}

	fp, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Print("err", err)
	}
	fmt.Println(string(fp))

