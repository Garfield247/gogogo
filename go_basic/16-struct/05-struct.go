package main

import "fmt"

func main() {
	s1 := struct {
		string
		int
	}{
		"xiaoming",
		10,
	}
	fmt.Println(s1)
}
