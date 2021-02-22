package main

import "fmt"

func main() {
	// continue
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d,", i)
	}

	//break
	for i := 1; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Printf("%d,", i)
	}

}
