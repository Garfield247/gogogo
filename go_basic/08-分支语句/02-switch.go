package main

import "fmt"

func main() {
	var grade string = "B"
	var marks int = 0

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	case 50:
		grade = "D"
	default:
		grade = "F"
	}
	fmt.Println(grade)

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B", grade == "C":
		fmt.Println("良好")
	case grade == "D":
		fmt.Println("及格")
	case grade == "F":
		fmt.Println("不及格")
	}
}
