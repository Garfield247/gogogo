package main

import (
	"fmt"
	"math"
)

// 值传递
func main() {
	getSquareRoot := func(x float64) float64 {

		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}
