package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("./test.txt")

	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(fp)
}
